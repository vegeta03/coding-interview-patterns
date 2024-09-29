//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	data T
	next *Node[T]
}

type LinkedList[T constraints.Ordered] struct {
	head *Node[T]
}

func (l *LinkedList[T]) CreateLinkedList(arr []T) {
	var tail *Node[T]
	for _, e := range arr {
		newNode := &Node[T]{
			data: e,
			next: nil,
		}

		if l.head == nil {
			l.head = newNode
			tail = newNode
		} else {
			tail.next = newNode
			tail = newNode
		}
	}
}

func (l *LinkedList[T]) PrintList() {
	current := l.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (l *LinkedList[T]) Reverse() {
	var prev *Node[T]
	current := l.head
	var next *Node[T]

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

func main() {
	arr := []int{1, 2, 3}

	var linkedList LinkedList[int]

	linkedList.CreateLinkedList(arr)
	linkedList.PrintList()
}
