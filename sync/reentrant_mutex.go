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

const (
	// mutexFree indicates that the mutex is not held by any goroutine.
	mutexFreeFlag = -1
)

// ReentrantMutex is a reentrant mutual exclusion lock.
// The zero value for a ReentrantMutex is an unlocked mutex.
//
// ReentrantMutex allows the same goroutine to acquire the lock multiple times
// without deadlock. It tracks the owning goroutine and recursion level.
// Only the owning goroutine may unlock it, and the mutex is released when
// the recursion level reaches zero.
//
// ReentrantMutex implements the sync.Locker interface.
type ReentrantMutex struct {
	_ noCopy

	hCall atomic.Int64
	hID   atomic.Int64
}

// New creates and initializes a new ReentrantMutex.
//
// The returned mutex is initially free and implements the sync.Locker interface.
func New() *ReentrantMutex {
	rmu := ReentrantMutex{
		hID:   atomic.Int64{},
		hCall: atomic.Int64{},
	}

	rmu.hID.Store(mutexFreeFlag)

	return &rmu
}

// Lock locks rm.
// If the lock is already held by the current goroutine, the recursion count is incremented.
// Otherwise, Lock spins until the lock is acquired.
func (rm *ReentrantMutex) Lock() {
	gID := goid.Get()

	if rm.hID.Load() == gID {
		rm.hCall.Add(1)
		return
	}

	for !rm.hID.CompareAndSwap(mutexFreeFlag, gID) {
		runtime.Gosched()
	}

	rm.hCall.Store(1)
}

// Unlock unlocks rm.
// It must be called by the goroutine that owns the lock.
// If the recursion count is >1, it is decremented and nil is returned.
// If the recursion count reaches 0, the lock is released.
// Unlock returns an error if called on an unlocked mutex or by a non-owner goroutine.
func (rm *ReentrantMutex) Unlock() {
	gID := goid.Get()

	if rm.hID.Load() == mutexFreeFlag {
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
		rm.hID.Store(mutexFreeFlag)
	}
}
