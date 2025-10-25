package internal_test

import (
	"testing"

	"github.com/lif0/pkg/utils/internal"
	"github.com/stretchr/testify/assert"
)

func collect[T any](l *internal.LinkedList[T]) []T {
	var out []T
	for _, v := range l.Iter() {
		out = append(out, v)
	}
	return out
}

func Test_LinkedList_Append(t *testing.T) {
	t.Parallel()

	t.Run("ok/append-to-empty", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]

		// act
		l.Append(&internal.Node[int]{Val: 1})

		// assert
		assert.Equal(t, 1, l.Len())
		head := l.GetHead()
		assert.NotNil(t, head)
		assert.Equal(t, 1, head.Val)
		assert.Nil(t, head.Prev)
		assert.Nil(t, head.Next)
	})

	t.Run("ok/append-to-non-empty", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		n1 := &internal.Node[int]{Val: 1}
		n2 := &internal.Node[int]{Val: 2}

		// act
		l.Append(n1)
		l.Append(n2)

		// assert
		assert.Equal(t, 2, l.Len())
		head := l.GetHead()
		assert.Equal(t, 1, head.Val)
		assert.Equal(t, 2, head.Next.Val)
		assert.Equal(t, head, head.Next.Prev)
	})

	t.Run("bug/append-preserves-external-next", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		externalTail := &internal.Node[int]{Val: 9}
		n := &internal.Node[int]{Val: 1, Next: externalTail}

		// act
		l.Append(n)

		// assert
		assert.Equal(t, 1, l.Len())
		got := collect(&l)
		// Итерация “видит” 2 узла, хотя Len()==1 — следствие того, что Append не сбрасывает node.Next
		assert.Equal(t, []int{1, 9}, got)
	})
}

func Test_LinkedList_Remove(t *testing.T) {
	t.Parallel()

	t.Run("ok/remove-head", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		n1 := &internal.Node[int]{Val: 1}
		n2 := &internal.Node[int]{Val: 2}
		n3 := &internal.Node[int]{Val: 3}
		l.Append(n1)
		l.Append(n2)
		l.Append(n3)

		// act
		l.Remove(n1)

		// assert
		assert.Equal(t, 2, l.Len())
		head := l.GetHead()
		assert.Equal(t, 2, head.Val)
		assert.Nil(t, head.Prev)
		assert.Equal(t, 3, head.Next.Val)
	})

	t.Run("ok/remove-tail", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		n1 := &internal.Node[int]{Val: 1}
		n2 := &internal.Node[int]{Val: 2}
		l.Append(n1)
		l.Append(n2)

		// act
		l.Remove(n2)

		// assert
		assert.Equal(t, 1, l.Len())
		head := l.GetHead()
		assert.Equal(t, 1, head.Val)
		assert.Nil(t, head.Next)
	})

	t.Run("ok/remove-middle", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		n1 := &internal.Node[int]{Val: 1}
		n2 := &internal.Node[int]{Val: 2}
		n3 := &internal.Node[int]{Val: 3}
		l.Append(n1)
		l.Append(n2)
		l.Append(n3)

		// act
		l.Remove(n2)

		// assert
		assert.Equal(t, 2, l.Len())
		head := l.GetHead()
		assert.Equal(t, 1, head.Val)
		assert.Equal(t, 3, head.Next.Val)
		assert.Equal(t, head, head.Next.Prev)
	})

	t.Run("ok/remove-singleton", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		n1 := &internal.Node[int]{Val: 1}
		l.Append(n1)

		// act
		l.Remove(n1)

		// assert
		assert.Equal(t, 0, l.Len())
		assert.Nil(t, l.GetHead())
	})

	t.Run("ok/remove-nil-noop", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		l.Append(&internal.Node[int]{Val: 1})

		// act
		l.Remove(nil)

		// assert
		assert.Equal(t, 1, l.Len())
		assert.Equal(t, []int{1}, collect(&l))
	})

	t.Run("bug/remove-foreign-node-decrements-size", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		l.Append(&internal.Node[int]{Val: 1})
		foreign := &internal.Node[int]{Val: 999} // не в списке

		// act
		l.Remove(foreign)

		// assert
		// Структура связей не изменилась (узел 1 всё ещё голова), но размер уменьшился с 1 до 0
		assert.Equal(t, []int{1}, collect(&l))
		assert.Equal(t, 0, l.Len())
	})

	t.Run("bug/remove-twice-size-negative", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		n := &internal.Node[int]{Val: 1}
		l.Append(n)

		// act
		l.Remove(n)
		l.Remove(n) // повторно по тому же узлу

		// assert
		assert.Equal(t, -1, l.Len())
		assert.Nil(t, l.GetHead())
	})
}

