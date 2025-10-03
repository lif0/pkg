package concurrency_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/lif0/pkg/concurrency"
)

func Test_FutureAction(t *testing.T) {
	callback := func() any {
		time.Sleep(time.Millisecond)
		return "success"
	}

	future := concurrency.NewFutureAction(callback)
	time.Sleep(time.Millisecond * 500)
	result := future.Get()
	assert.Equal(t, "success", result)
}
