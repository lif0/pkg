package concurrency_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/lif0/pkg/concurrency"
)

func Test_FutureAction(t *testing.T) {
	callback := func() any {
		time.Sleep(time.Second)
		return "success"
	}

	future := concurrency.NewFutureAction(callback)
	result := future.Get()
	assert.Equal(t, "success", result)
}
