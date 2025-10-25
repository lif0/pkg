package internal

type LinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

type Node[T any] struct {
	Val  T
	Prev *Node[T]
	Next *Node[T]
}

// Remove ...
// time: O(1); mem: O(1)
func (l *LinkedList[T]) Remove(node *Node[T]) {
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
	if node == l.head {
		l.head = l.head.Next
	}

	// if the node and the tail is equal, set to tail node's previous.
	if node == l.tail {
		l.tail = l.tail.Prev
	}

	l.size -= 1
}

// Append ...
// time: O(1); mem: O(1)
func (l *LinkedList[T]) Append(node *Node[T]) {
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		node.Prev = l.tail
		l.tail = node
	}

	l.size += 1
}

// GetHead ...
// time: O(1); mem: O(1)
func (l *LinkedList[T]) GetHead() *Node[T] {
	return l.head
}

// Len ...
// time: O(1); mem: O(1)
func (l *LinkedList[T]) Len() int {
	return l.size
}

// Iter ...
func (l *LinkedList[T]) Iter() func(func(int, T) bool) {
	return func(yield func(int, T) bool) {
		i := 0
		for n := l.head; n != nil; n = n.Next {
			if !yield(i, n.Val) {
				return
			}
			i++
		}
	}
}
