//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"strings"

	modifiedbinarysearch "github.com/vegeta03/coding-interview-patterns/modified-binary-search"
)

func main() {
	numsList := [][]int{
		{5, 6, 7, 1, 2, 3, 4},
		{40, 50, 60, 10, 20, 30},
		{47, 58, 69, 72, 83, 94, 12, 24, 35},
		{77, 82, 99, 105, 5, 13, 28, 41, 56, 63},
		{48, 52, 57, 62, 68, 72, 5, 7, 12, 17, 21, 28, 33, 37, 41},
	}
	targetList := []int{1, 50, 12, 56, 5}

	for index, value := range numsList {
		fmt.Printf("%d.\tRotated array: %s\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		fmt.Printf("\tTarget %d found at index %d\n", targetList[index], modifiedbinarysearch.IterativeBinarySearchRotated(value, targetList[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
