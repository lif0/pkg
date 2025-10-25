package concurrency

import "sync"

type SyncValue[T any] struct {
	v  T
	mu sync.Mutex
}

func NewSyncValue[T any](value T) *SyncValue[T] {
	return &SyncValue[T]{
		v:  value,
		mu: sync.Mutex{},
	}
}

func (sv *SyncValue[T]) MutateValue(f func(v *T)) {
	sv.mu.Lock()
	defer sv.mu.Unlock()

	f(&sv.v)
}

func (sv *SyncValue[T]) GetValue() T {
	sv.mu.Lock()
	defer sv.mu.Unlock()

	return sv.v
}
