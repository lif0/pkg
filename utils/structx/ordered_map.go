package structx

import (
	"github.com/lif0/pkg/utils/internal"
)

// OrderedMap is a map[Type]Type1-like collection that preserves the order
// in which keys were inserted. It behaves like a regular map but
// allows deterministic iteration over its elements.
//
// OrderedMap is useful when both quick key-based access and
// predictable iteration order are desired.
type OrderedMap[K comparable, V any] struct {
	dict map[K]*internal.Node[V]
	list internal.LinkedList[V]
}

// NewOrderedMap returns a new empty OrderedMap.
func NewOrderedMap[K comparable, V any](cap ...int) *OrderedMap[K, V] {
	var dictCap int
	if len(cap) > 0 {
		dictCap = cap[0]
	}

	return &OrderedMap[K, V]{
		dict: make(map[K]*internal.Node[V], dictCap),
		list: internal.LinkedList[V]{},
	}
}

// Get retrieves the value stored under the given key.
// The second return value reports whether the key was present.
//
// Complexity:
// - time: O(1)
// - mem: O(1)
func (this *OrderedMap[K, V]) Get(key K) (V, bool) {
	if node, ok := this.dict[key]; ok {
		return node.Val, true
	}

	var zeroVal V
	return zeroVal, false
}

// Put sets the value for the given key.
// If the key already exists, its value is updated.
// Otherwise, a new entry is added to the end of the order.
//
// Complexity:
// - time: O(1)
// - mem: O(1)
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

// Delete removes the element with the specified key.
// If the key does not exist, Delete does nothing.
//
// Complexity:
// - time: O(1)
// - mem: O(1)
func (this *OrderedMap[K, V]) Delete(key K) {
	if node, ok := this.dict[key]; ok {
		this.list.Remove(node)
		delete(this.dict, key)
	}
}

// GetValues returns all values in insertion order.
// The returned slice has the same length as the number of elements.
//
// Complexity:
// - time: O(N)
// - mem: O(N)
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

// Iter iteration on map
//
// Example:
//
//	m := NewOrderedMap[int, string]()
//	for i, v := range m.Iter() {
//		fmt.Println(i,v)
//	}
func (this *OrderedMap[K, V]) Iter() func(func(int, V) bool) {
	return this.list.Iter()
}

// Delete built-in function deletes the element with the specified key
// (m[key]) from the OrderedMap. If m is nil or there is no such element, delete
// is a no-op.
//
// Example:
//
//	var om = NewOrderedMap[string, int]()
//	om.Put("x", 1)
//	structx.Delete(om, "x")
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
