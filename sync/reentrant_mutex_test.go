package sync

import (
	"math/rand/v2"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMutexRecursive(t *testing.T) {
	t.Parallel()

	mx := NewReentrantMutex()

	mx.Lock()
	mx.Lock()
	mx.Lock()
	mx.Lock()
	mx.Lock()

	mx.Unlock()
	mx.Unlock()
	mx.Unlock()
	mx.Unlock()
	mx.Unlock()
}

func TestUnlockOfUnlockedMutex(t *testing.T) {
	t.Parallel()

	mx := NewReentrantMutex()
	require.PanicsWithError(t, ErrUnlockOfUnlockedMutex.Error(), func() {
		mx.Unlock()
	})
}

func TestUnlockFromAnotherGoroutine(t *testing.T) {
	t.Parallel()

	mx := NewReentrantMutex()

	mx.Lock()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		require.PanicsWithError(t, ErrUnlockFromAnotherGoroutine.Error(), func() {
			mx.Unlock()
		})
	}()

	wg.Wait()

	mx.Unlock()
}

func TestMutualExclusion(t *testing.T) {
	t.Parallel()

	v := make(map[int]int)
	rm := NewReentrantMutex()

	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			rm.Lock()
			v[rand.N[int](10e9)]++
			defer rm.Unlock()
		}()
	}

	wg.Wait()
}

func TestUnlockWithNegativeCount(t *testing.T) {
	t.Parallel()

	mx := NewReentrantMutex()

	mx.Lock()
	mx.hCall.Store(0) // Artificially set to trigger negative count branch

	require.PanicsWithError(t, ErrUnlockWithNegativeCount.Error(), func() {
		mx.Unlock()
	})
}

func TestRecursivePartialUnlockContention(t *testing.T) {
	t.Parallel()

	mx := NewReentrantMutex()

	mx.Lock()
	mx.Lock() // recursion level 2

	ch := make(chan bool)

	go func() {
		mx.Lock()
		ch <- true
		mx.Unlock()
	}()

	time.Sleep(100 * time.Millisecond)

	select {
	case <-ch:
		t.Fatal("acquired lock too early, while still recursively held")
	default:
		// good, blocked
	}

	mx.Unlock() // now recursion level 1

	time.Sleep(100 * time.Millisecond)

	select {
	case <-ch:
		t.Fatal("acquired lock too early, while still held")
	default:
		// good, still blocked
	}

	mx.Unlock() // now released

	select {
	case <-ch:
		// good, acquired
	case <-time.After(1 * time.Second):
		t.Fatal("failed to acquire lock after full unlock")
	}
}

func TestContentionSpin(t *testing.T) {
	t.Parallel()

	mx := NewReentrantMutex()

	mx.Lock()

	done := make(chan bool)

	go func() {
		mx.Lock()
		mx.Unlock()
		done <- true
	}()

	time.Sleep(100 * time.Millisecond) // ensure spinning occurs

	mx.Unlock()

	select {
	case <-done:
		// good
	case <-time.After(1 * time.Second):
		t.Fatal("failed to acquire after unlock")
	}
}

func TestMultipleRecursiveLevels(t *testing.T) {
	t.Parallel()

	mx := NewReentrantMutex()

	mx.Lock()
	mx.Lock()
	mx.Lock()

	mx.Unlock()
	mx.Unlock()

	// Still held at level 1
	ch := make(chan bool)
	go func() {
		mx.Lock()
		ch <- true
		mx.Unlock()
	}()

	time.Sleep(100 * time.Millisecond)

	select {
	case <-ch:
		t.Fatal("acquired lock too early")
	default:
	}

	mx.Unlock() // now released

	select {
	case <-ch:
		// good
	case <-time.After(1 * time.Second):
		t.Fatal("failed to acquire")
	}
}

func TestNewMutexIsUnlocked(t *testing.T) {
	t.Parallel()

	mx := NewReentrantMutex()

	// Should be able to lock immediately
	mx.Lock()
	mx.Unlock()

	// And unlock should panic if tried again
	require.PanicsWithError(t, ErrUnlockOfUnlockedMutex.Error(), func() {
		mx.Unlock()
	})
}

func TestMutexPerformance(t *testing.T) {
	// t.Skip()

	stdLibMx := testing.Benchmark(func(b *testing.B) {
		b.SetParallelism(10)

		v := make(map[int]int)
		mx := new(sync.Mutex)

		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mx.Lock()
				v[rand.N[int](10e9)]++
				mx.Unlock()
			}
		})
	})

	rmMx := testing.Benchmark(func(b *testing.B) {
		b.SetParallelism(10)

		v := make(map[int]int)
		mx := NewReentrantMutex()

		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mx.Lock()
				v[rand.N[int](10e9)]++
				mx.Unlock()
			}
		})
	})

	require.LessOrEqual(t, float64(stdLibMx.NsPerOp())/float64(rmMx.NsPerOp()), 4.0)
}
