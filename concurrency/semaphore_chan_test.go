package concurrency_test

import (
	"testing"
	"time"

	"github.com/lif0/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

func Test_Semaphore(t *testing.T) {
	t.Run("semaphore", func(t *testing.T) {
		sem := concurrency.NewSemaphore(2)

		go func() {
			time.Sleep(time.Millisecond * 200)
			sem.Release()
			sem.Release()
			assert.True(t, true)
		}()

		sem.Acquire()
		sem.Acquire()

		assert.True(t, true)
	})
}
