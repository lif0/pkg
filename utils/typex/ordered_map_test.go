package typex_test

import (
	"testing"

	"github.com/lif0/pkg/utils/typex"
	"github.com/stretchr/testify/assert"
)

// // arrange
// func newOrderedMap[K comparable, V any]() *typex.OrderedMap[K, V] {
// 	var m typex.OrderedMap[K, V]
// 	initOrderedMapDict(&m)
// 	return &m
// }

// // arrange
// func initOrderedMapDict[K comparable, V any](m *typex.OrderedMap[K, V]) {
// 	mv := reflect.ValueOf(m).Elem()
// 	dictField := mv.FieldByName("dict")
// 	// assert.NotNil(t, dictField, "dict field must exist")

// 	// make map with the exact field type, without importing internal types
// 	newMap := reflect.MakeMapWithSize(dictField.Type(), 0)

// 	// set unexported field via unsafe
// 	dictPtr := unsafe.Pointer(dictField.UnsafeAddr())
// 	reflect.NewAt(dictField.Type(), dictPtr).Elem().Set(newMap)
// }

func Test_OrderedMap_Get(t *testing.T) {
	t.Parallel()

	t.Run("ok/after-put", func(t *testing.T) {
		t.Parallel()
		m := typex.NewOrderedMap[string, int]()

		// act
		m.Put("x", 42)
		got, ok := m.Get("x")

		// assert
		assert.True(t, ok)
		assert.Equal(t, 42, got)
	})

	t.Run("edge/unknown-key", func(t *testing.T) {
		t.Parallel()
		m := typex.NewOrderedMap[string, int]()

		// act
		v, ok := m.Get("b")

		// assert
		assert.False(t, ok)
		assert.Equal(t, 0, v)
	})

	t.Run("edge/zero-value-safe-get", func(t *testing.T) {
		t.Parallel()
		var m typex.OrderedMap[string, int]

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
		var m typex.OrderedMap[int, int]

		// act + assert
		assert.Panics(t, func() {
			m.Put(1, 10) // запись в nil map внутри dict
		})
	})

	t.Run("ok/insert-order", func(t *testing.T) {
		t.Parallel()
		m := typex.NewOrderedMap[string, int]()

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
		m := typex.NewOrderedMap[string, int]()
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
		m := typex.NewOrderedMap[string, int]()
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
		m := typex.NewOrderedMap[string, int]()
		m.Put("a", 1)
		m.Put("c", 3)

		// act
		typex.Delete(m, "b")

		// assert
		_, okB := m.Get("b")
		assert.False(t, okB)

		vals := m.GetValues()
		assert.Equal(t, []int{1, 3}, vals)
	})

	t.Run("edge/nil-dict", func(t *testing.T) {
		t.Parallel()
		var m typex.OrderedMap[int, int] // nil

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
		m := typex.NewOrderedMap[string, int]()
		m.Put("a", 1)
		m.Put("b", 2)
		m.Put("c", 3)

		// act
		typex.Delete(m, "b")

		// assert
		_, okB := m.Get("b")
		assert.False(t, okB)

		vals := m.GetValues()
		assert.Equal(t, []int{1, 3}, vals)
	})

	t.Run("ok/delete-not-existing", func(t *testing.T) {
		t.Parallel()
		m := typex.NewOrderedMap[string, int]()
		m.Put("a", 1)
		m.Put("c", 3)

		// act
		typex.Delete(m, "b")

		// assert
		_, okB := m.Get("b")
		assert.False(t, okB)

		vals := m.GetValues()
		assert.Equal(t, []int{1, 3}, vals)
	})

	t.Run("ok/nil", func(t *testing.T) {
		t.Parallel()
		var m *typex.OrderedMap[int, int] // nil

		// act: should't panic
		typex.Delete(m, 1)

		// assert
		assert.Nil(t, m)
	})

	t.Run("ok/empty", func(t *testing.T) {
		t.Parallel()
		m := typex.NewOrderedMap[int, int]()

		// act: should't panic
		typex.Delete(m, 1)

		// assert: the struct still empty
		assert.Equal(t, 0, len(m.GetValues()))
	})
}

func Test_OrderedMap_GetValues(t *testing.T) {
	t.Parallel()

	t.Run("edge/empty", func(t *testing.T) {
		t.Parallel()
		var m typex.OrderedMap[string, int]

		// act
		vals := m.GetValues()

		// assert
		assert.NotNil(t, vals)
		assert.Equal(t, 0, len(vals))
	})

	t.Run("ok/single", func(t *testing.T) {
		t.Parallel()
		m := typex.NewOrderedMap[string, int]()
		m.Put("a", 7)

		// act
		vals := m.GetValues()

		// assert
		assert.Equal(t, 1, len(vals))
		assert.Equal(t, []int{7}, vals)
	})

	t.Run("ok/multiple", func(t *testing.T) {
		t.Parallel()
		m := typex.NewOrderedMap[int, string]()
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
		m := typex.NewOrderedMap[string, int]()

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