func Test_LinkedList_GetHead(t *testing.T) {
	t.Parallel()

	t.Run("edge/empty", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]

		// act
		h := l.GetHead()

		// assert
		assert.Nil(t, h)
	})

	t.Run("ok/non-empty", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		n1 := &internal.Node[int]{Val: 1}
		n2 := &internal.Node[int]{Val: 2}
		l.Append(n1)
		l.Append(n2)

		// act
		h := l.GetHead()

		// assert
		assert.NotNil(t, h)
		assert.Equal(t, 1, h.Val)
	})
}

func Test_LinkedList_Len(t *testing.T) {
	t.Parallel()

	t.Run("ok/increments-and-decrements", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		n1 := &internal.Node[int]{Val: 1}
		n2 := &internal.Node[int]{Val: 2}
		n3 := &internal.Node[int]{Val: 3}

		// act
		l.Append(n1)
		l.Append(n2)
		l.Append(n3)
		l.Remove(n2)

		// assert
		assert.Equal(t, 2, l.Len())
		assert.Equal(t, []int{1, 3}, collect(&l))
	})
}

func Test_LinkedList_Iter(t *testing.T) {
	t.Parallel()

	t.Run("edge/empty", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		iter := l.Iter()

		// arrange
		calls := 0

		// act
		iter(func(i int, v int) bool {
			calls++
			return true
		})

		// assert
		assert.Equal(t, 0, calls)
	})

	t.Run("ok/full-iteration", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[string]
		l.Append(&internal.Node[string]{Val: "a"})
		l.Append(&internal.Node[string]{Val: "b"})
		l.Append(&internal.Node[string]{Val: "c"})
		iter := l.Iter()

		// arrange
		var gotIdx []int
		var gotVal []string

		// act
		iter(func(i int, v string) bool {
			gotIdx = append(gotIdx, i)
			gotVal = append(gotVal, v)
			return true
		})

		// assert
		assert.Equal(t, []int{0, 1, 2}, gotIdx)
		assert.Equal(t, []string{"a", "b", "c"}, gotVal)
	})

	t.Run("ok/stop-immediately", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		l.Append(&internal.Node[int]{Val: 10})
		l.Append(&internal.Node[int]{Val: 20})
		l.Append(&internal.Node[int]{Val: 30})
		iter := l.Iter()

		// arrange
		var gotIdx []int
		var gotVal []int
		calls := 0

		// act
		iter(func(i int, v int) bool {
			calls++
			gotIdx = append(gotIdx, i)
			gotVal = append(gotVal, v)
			return false
		})

		// assert
		assert.Equal(t, 1, calls)
		assert.Equal(t, []int{0}, gotIdx)
		assert.Equal(t, []int{10}, gotVal)
	})

	t.Run("ok/stop-middle", func(t *testing.T) {
		t.Parallel()
		var l internal.LinkedList[int]
		l.Append(&internal.Node[int]{Val: 1})
		l.Append(&internal.Node[int]{Val: 2})
		l.Append(&internal.Node[int]{Val: 3})
		l.Append(&internal.Node[int]{Val: 4})
		iter := l.Iter()

		// arrange
		var gotIdx []int
		var gotVal []int
		calls := 0

		// act
		iter(func(i int, v int) bool {
			calls++
			gotIdx = append(gotIdx, i)
			gotVal = append(gotVal, v)
			return i < 1
		})

		// assert
		assert.Equal(t, 2, calls)
		assert.Equal(t, []int{0, 1}, gotIdx)
		assert.Equal(t, []int{1, 2}, gotVal)
	})
}
