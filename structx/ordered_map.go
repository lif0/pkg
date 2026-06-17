package structx

import (
	"github.com/lif0/pkg/structx/internal"
)

type kv[K any, V any] struct {
	K K
	V V
}

// OrderedMap is a map[Type]Type1-like collection that preserves the order
// in which keys were inserted. It behaves like a regular map but
// allows deterministic iteration over its elements.
//
// OrderedMap is useful when both quick key-based access and
// predictable iteration order are desired.
type OrderedMap[K comparable, V any] struct {
	dict    map[K]*internal.ChainLink[kv[K, V]]
	list    internal.Chain[kv[K, V]]
	objPool *ObjectPool[internal.ChainLink[kv[K, V]]]
}

// NewOrderedMap returns a new empty OrderedMap.
func NewOrderedMap[K comparable, V any](size ...uint32) *OrderedMap[K, V] {
	var capacity uint32
	if len(size) > 0 && size[0] > 0 {
		capacity = size[0]
	}

	return &OrderedMap[K, V]{
		dict:    make(map[K]*internal.ChainLink[kv[K, V]], capacity),
		list:    internal.Chain[kv[K, V]]{},
		objPool: NewObjectPool[internal.ChainLink[kv[K, V]]](capacity),
	}
}

// Get retrieves the value stored under the given key.
// The second return value reports whether the key was present.
//
// Complexity:
// - time: O(1)
// - mem: O(1)
func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	if node, ok := m.dict[key]; ok {
		return node.Val.V, true
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
func (m *OrderedMap[K, V]) Put(key K, value V) {
	if node, ok := m.dict[key]; ok {
		node.Val.V = value
	} else {
		node = m.objPool.Get() // &internal.Node[kv[K,V]]{Val: value}
		node.Val.K = key
		node.Val.V = value
		node.Prev = nil // overcautiousness
		node.Next = nil // overcautiousness

		m.list.Append(node)
		m.dict[key] = node
	}
}

// Delete removes the element with the specified key.
// If the key does not exist, Delete does nothing.
//
// Complexity:
// - time: O(1)
// - mem: O(1)
func (m *OrderedMap[K, V]) Delete(key K) {
	Delete(m, key)
}

// GetValues returns all values in insertion order.
// The returned slice has the same length as the number of elements.
//
// Complexity:
// - time: O(N)
// - mem: O(N)
func (m *OrderedMap[K, V]) GetValues() []V {
	result := make([]V, m.list.Len())

	if cap(result) == 0 {
		return result
	}

	if cap(result) == 1 {
		result[0] = m.list.GetHead().Val.V
	}

	for i, v := range m.list.Iter() {
		result[i] = v.V
	}

	return result
}

// Iter iteration on map in insertion order
//
// Example:
//
//	m := NewOrderedMap[int, string]()
//
//	for k, v := range m.Iter() {
//		fmt.Println(k,v)
//	}
func (m *OrderedMap[K, V]) Iter() func(func(K, V) bool) {
	return func(yield func(K, V) bool) {
		h := m.list.GetHead()
		for n := h; n != nil; n = n.Next {
			if !yield(n.Val.K, n.Val.V) {
				return
			}
		}
	}
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
