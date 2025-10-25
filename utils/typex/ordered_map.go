package typex

import (
	"github.com/lif0/pkg/utils/internal"
)

type OrderedMap[K comparable, V any] struct {
	dict map[K]*internal.Node[V]
	list internal.LinkedList[V]
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		dict: make(map[K]*internal.Node[V]),
		list: internal.LinkedList[V]{},
	}
}

// Get ...
// time: O(1); mem: O(1)
func (this *OrderedMap[K, V]) Get(key K) (V, bool) {
	if node, ok := this.dict[key]; ok {
		return node.Val, true
	}

	var zeroVal V
	return zeroVal, false
}

// Put ...
// time: O(1); mem: O(1)
func (this *OrderedMap[K, V]) Put(key K, value V) {
	if node, ok := this.dict[key]; ok {
		// this.removeNode(node)
		// node.Val = value
		// this.addNodeToTail(node)

		node.Val = value
	} else {
		node = &internal.Node[V]{Val: value}
		this.list.Append(node)
		this.dict[key] = node
	}
}

// Delete ...
// time: O(1); mem: O(1)
func (this *OrderedMap[K, V]) Delete(key K) {
	if node, ok := this.dict[key]; ok {
		this.list.Remove(node)
		delete(this.dict, key)
	}
}

// GetValues ...
// time: O(N); mem: O(N)
func (this *OrderedMap[K, V]) GetValues() []V {
	result := make([]V, this.list.Len())

	if cap(result) == 0 {
		return result
	}

	if cap(result) == 1 {
		result[0] = this.list.GetHead().Val
	}

	for i, v := range this.list.Iter() {
		result[i] = v
	}

	return result
}

// Delete built-in function deletes the element with the specified key
// (m[key]) from the OrderedMap. If m is nil or there is no such element, delete
// is a no-op.
func Delete[Type comparable, Type1 any](m *OrderedMap[Type, Type1], key Type) {
	if m == nil {
		return
	}

	if m.list.Len() == 0 {
		return
	}

	if node, ok := m.dict[key]; ok {
		m.list.Remove(node)
		delete(m.dict, key)
	}
}
