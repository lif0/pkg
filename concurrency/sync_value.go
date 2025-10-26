package concurrency

import "sync"

// SyncValue is a generic wrapper around a value of any type `T` that allows
// concurrent access with safe read and write operations protected by an RWMutex.
//
// It provides two methods: MutateValue and ReadValue.
//   - MutateValue gives exclusive write access to the wrapped value.
//   - ReadValue gives shared read access to the wrapped value.
//
// Inside MutateValue, mu.Lock() / mu.Unlock() are used.
// Inside ReadValue, mu.RLock() / mu.RUnlock() are used.
//
// Example:
//
//	var obj ComplexObject
//
//	m := NewSyncValue[ComplexObject](ComplexObject{Slice: ..., Map: ...})
//
//	// Safe mutation
//	m.MutateValue(func(v *ComplexObject) {
//		v.Slice = append(v.Slice, 100)
//	})
//
//	// Safe read
//	var safeCopy ComplexObject
//	m.ReadValue(func(v *ComplexObject) {
//		safeCopy = v.DeepCopy() // good way
//		// safeCopy = *v // BAD WAY: shallow copy may share internal memory
//	})
//
// Note:
//   - Do NOT store the pointer passed into the callback for later use; it is only
//     valid within the callback.
//   - When T is a reference type (e.g., slice, map, pointer, channel), copying `*v`
//     only performs a shallow copy, meaning the underlying data may still be shared.
//     Use DeepCopy or an explicit cloning function to avoid data races.
type SyncValue[T any] struct {
	mu sync.RWMutex
	v  T
}

// NewSyncValue constructs a SyncValue initialized with the provided value.
func NewSyncValue[T any](value ...T) *SyncValue[T] {
	var val T

	if len(value) > 0 {
		val = value[0]
	}

	return &SyncValue[T]{
		v:  val,
		mu: sync.RWMutex{},
	}
}

// MutateValue provides exclusive, write access to the wrapped value by invoking
// the supplied function while holding an exclusive lock. The callback receives a
// pointer to the internal value, allowing in-place updates.
//
// IMPORTANT:
//   - Do not store the pointer beyond the duration of the callback.
//   - Avoid calling other methods of SyncValue from within the callback to
//     prevent lock re-entrancy issues (Go mutexes are not re-entrant).
//
// Example (mutating a struct field and a slice in-place):
//
//	type Config struct {
//	    Enabled bool
//	    Items   []int
//	}
//
//	cfg := NewSyncValue(Config{Enabled: false, Items: []int{1, 2}})
//	cfg.MutateValue(func(v *Config) {
//	    v.Enabled = true
//	    v.Items = append(v.Items, 3)
//	})
func (sv *SyncValue[T]) MutateValue(f func(v *T)) {
	sv.mu.Lock()
	defer sv.mu.Unlock()

	f(&sv.v)
}

// ReadValue provides shared, read access to the wrapped value by invoking the
// supplied function while holding a shared lock. The callback receives a pointer
// to the internal value for efficient inspection.
//
// IMPORTANT:
//   - The pointer is only safe to use within the callback.
//   - If T (or its fields) contains reference types (e.g., slices/maps), copying
//     *v will be shallow. To safely use the value after the callback returns,
//     make a defensive deep copy.
//
// Example (creating a defensive copy of a slice):
//
//	sv := NewSyncValue([]int{1, 2, 3})
//	var snapshot []int
//	sv.ReadValue(func(v *[]int) {
//		snapshot = make([]int, len(*v)) // copy does't alloc memory!! You should call make for it.
//		copy(snapshot, *v) // Defensive copy: underlying array is not shared.
//	})
func (sv *SyncValue[T]) ReadValue(f func(v *T)) {
	sv.mu.RLock()
	defer sv.mu.RUnlock()

	f(&sv.v)
}
