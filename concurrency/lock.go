package concurrency

import "sync"

// WithLock executes the given action while holding the provided lock.
//
// It accepts any sync.Locker (e.g., *sync.Mutex, *sync.RWMutex) and a function
// with no parameters or return values. If the action is nil, nothing is executed.
//
// The lock is guaranteed to be released after the action completes,
// even if the action panics or returns early.
func WithLock(mutex sync.Locker, action func()) {
	if action == nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	action()
}
