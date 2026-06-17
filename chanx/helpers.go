package chanx

// ToRecvChans converts a slice of bidirectional channels to a slice of receive-only channels.
// This ensures that the channels can only be used for receiving values, preventing accidental sends.
// The function creates a new slice with the same length and copies the references.
//
// Example:
//
//	chans := []chan int{make(chan int), make(chan int)}
//	recvChans := ToRecvChans(chans)
//
// Now recvChans can be passed to functions expecting []<-chan int.
//
// Complexity:
//
//	time: O(n)
//	mem: O(n)
func ToRecvChans[T any](in []chan T) []<-chan T {
	out := make([]<-chan T, len(in))

	for i := 0; i < len(in); i++ {
		out[i] = in[i]
	}

	return out
}

// ToSendChans converts a slice of bidirectional channels to a slice of send-only channels.
// This ensures that the channels can only be used for sending values, preventing accidental receives.
// The function creates a new slice with the same length and copies the references.
//
// Example:
//
//		chans := []chan string{make(chan string), make(chan string)}
//		sendChans := ToSendChans(chans)
//
//	 Now sendChans can be passed to functions expecting []chan<- string.
//
// Complexity:
//
//	time: O(n)
//	mem: O(n)
func ToSendChans[T any](in []chan T) []chan<- T {
	out := make([]chan<- T, len(in))

	for i := 0; i < len(in); i++ {
		out[i] = in[i]
	}

	return out
}
