package concurrency

import (
	"errors"
	"math/bits"
	"sync/atomic"
)

type Semaphore64 struct {
	slots uint64
}

var ErrNoFreeSlot = errors.New("no free slot")

// NewSemaphore64 creates a new Semaphore64 instance.
// Each call to Acquire occupies one slot; each call to Release frees one.
func NewSemaphore64() *Semaphore64 {
	return &Semaphore64{}
}

// Acquire tries to occupy a free slot.
// Returns the index of the acquired slot, or an error if all slots are in use.
func (s *Semaphore64) Acquire() (int, error) {
	for {
		old := atomic.LoadUint64(&s.slots)
		if old == ^uint64(0) { // all 64 slots are occupied
			return -1, ErrNoFreeSlot
		}

		freeBit := bits.TrailingZeros64(^old)
		if freeBit >= 64 {
			return -1, ErrNoFreeSlot
		}

		mask := uint64(1) << freeBit
		if atomic.CompareAndSwapUint64(&s.slots, old, old|mask) {
			return freeBit, nil
		}
	}
}

// Release frees a previously acquired slot by its index.
func (s *Semaphore64) Release(index int) {
	mask := ^(uint64(1) << index)
	atomic.AndUint64(&s.slots, mask)
}

// Used returns the number of currently occupied slots.
func (s *Semaphore64) Used() int {
	return bits.OnesCount64(atomic.LoadUint64(&s.slots))
}
