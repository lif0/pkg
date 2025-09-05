package concurrency_test

import (
	"context"
	"testing"
	"time"

	"github.com/lif0/pkg/concurrency"
)

func mustPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic, got none")
		}
	}()
	fn()
}

func TestNewSemaphore_UnlimitedAndLimited(t *testing.T) {
	s0 := concurrency.NewSemaphore(0)
	if got := s0.Cap(); got != 0 {
		t.Fatalf("Cap(unlimited) = %d, want 0", got)
	}
	if got := s0.InUse(); got != 0 {
		t.Fatalf("InUse(unlimited) = %d, want 0", got)
	}

	s3 := concurrency.NewSemaphore(3)
	if got := s3.Cap(); got != 3 {
		t.Fatalf("Cap(limited) = %d, want 3", got)
	}
	if got := s3.InUse(); got != 0 {
		t.Fatalf("InUse(initial limited) = %d, want 0", got)
	}
}

func TestNilSemaphore_Behavior(t *testing.T) {
	var s *concurrency.Semaphore // nil receiver must be safe

	// no panics and "unlimited" utils
	s.Acquire()
	if err := s.AcquireContext(context.Background()); err != nil {
		t.Fatalf("AcquireContext(nil) error = %v, want nil", err)
	}
	if ok := s.TryAcquire(); !ok {
		t.Fatalf("TryAcquire(nil) = false, want true")
	}
	s.Release() // no panic
	if s.InUse() != 0 {
		t.Fatalf("InUse(nil) = %d, want 0", s.InUse())
	}
	if s.Cap() != 0 {
		t.Fatalf("Cap(nil) = %d, want 0", s.Cap())
	}
}

func TestUnlimitedSemaphore_Methods(t *testing.T) {
	s := concurrency.NewSemaphore(0)

	// Acquire / Release are no-ops
	s.Acquire()
	s.Release()

	// TryAcquire always true
	if ok := s.TryAcquire(); !ok {
		t.Fatalf("TryAcquire(unlimited) = false, want true")
	}

	// AcquireContext returns nil
	if err := s.AcquireContext(context.Background()); err != nil {
		t.Fatalf("AcquireContext(unlimited) error = %v, want nil", err)
	}

	// InUse/Cap stay 0
	if s.InUse() != 0 || s.Cap() != 0 {
		t.Fatalf("unlimited InUse=%d Cap=%d, want 0/0", s.InUse(), s.Cap())
	}
}

func TestAcquireRelease_LimitedAndTryAcquire(t *testing.T) {
	s := concurrency.NewSemaphore(2)

	// acquire to capacity
	s.Acquire()
	if s.InUse() != 1 {
		t.Fatalf("InUse after 1 Acquire = %d, want 1", s.InUse())
	}
	s.Acquire()
	if s.InUse() != 2 {
		t.Fatalf("InUse after 2 Acquire = %d, want 2", s.InUse())
	}

	// capacity full -> TryAcquire false
	if ok := s.TryAcquire(); ok {
		t.Fatalf("TryAcquire when full = true, want false")
	}

	// release one -> TryAcquire becomes true
	s.Release()
	if s.InUse() != 1 {
		t.Fatalf("InUse after 1 Release = %d, want 1", s.InUse())
	}
	if ok := s.TryAcquire(); !ok { // this also increments InUse
		t.Fatalf("TryAcquire after Release = false, want true")
	}
	if s.InUse() != 2 {
		t.Fatalf("InUse after successful TryAcquire = %d, want 2", s.InUse())
	}

	// drain
	s.Release()
	s.Release()
	if s.InUse() != 0 {
		t.Fatalf("InUse after draining = %d, want 0", s.InUse())
	}
}

func TestAcquireBlocksUntilRelease(t *testing.T) {
	s := concurrency.NewSemaphore(1)

	// fill the only slot
	s.Acquire()

	done := make(chan struct{})
	go func() {
		s.Acquire() // should block until Release
		close(done)
	}()

	// ensure goroutine is blocked
	select {
	case <-done:
		t.Fatalf("Acquire did not block at capacity")
	case <-time.After(50 * time.Millisecond):
		// still blocked — good
	}

	// now release and expect goroutine to proceed
	s.Release()
	select {
	case <-done:
		// ok
	case <-time.After(250 * time.Millisecond):
		t.Fatalf("Acquire did not unblock after Release")
	}

	// cleanup
	s.Release()
}

func TestAcquireContext_Canceled(t *testing.T) {
	s := concurrency.NewSemaphore(1)
	s.Acquire() // fill

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	start := time.Now()
	err := s.AcquireContext(ctx)
	if err == nil {
		t.Fatalf("AcquireContext expected error on cancel/deadline, got nil")
	}
	// should complete roughly around the timeout (not block indefinitely)
	if time.Since(start) > time.Second {
		t.Fatalf("AcquireContext took too long, likely blocked")
	}

	// cleanup
	s.Release()
}

func TestAcquireContext_Succeeds(t *testing.T) {
	s := concurrency.NewSemaphore(2)
	if err := s.AcquireContext(context.Background()); err != nil {
		t.Fatalf("AcquireContext should succeed, got %v", err)
	}
	if s.InUse() != 1 {
		t.Fatalf("InUse after AcquireContext = %d, want 1", s.InUse())
	}
	s.Release()
}

func TestRelease_PanicsOnOverRelease(t *testing.T) {
	s := concurrency.NewSemaphore(1)

	// over-release on limited semaphore must panic
	mustPanic(t, func() {
		s.Release()
	})

	// but on unlimited it's a no-op
	su := concurrency.NewSemaphore(0)
	su.Release() // must not panic
}
