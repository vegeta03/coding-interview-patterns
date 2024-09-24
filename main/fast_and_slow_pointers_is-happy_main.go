//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"strings"

	fastslowpointers "github.com/vegeta03/coding-interview-patterns/fast-and-slow-pointers"
)

func main() {
	array := []int{1, 5, 19, 25, 7}
	for i, v := range array {
		fmt.Printf("%d.\tInput number: %d\n", i+1, v)
		fmt.Printf("\n\tIs it a happy number? %t\n", fastslowpointers.IsHappy(v))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
