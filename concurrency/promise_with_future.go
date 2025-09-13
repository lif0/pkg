package concurrency

import (
	"sync"
	"sync/atomic"
)

// PromiseError and FutureError are type aliases for Promise and Future specialized for error handling.
// This allows for easy propagation of errors in asynchronous operations.
type (
	PromiseError = Promise[error]
	FutureError  = Future[error]
)

// Promise represents a writable, single-assignment container for a future value.
// It allows setting a value exactly once. Attempting to set the value more than once is ignored.
// The internal channel is buffered to hold one value and is closed after setting.
// Synchronization is handled via atomic operations and a mutex for thread safety.
//
// Example usage:
//
//	func main() {
//		promise := NewPromise[string]()
//		go func() {
//			time.Sleep(time.Second)
//			promise.Set("Cake")
//		}()
//
//	    future := promise.GetFuture()
//	    value := future.Get()
//	    fmt.Println(value) // Output: Cake
//	}
type Promise[T any] struct {
	result   chan T
	promised atomic.Bool
	mu       sync.Mutex
}

// Future represents a read-only view of a promised value.
// It provides a way to retrieve the value asynchronously, blocking if necessary.
type Future[T any] struct {
	result <-chan T
}

// NewPromise creates and returns a new Promise.
// The internal channel is buffered with capacity 1 to hold the future value.
func NewPromise[T any]() Promise[T] {
	return Promise[T]{
		result: make(chan T, 1),
	}
}

// Set assigns the value to the Promise.
// This can be called only once; subsequent calls are ignored.
// After setting, the value is sent to the channel, and the channel is closed.
func (p *Promise[T]) Set(value T) {
	if !p.promised.CompareAndSwap(false, true) {
		return
	}

	p.result <- value
	close(p.result)
}

// GetFuture returns a Future associated with this Promise.
// The Future can be used to retrieve the value once it's set.
func (p *Promise[T]) GetFuture() *Future[T] {
	return NewFuture[T](p.result)
}

// NewFuture creates a new Future from a given receive-only channel.
// This allows wrapping an existing channel as a Future.
func NewFuture[T any](result <-chan T) *Future[T] {
	return &Future[T]{
		result: result,
	}
}

// Get retrieves the value from the Future, blocking until it's available.
// If the channel is closed without a value (though not typical in this pattern), it returns the zero value.
func (f *Future[T]) Get() T {
	return <-f.result
}
