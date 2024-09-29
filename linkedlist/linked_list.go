package linkedlist

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	data T
	next *Node[T]
}

type LinkedList[T constraints.Ordered] struct {
	Head *Node[T]
}

func (l *LinkedList[T]) CreateLinkedList(arr []T) {
	var tail *Node[T]

	for _, val := range arr {
		newNode := &Node[T]{
			data: val,
			next: nil,
		}

		if l.Head == nil {
			l.Head = newNode
			tail = newNode
		} else {
			tail.next = newNode
			tail = newNode
		}
	}
}

func (l *LinkedList[T]) PrintList() {
	current := l.Head

	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}
