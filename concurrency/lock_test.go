package concurrency_test

import (
	"sync"
	"testing"

	"github.com/lif0/pkg/concurrency"
)

func TestWithLock_ActionExecuted(t *testing.T) {
	var mu sync.Mutex
	called := false

	concurrency.WithLock(&mu, func() {
		called = true
	})

	if !called {
		t.Error("expected action to be executed, but it was not")
	}
}

func TestWithLock_ActionNil(t *testing.T) {
	var mu sync.Mutex
	concurrency.WithLock(&mu, nil)
}

func TestWithLock_LockerNil(t *testing.T) {
	var mu *sync.Mutex = nil

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic when mutex is nil, but did not panic")
		}
	}()

	concurrency.WithLock(mu, func() {})
}

func TestWithLock_MutexActuallyLocks(t *testing.T) {
	var mu sync.Mutex
	var counter int

	const goroutines = 100
	const increments = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				concurrency.WithLock(&mu, func() {
					counter++
				})
			}
		}()
	}

	wg.Wait()

	expected := goroutines * increments
	if counter != expected {
		t.Errorf("expected counter = %d, got %d", expected, counter)
	}
}
