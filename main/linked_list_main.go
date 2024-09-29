//go:build ignore
// +build ignore

package main

import "github.com/vegeta03/coding-interview-patterns/linkedlist"

func main() {
	arr := []int{1, 2, 3, 4}
	var ll linkedlist.LinkedList[int]
	ll.CreateLinkedList(arr)
	ll.PrintList()
}
