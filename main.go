package main

import (
	"fmt"
	"strings"

	fastslowpointers "github.com/vegeta03/coding-interview-patterns/fast-and-slow-pointers"
	twopointers "github.com/vegeta03/coding-interview-patterns/two-pointers"
)

func main() {

	fmt.Printf("Is Palindrome START\n\n")
	str := []string{"RACEACAR", "A", "ABCDEFGFEDCBA", "ABC", "ABCBA", "ABBA", "RACEACAR"}
	for i, s := range str {
		fmt.Printf("Test Case # %d\n", i+1)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
		fmt.Printf("The input string is '%s' and the length of the string is %d.\n", s, len(s))
		fmt.Printf("\nIs it a palindrome?.....%v\n", twopointers.IsPalindrome(s))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
	fmt.Println("Is Palindrome END")
	fmt.Printf("%s\n\n", strings.Repeat("*", 100))

	fmt.Printf("SumOfThree START\n\n")
	numsLists := [][]int{
		{3, 7, 1, 2, 8, 4, 5},
		{-1, 2, 1, -4, 5, -3},
		{2, 3, 4, 1, 7, 9},
		{1, -1, 0},
		{2, 4, 2, 7, 6, 3, 1},
	}
	tList := []int{10, 7, 20, -1, 8}
	for i, nList := range numsLists {
		fmt.Printf("%d. Input array: %s\n", i+1, strings.Replace(fmt.Sprint(nList), " ", ", ", -1))

		if twopointers.FindSumOfThree(nList, tList[i]) {
			fmt.Printf("   Sum for %d exists\n", tList[i])
		} else {
			fmt.Printf("   Sum for %d does not exist\n", tList[i])
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
	fmt.Println("SumOfThree END")
	fmt.Printf("%s\n\n", strings.Repeat("*", 100))

	fmt.Printf("Happy number START\n\n")
	array := []int{1, 5, 19, 25, 7}
	for i, v := range array {
		fmt.Printf("%d.\tInput number: %d\n", i+1, v)
		fmt.Printf("\n\tIs it a happy number? %t\n", fastslowpointers.IsHappy(v))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
	fmt.Println("Happy number END")
	fmt.Printf("%s\n\n", strings.Repeat("*", 100))
}
