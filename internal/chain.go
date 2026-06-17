package internal

type Chain[T any] struct {
	size int
	head *ChainLink[T]
	tail *ChainLink[T]
}

type ChainLink[T any] struct {
	Val  T
	Prev *ChainLink[T]
	Next *ChainLink[T]
}

// Remove ...
// time: O(1); mem: O(1)
func (c *Chain[T]) Remove(node *ChainLink[T]) {
	if node == nil {
		return
	}

	// If a previous node exists, link it to the next one, skipping the current node.
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	//  If a next node exists, set its previous pointer to the current node’s previous.
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	// if the node and the head is equal, set to head node's next.
	if node == c.head {
		c.head = c.head.Next
	}

	// if the node and the tail is equal, set to tail node's previous.
	if node == c.tail {
		c.tail = c.tail.Prev
	}

	c.size -= 1
}

// Append ...
// time: O(1); mem: O(1)
func (c *Chain[T]) Append(node *ChainLink[T]) {
	if c.tail == nil {
		c.head = node
		c.tail = node
	} else {
		c.tail.Next = node
		node.Prev = c.tail
		c.tail = node
	}

	c.size += 1
}

// GetHead ...
// time: O(1); mem: O(1)
func (c *Chain[T]) GetHead() *ChainLink[T] {
	return c.head
}

// Len ...
// time: O(1); mem: O(1)
func (c *Chain[T]) Len() int {
	return c.size
}

// Iter iteration on chain
//
// Example:
//
//	m := Chain[int, string]()
//	for i, v := range m.Iter() {
//		fmt.Println(i,v)
//	}
func (c *Chain[T]) Iter() func(func(int, T) bool) {
	return func(yield func(int, T) bool) {
		i := 0
		for n := c.head; n != nil; n = n.Next {
			if !yield(i, n.Val) {
				return
			}
			i++
		}
	}
}
