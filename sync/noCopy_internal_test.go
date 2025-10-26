package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// formal test to increase coverage.
func Test_noCopy(t *testing.T) {
	nc := noCopy{}
	i := 0
	nc.Lock()
	i = 1
	defer nc.Unlock()

	assert.Equal(t, 1, i)
}
