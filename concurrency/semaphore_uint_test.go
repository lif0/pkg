package concurrency_test

import (
	"testing"
	"time"

	"github.com/lif0/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

func Test_Semaphore64(t *testing.T) {
	t.Run("semaphore64", func(t *testing.T) {
		sem := concurrency.NewSemaphore64()

		var ac1, ac2 int

		go func() {
			time.Sleep(time.Millisecond * 200)
			sem.Release(ac1)
			sem.Release(ac2)
			assert.True(t, true)
		}()

		ac1, _ = sem.Acquire()
		ac2, _ = sem.Acquire()

		assert.True(t, true)
	})
}
