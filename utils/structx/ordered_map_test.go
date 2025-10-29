package structx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lif0/pkg/utils/structx"
)

func Test_OrderedMap_Iter(t *testing.T) {
	t.Parallel()

	t.Run("ok/iter-kv", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[int, int](10)
		res := 0

		// act
		for i := range 10 {
			m.Put(i+1, i+1)
		}
		for k, v := range m.Iter() {
			res += k + v
		}

		// assert
		assert.Equal(t, 110, res)
	})

	t.Run("ok/iter-kv-2", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[int, int]()

		// act
		for i := range 10 {
			m.Put(i, i)
		}

		// assert
		expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		gotKey := make([]int, 0, 10)
		gotValue := make([]int, 0, 10)
		for k, v := range m.Iter() {
			gotKey = append(gotKey, k)
			gotValue = append(gotValue, v)
		}
		assert.Equal(t, expected, gotKey)
		assert.Equal(t, expected, gotValue)
	})

	t.Run("ok/iter-kv-3", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[int, int]()

		m.Put(1, 1)

		m.Iter()(func(i1, i2 int) bool {
			return false
		})
	})
}

func Test_OrderedMap_Get(t *testing.T) {
	t.Parallel()

	t.Run("ok/after-put", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[string, int]()

		// act
		m.Put("x", 42)
		got, ok := m.Get("x")

		// assert
		assert.True(t, ok)
		assert.Equal(t, 42, got)
	})

	t.Run("edge/unknown-key", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[string, int]()

		// act
		v, ok := m.Get("b")

		// assert
		assert.False(t, ok)
		assert.Equal(t, 0, v)
	})

	t.Run("edge/zero-value-safe-get", func(t *testing.T) {
		t.Parallel()
		var m structx.OrderedMap[string, int]

		// act
		v, ok := m.Get("kek")

		// assert
		assert.False(t, ok) // safe read from nil-map
		assert.Equal(t, 0, v)
	})
}

func Test_OrderedMap_Put(t *testing.T) {
	t.Parallel()

	t.Run("panic/nil-dict", func(t *testing.T) {
		t.Parallel()
		var m structx.OrderedMap[int, int]

		// act + assert
		assert.Panics(t, func() {
			m.Put(1, 10) // запись в nil map внутри dict
		})
	})

	t.Run("ok/insert-order", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[string, int]()

		// act
		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)

		// assert
		vals := m.GetValues()
		assert.Equal(t, []int{1, 2, 3}, vals)
	})

	t.Run("ok/update-existing-preserve-order", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[string, int]()
		m.Put("a", 1)
		m.Put("b", 2)

		// act
		m.Put("a", 100)

		// assert
		gotA, okA := m.Get("a")
		assert.True(t, okA)
		assert.Equal(t, 100, gotA)

		vals := m.GetValues()
		// order var is not changed: [a, b]
		assert.Equal(t, []int{100, 2}, vals)
	})
}

func Test_OrderedMap_Delete(t *testing.T) {
	t.Parallel()

	t.Run("ok/delete-existing", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[string, int]()
		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)

		// act
		m.Delete("b")

		// assert
		_, okB := m.Get("b")
		assert.False(t, okB)

		vals := m.GetValues()
		assert.Equal(t, []int{1, 3}, vals)
	})

	t.Run("ok/build-in-delete-not-existing", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[string, int]()
		m.Put("a", 1)
		m.Put("c", 3)

		// act
		structx.Delete(m, "b")

		// assert
		_, okB := m.Get("b")
		assert.False(t, okB)

		vals := m.GetValues()
		assert.Equal(t, []int{1, 3}, vals)
	})

	t.Run("edge/nil-dict", func(t *testing.T) {
		t.Parallel()
		var m structx.OrderedMap[int, int] // nil

		// act: should't panic
		m.Delete(1)

		// assert: the struct still empty
		assert.Equal(t, 0, len(m.GetValues()))
	})
}

func Test_OrderedMap_BuiltInDelete(t *testing.T) {
	t.Parallel()

	t.Run("ok/delete", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[string, int]()
		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)

		// act
		structx.Delete(m, "b")

		// assert
		_, okB := m.Get("b")
		assert.False(t, okB)

		vals := m.GetValues()
		assert.Equal(t, []int{1, 3}, vals)
	})

	t.Run("ok/delete-not-existing", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[string, int]()
		m.Put("a", 1)
		m.Put("c", 3)

		// act
		structx.Delete(m, "b")

		// assert
		_, okB := m.Get("b")
		assert.False(t, okB)

		vals := m.GetValues()
		assert.Equal(t, []int{1, 3}, vals)
	})

	t.Run("ok/nil", func(t *testing.T) {
		t.Parallel()
		var m *structx.OrderedMap[int, int] // nil

		// act: should't panic
		structx.Delete(m, 1)

		// assert
		assert.Nil(t, m)
	})

	t.Run("ok/empty", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[int, int]()

		// act: should't panic
		structx.Delete(m, 1)

		// assert: the struct still empty
		assert.Equal(t, 0, len(m.GetValues()))
	})
}

func Test_OrderedMap_GetValues(t *testing.T) {
	t.Parallel()

	t.Run("edge/empty", func(t *testing.T) {
		t.Parallel()
		var m structx.OrderedMap[string, int]

		// act
		vals := m.GetValues()

		// assert
		assert.NotNil(t, vals)
		assert.Equal(t, 0, len(vals))
	})

	t.Run("ok/single", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[string, int]()
		m.Put("a", 7)

		// act
		vals := m.GetValues()

		// assert
		assert.Equal(t, 1, len(vals))
		assert.Equal(t, []int{7}, vals)
	})

	t.Run("ok/multiple", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[int, string]()
		m.Put(10, "x")
		m.Put(20, "y")
		m.Put(30, "z")

		// act
		vals := m.GetValues()

		// assert
		assert.Equal(t, []string{"x", "y", "z"}, vals)
	})
}

func Test_NewOrderedMap(t *testing.T) {
	t.Parallel()

	t.Run("ok/empty", func(t *testing.T) {
		t.Parallel()
		m := structx.NewOrderedMap[string, int]()

		// act
		v, ok := m.Get("missing")
		values := m.GetValues()

		// assert
		assert.False(t, ok)
		assert.Equal(t, 0, v)
		assert.NotNil(t, values)
		assert.Equal(t, 0, len(values))
	})
}
