// Package sync provides sync utilities.
package sync

import (
	"errors"
	"runtime"
	"sync/atomic"

	"github.com/petermattis/goid"
)

var (
	// ErrUnlockOfUnlockedMutex reports an attempt to unlock an unlocked mutex.
	ErrUnlockOfUnlockedMutex = errors.New("unlock of unlocked reentrant mutex")

	// ErrUnlockFromAnotherGoroutine reports an attempt to unlock a mutex owned by another goroutine.
	ErrUnlockFromAnotherGoroutine = errors.New("unlock from non-owner goroutine")

	// ErrUnlockWithNegativeCount Unlock with negative count.
	ErrUnlockWithNegativeCount = errors.New("unlock with negative count")
)

// A ReentrantMutex is a reentrant mutual exclusion lock.
// The zero value for a ReentrantMutex is an unlocked mutex.
//
// A ReentrantMutex must not be copied after first use.
//
// In the terminology of [the Go memory model],
// the n'th call to [ReentrantMutex.Unlock] by the owning goroutine "synchronizes before"
// the m'th call to [ReentrantMutex.Lock] for any n < m, accounting for recursion levels.
// ReentrantMutex allows the same goroutine to acquire the lock multiple times without deadlock.
// It tracks the owning goroutine and recursion level; only the owning goroutine may unlock it,
// and the mutex is released when the recursion level reaches zero.
//
// ReentrantMutex implements the sync.Locker interface.
//
// [the Go memory model]: https://go.dev/ref/mem
type ReentrantMutex struct {
	_ noCopy

	hCall   atomic.Int64
	hID     atomic.Int64
	notFree atomic.Bool
}

// NewReentrantMutex creates and initializes a new ReentrantMutex.
func NewReentrantMutex() *ReentrantMutex {
	return &ReentrantMutex{}
}

// Lock locks rm.
// If the lock is already held by the current goroutine, the recursion count is incremented.
// Otherwise, the calling goroutine blocks until the rmutex is available.
func (rm *ReentrantMutex) Lock() {
	gID := goid.Get()

	if rm.hID.Load() == gID {
		rm.hCall.Add(1)
		return
	}

	for !rm.notFree.CompareAndSwap(false, true) {
		runtime.Gosched()
	}

	rm.hID.Store(gID)
	rm.hCall.Store(1)
}

// Unlock unlocks rm.
// It panics if rm is not locked on entry to Unlock.
//
// Unlock must be called by the goroutine that owns the lock.
// If the recursion count is greater than 1, it is decremented.
// If the recursion count reaches 0, the lock is released.
//
// A locked [ReentrantMutex] is associated with a particular goroutine.
// It is not allowed for one goroutine to lock a ReentrantMutex and then
// arrange for another goroutine to unlock it.
func (rm *ReentrantMutex) Unlock() {
	gID := goid.Get()

	if !rm.notFree.Load() {
		panic(ErrUnlockOfUnlockedMutex)
	}

	if rm.hID.Load() != gID {
		panic(ErrUnlockFromAnotherGoroutine)
	}

	newCount := rm.hCall.Add(-1)
	if newCount < 0 {
		panic(ErrUnlockWithNegativeCount)
	}
	if newCount == 0 {
		rm.hID.Store(-1)
		rm.notFree.Store(false)
	}
}
