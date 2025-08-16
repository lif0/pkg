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

// NewSemaphore64 creates a new Semaphore with a fixed number of available slots (tickets).
// Each call to Acquire consumes one ticket; each call to Release returns one.
func NewSemaphore64() *Semaphore64 {
	return &Semaphore64{}
}

// Acquire пытается занять свободный слот, возвращает индекс занятого слота или ошибку
func (s *Semaphore64) Acquire() (int, error) {
	for {
		old := atomic.LoadUint64(&s.slots)
		if old == ^uint64(0) { // все 64 слота заняты
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

// Release освобождает слот по индексу
func (s *Semaphore64) Release(index int) {
	mask := ^(uint64(1) << index)
	atomic.AndUint64(&s.slots, mask)
}

// Used возвращает число занятых слотов
func (s *Semaphore64) Used() int {
	return bits.OnesCount64(atomic.LoadUint64(&s.slots))
}
