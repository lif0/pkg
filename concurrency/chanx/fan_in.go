package chanx

import (
	"context"
	"sync"
)

// FanIn merges multiple input channels into a single output channel.
// It reads concurrently from each input channel and forwards values to the output.
// The output channel is closed when all input channels are closed or the context is canceled.
// This is a non-blocking, concurrent fan-in implementation that respects context cancellation.
//
// The order of values in the output is not guaranteed, as it depends on the race between goroutines.
//
// Example usage:
//
//	func main() {
//		ctx := context.Background()
//		chans := make([]chan string, 100)
//
//		for i := 0; i < numChans; i++ {
//			chans[i] = make(chan int)
//			go func(chInx int) {
//				defer close(chans[chInx])
//
//				for j := 1; j <= 100; j++ {
//					chans[chInx] <- j
//				}
//			}(i)
//		}
//
//		out := chanx.FanIn[int](ctx, chanx.ToRecvChans[int](chans)...)
//		for v := range out {
//			fmt.Println(v)
//		}
//	}
func FanIn[T any](ctx context.Context, chans ...<-chan T) <-chan T {
	res := make(chan T)
	wg := &sync.WaitGroup{}

	for i := 0; i < len(chans); i++ {
		wg.Add(1)
		go fia(ctx, wg, &chans[i], &res)
	}

	go func() {
		defer close(res)
		wg.Wait()
	}()

	return res
}

// FanIn action
func fia[T any](ctx context.Context, wg *sync.WaitGroup, argCh *<-chan T, argResult *chan T) {
	ch := *argCh
	result := *argResult

	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-ch:
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			case result <- v:
			}
		}
	}
}
