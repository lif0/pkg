package concurrency

import "context"

// Semaphore is a counting semaphore that bounds the number of concurrent holders.
//
// The zero value (and a nil *Semaphore) is an unlimited semaphore: all acquire
// operations succeed immediately and Release is a no-op.
//
// All methods are safe for concurrent use by multiple goroutines.
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore returns a semaphore with the provided capacity.
//
// If capacity <= 0, it returns an unlimited semaphore, for which all acquire
// operations succeed immediately and Release does nothing.
func NewSemaphore(capacity uint) *Semaphore {
	if capacity <= 0 {
		return &Semaphore{ch: nil} // ch == nil -> unlimited semaphore
	}

	return &Semaphore{ch: make(chan struct{}, capacity)}
}

// Acquire obtains one slot from s, blocking until a slot is available.
// For an unlimited semaphore, Acquire is a no-op.
func (s *Semaphore) Acquire() {
	if s == nil || s.ch == nil {
		return
	}

	s.ch <- struct{}{}
}

// AcquireContext attempts to obtain one slot, blocking until a slot is available
// or the context is canceled or its deadline is exceeded.
// It returns ctx.Err() if the context is done first.
// For an unlimited semaphore, AcquireContext returns nil immediately.
func (s *Semaphore) AcquireContext(ctx context.Context) error {
	if s == nil || s.ch == nil {
		return nil
	}
	select {
	case s.ch <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// TryAcquire attempts to obtain one slot without blocking.
// It returns true if a slot was acquired and false otherwise.
// For an unlimited semaphore, TryAcquire always returns true.
func (s *Semaphore) TryAcquire() bool {
	if s == nil || s.ch == nil {
		return true // без лимита — всегда успешно
	}
	select {
	case s.ch <- struct{}{}:
		return true
	default:
		return false
	}
}

// Release releases one previously acquired slot.
// On a limited semaphore, calling Release without a matching acquire panics.
// On an unlimited semaphore, Release is a no-op.
func (s *Semaphore) Release() {
	if s == nil || s.ch == nil {
		return
	}
	select {
	case <-s.ch:
		return
	default:
		panic("concurrency.Semaphore: release without matching acquire")
	}
}

// InUse reports the current number of acquired slots.
// For an unlimited semaphore, InUse returns 0.
func (s *Semaphore) InUse() int {
	if s == nil || s.ch == nil {
		return 0
	}
	return len(s.ch)
}

// Cap returns the maximum number of concurrent holders (the capacity).
// For an unlimited semaphore, Cap returns 0.
func (s *Semaphore) Cap() int {
	if s == nil || s.ch == nil {
		return 0
	}
	return cap(s.ch)
}
