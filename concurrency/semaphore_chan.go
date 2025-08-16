package concurrency

// Semaphore is a simple counting semaphore using a buffered channel.
type Semaphore struct {
	tickets chan struct{}
}

// NewSemaphore creates a new Semaphore with a fixed number of available slots (tickets).
// Each call to Acquire consumes one ticket; each call to Release returns one.
func NewSemaphore(ticketsNumber uint) *Semaphore {
	return &Semaphore{
		tickets: make(chan struct{}, ticketsNumber),
	}
}

// Acquire blocks until a ticket is available.
// If the maximum number of concurrent acquisitions has been reached, this call will block until Release is called.
func (s *Semaphore) Acquire() {
	if s == nil || s.tickets == nil {
		return
	}

	s.tickets <- struct{}{}
}

// Release frees up one ticket previously acquired via Acquire.
// If no goroutines are waiting, it simply makes space in the channel buffer.
func (s *Semaphore) Release() {
	if s == nil || s.tickets == nil {
		return
	}

	<-s.tickets
}
