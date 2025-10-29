package structx_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/lif0/pkg/utils/structx"
	"github.com/stretchr/testify/assert"
)

type complexStruct struct {
	Val  int
	Nums []int
}

func Test_ObjectPool(t *testing.T) {
	t.Run("ok/uniq_ptr_for_each_object", func(t *testing.T) {
		pool := structx.NewObjectPool[complexStruct](1)

		hash := map[unsafe.Pointer]struct{}{}

		for i := 0; i < 10_000_000; i++ {
			csp := pool.Get()
			ptr := reflect.ValueOf(csp).UnsafePointer()
			if _, ok := hash[ptr]; ok {
				t.Fatal("sss")
			}

			hash[ptr] = struct{}{}
		}

		assert.Len(t, hash, 10_000_000)
	})

	t.Run("ok/", func(t *testing.T) {
		pool := structx.NewObjectPool[complexStruct](0)

		x := pool.Get()

		assert.NotNil(t, x)
	})
}
