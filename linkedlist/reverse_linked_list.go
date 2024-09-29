package linkedlist

func (l *LinkedList[T]) Reverse() {
	var prev *Node[T] = nil
	var next *Node[T] = nil
	current := l.Head

	for current != nil {
		next = current.next

		current.next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
