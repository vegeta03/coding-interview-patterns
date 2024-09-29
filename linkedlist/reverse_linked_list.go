package linkedlist

func (l *LinkedList[T]) Reverse() {
	var prev *Node[T] = nil
	current := l.Head
	var next *Node[T] = nil

	for current != nil {
		next = current.next

		current.next = prev

		prev = current
		current = next
	}
	l.Head = prev
}
