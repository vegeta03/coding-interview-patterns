package linkedlist

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func CreateLinkedList(arr []int) *LinkedList {
	linkedList := &LinkedList{}

	for _, val := range arr {
		node := &Node{Data: val}

		if linkedList.Head == nil {
			linkedList.Head = node
		} else {
			current := linkedList.Head

			for current.Next != nil {
				current = current.Next
			}

			current.Next = node
		}
	}
	return linkedList
}

func (l *LinkedList) PrintList() {
	current := l.Head

	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}
