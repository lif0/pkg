package concurrency

// FutureAction is an abstraction over a channel that models a task and its result.
// It allows executing a computation asynchronously in a goroutine and retrieving
// the result later via a blocking call. This is similar to the Future pattern in
// other languages, providing a simple way to handle asynchronous results without
// manual channel management.
//
// The channel is closed after the result is sent, ensuring proper resource cleanup.
//
// Example usage:
//
//	func main() {
//		callback := func() any {
//			time.Sleep(time.Second)
//			return "success"
//		}
//
//		future := NewFutureAction(callback)
//		result := future.Get()
//		fmt.Println(result) // Output: success
//	}
type FutureAction[T any] struct {
	result chan T
}

// NewFutureAction creates and returns a new FutureAction.
// It starts the provided action function in a separate goroutine.
// The action's return value is sent to the internal channel.
// The channel is closed after sending the result to allow safe ranging or detection of completion.
func NewFutureAction[T any](action func() T) *FutureAction[T] {
	future := &FutureAction[T]{
		result: make(chan T, 1),
	}

	go func() {
		defer close(future.result)
		future.result <- action()
	}()

	return future
}

// Get returns the result of the asynchronous task.
// This method blocks until the result is available from the channel.
// If the action function blocks indefinitely (e.g., due to an infinite loop or deadlock),
// Get will never return, potentially causing the caller to hang.
// It is the caller's responsibility to ensure the action completes.
func (f *FutureAction[T]) Get() T {
	return <-f.result
}
