package concurrency_test

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/lif0/pkg/concurrency"
)

func TestPromiseSet(t *testing.T) {
	p := concurrency.NewPromise[string]()
	p.Set("test")

	f := p.GetFuture()
	value := f.Get()
	if value != "test" {
		t.Errorf("Set: expected value 'test', got '%s'", value)
	}
}

func TestPromiseSetMultiple(t *testing.T) {
	p := concurrency.NewPromise[string]()
	p.Set("first")
	p.Set("second") // Should be ignored.

	f := p.GetFuture()
	value := f.Get()
	if value != "first" {
		t.Errorf("SetMultiple: expected value 'first', got '%s'", value)
	}
}

func TestGetFuture(t *testing.T) {
	p := concurrency.NewPromise[string]()
	p.Set("test")

	f := p.GetFuture()
	value := f.Get()
	if value != "test" {
		t.Errorf("GetFuture: expected value 'test', got '%s'", value)
	}
}

func TestNewFuture(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "test"
	close(ch)

	f := concurrency.NewFuture[string](ch)
	value := f.Get()
	if value != "test" {
		t.Errorf("NewFuture: expected value 'test', got '%s'", value)
	}
}

func TestFutureGet(t *testing.T) {
	p := concurrency.NewPromise[string]()
	go func() {
		time.Sleep(time.Millisecond * 50)
		p.Set("test")
	}()

	f := p.GetFuture()
	value := f.Get()
	if value != "test" {
		t.Errorf("Get: expected value 'test', got '%s'", value)
	}
}

func TestFutureGetEmptyChannel(t *testing.T) {
	ch := make(chan string, 1)
	close(ch)
	f := concurrency.NewFuture[string](ch)
	value := f.Get()
	if value != "" {
		t.Errorf("Get: expected zero value for closed empty channel, got '%s'", value)
	}
}

func TestPromiseError(t *testing.T) {
	p := concurrency.NewPromise[error]()
	expectedErr := errors.New("test error")
	go func() {
		time.Sleep(time.Millisecond * 50)
		p.Set(expectedErr)
	}()

	f := p.GetFuture()
	err := f.Get()
	if err.Error() != expectedErr.Error() {
		t.Errorf("PromiseError: expected error '%v', got '%v'", expectedErr, err)
	}
}

func TestConcurrentSet(t *testing.T) {
	p := concurrency.NewPromise[string]()
	var wg sync.WaitGroup
	n := 10
	wg.Add(n)
	values := make(map[string]bool)
	mu := sync.Mutex{}

	for i := 0; i < n; i++ {
		i := i
		go func() {
			defer wg.Done()
			val := "value" + string(rune('0'+i))
			p.Set(val)

			mu.Lock()
			values[val] = true
			mu.Unlock()
		}()
	}

	wg.Wait()
	f := p.GetFuture()
	value := f.Get()

	// Verify that the value is one of the attempted ones.
	if !values[value] {
		t.Errorf("ConcurrentSet: got unexpected value '%s'", value)
	}

	// Since we can't check if channel has only one value directly, we rely on the fact that Get returns one,
	// and multiple Sets are ignored after the first.
	// Additional check: try to read again; should be zero since closed.
	zero := f.Get()
	if zero != "" {
		t.Errorf("ConcurrentSet: expected zero value on second Get, got '%s'", zero)
	}
}
