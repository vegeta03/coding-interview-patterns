//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"strings"

	twopointers "github.com/vegeta03/coding-interview-patterns/two-pointers"
)

func main() {

	fmt.Printf("Is Palindrome START\n\n")
	str := []string{"RACEACAR", "A", "ABCDEFGFEDCBA", "ABC", "ABCBA", "ABBA", "RACEACAR"}
	for i, s := range str {
		fmt.Printf("\nTest Case # %d\n", i+1)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
		fmt.Printf("The input string is '%s' and the length of the string is %d.\n", s, len(s))
		fmt.Printf("\nIs it a palindrome?.....%v\n", twopointers.IsPalindrome(s))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
	fmt.Println("Is Palindrome END")
	fmt.Printf("%s\n\n", strings.Repeat("*", 100))

}
