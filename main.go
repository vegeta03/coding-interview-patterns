package main

import (
	"fmt"
	"strings"

	fastslowpointers "github.com/vegeta03/coding-interview-patterns/fast-and-slow-pointers"
	modifiedbinarysearch "github.com/vegeta03/coding-interview-patterns/modified-binary-search"
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

	fmt.Printf("Binary Search START\n\n")
	numsLists = [][]int{
		{},
		{0, 1},
		{1, 2, 3},
		{-1, 0, 3, 5, 9, 12},
		{-100, -67, -55, -50, -49, -40, -33, -22, -10, -5},
	}

	targetList := []int{12, 1, 3, 9, -22}
	for i := 0; i < len(numsLists); i++ {
		nums := numsLists[i]
		target := targetList[i]
		index := modifiedbinarysearch.BinarySearch(nums, target)
		fmt.Printf("%d.\tArray to search: %v\n", i+1, strings.Replace(fmt.Sprint(nums), " ", ", ", -1))

		fmt.Printf("\tTarget: %d\n", target)
		if index != -1 {
			fmt.Printf("\t%d exists in the slice at index %d\n", target, index)
		} else {
			fmt.Printf("\t%d does not exist in the slice, so the return value is %d\n", target, index)
		}
		fmt.Println(strings.Repeat("-", 100))
	}
	fmt.Println("Binary Search END")
	fmt.Printf("%s\n\n", strings.Repeat("*", 100))

	fmt.Printf(" Iterative Binary Search Rotated START\n\n")
	numsList := [][]int{
		{5, 6, 7, 1, 2, 3, 4},
		{40, 50, 60, 10, 20, 30},
		{47, 58, 69, 72, 83, 94, 12, 24, 35},
		{77, 82, 99, 105, 5, 13, 28, 41, 56, 63},
		{48, 52, 57, 62, 68, 72, 5, 7, 12, 17, 21, 28, 33, 37, 41},
	}
	targetList = []int{1, 50, 12, 56, 5}

	for index, value := range numsList {
		fmt.Printf("%d.\tRotated array: %s\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		fmt.Printf("\tTarget %d found at index %d\n", targetList[index], modifiedbinarysearch.IterativeBinarySearchRotated(value, targetList[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
	fmt.Println("Iterative Binary Search Rotated END")
	fmt.Printf("%s\n\n", strings.Repeat("*", 100))

	fmt.Printf(" START\n\n")
	fmt.Println(" END")
	fmt.Printf("%s\n\n", strings.Repeat("*", 100))
}
