package linkedlist

func (l *LinkedList) Reverse() {
	var prev *Node = nil
	var next *Node = nil
	current := l.Head

	for current != nil {
		next = current.Next

		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
