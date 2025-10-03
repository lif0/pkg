package chanx_test

import (
	"context"
	"testing"
	"time"

	"github.com/lif0/pkg/concurrency/chanx"
	"github.com/stretchr/testify/assert"
)

func Sort[T string](arr []T) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Флаг для оптимизации: если не было перестановок, слайс уже отсортирован
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Обмен элементов
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// Если не было перестановок, выходим из цикла
		if !swapped {
			break
		}
	}
}

// TestFanInBasic verifies that FanIn merges values from multiple channels correctly.
func TestFanInBasic(t *testing.T) {
	ctx := context.Background()
	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)

	ch1 <- "a1"
	ch1 <- "a2"
	close(ch1)

	ch2 <- "b1"
	ch2 <- "b2"
	close(ch2)

	res := chanx.FanIn(ctx, ch1, ch2)

	expected := []string{"a1", "a2", "b1", "b2"}
	actual := make([]string, 0, len(expected))

	for v := range res {
		actual = append(actual, v)
	}

	if len(actual) != len(expected) {
		t.Errorf("Expected %d values, got %d: %v", len(expected), len(actual), actual)
	}

	Sort(actual) // because FaiIn not guaranteed order
	assert.EqualValues(t, expected, actual)
}

// TestFanInEmpty verifies that FanIn with no channels returns a closed channel immediately.
func TestFanInEmpty(t *testing.T) {
	ctx := context.Background()
	res := chanx.FanIn[any](ctx)

	// Should be closed immediately, no values.
	select {
	case _, ok := <-res:
		if ok {
			t.Error("Expected closed channel, but received a value")
		}
	default:
		// Channel is closed, as expected.
	}
}

// TestFanInOneChannel verifies that FanIn with a single channel forwards its values correctly.
func TestFanInOneChannel(t *testing.T) {
	ctx := context.Background()
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	res := chanx.FanIn(ctx, ch)

	var sum int
	for v := range res {
		sum += v
	}

	assert.Equal(t, 6, sum)
}

// TestFanInContextCancel verifies that canceling the context stops reading from channels and closes the result channel without forwarding values.
func TestFanInContextCancel(t *testing.T) {
	t.Run("ctx cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch := make(chan string, 1)
		res := chanx.FanIn(ctx, ch)

		cancel()

		// Send after cancel, but should not be received.
		go func() {
			defer close(ch)
			time.Sleep(10 * time.Millisecond)
			ch <- "delayed"
		}()

		select {
		case v, ok := <-res:
			if ok {
				t.Errorf("Received unexpected value %s after cancel", v)
			} else {
				// Channel closed without value - this is expected.
				return
			}
		case <-time.After(50 * time.Millisecond):
			t.Error("Result channel did not close in time after cancel")
		}
	})

	t.Run("select ctx.Done nested case", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		ch := make(chan string, 1)
		res := chanx.FanIn(ctx, ch)

		ch <- "ping"
		time.Sleep(time.Millisecond * 200) // anti-flak sleep
		cancel()

		select {
		case v, ok := <-res:
			if ok {
				t.Errorf("Received unexpected value %s after cancel", v)
			} else {
				// Channel closed without value - this is expected.
				return
			}
		case <-time.After(time.Second):
			t.Error("Result channel did not close in time after cancel")
		}
	})
}

// TestFanInMixed verifies that FanIn handles channels with different numbers of values.
func TestFanInMixed(t *testing.T) {
	ctx := context.Background()
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 3)

	ch1 <- "only-one"
	close(ch1)

	ch2 <- "b1"
	ch2 <- "b2"
	ch2 <- "b3"
	close(ch2)

	res := chanx.FanIn(ctx, ch1, ch2)

	expected := []string{"only-one", "b1", "b2", "b3"}
	actual := make([]string, 0, len(expected))

	for v := range res {
		actual = append(actual, v)
	}

	if len(actual) != len(expected) {
		t.Errorf("Expected %d values, got %d: %v", len(expected), len(actual), actual)
	}
}

// TestFanInChannelClose verifies that the result channel closes when all input channels are closed.
func TestFanInChannelClose(t *testing.T) {
	ctx := context.Background()
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Close inputs immediately.
	close(ch1)
	close(ch2)

	res := chanx.FanIn(ctx, ch1, ch2)

	// Should receive no values and close immediately.
	select {
	case _, ok := <-res:
		if ok {
			t.Error("Expected closed channel, but received a value")
		}
	default:
		// Closed, good.
	}
}

// TestFanInConcurrent verifies concurrent sending and fan-in behavior.
func TestFanInConcurrent(t *testing.T) {
	ctx := context.Background()

	// setup
	numChans := 10
	numValues := 100
	chans := make([]chan int, numChans)

	for i := 0; i < numChans; i++ {
		chans[i] = make(chan int)
		go func(chInx int) {
			defer close(chans[chInx])

			for j := 1; j <= numValues; j++ {
				chans[chInx] <- j
			}
		}(i)
	}

	// action
	resultIter := 0
	resultSum := 0
	res := chanx.FanIn[int](ctx, chanx.ToRecvChans[int](chans)...)

	for v := range res {
		resultIter++
		resultSum += v
	}

	// assert
	expectedIter := numChans * numValues
	expectedSum := ((numValues * (numValues + 1)) / 2) * numChans

	assert.Equal(t, resultIter, expectedIter, "Expected %d iteration count, got %d", resultIter, expectedIter)
	assert.Equal(t, resultSum, expectedSum, "Expected %d sum values, got %d", resultIter, expectedSum)
}
