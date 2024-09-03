package twopointers

import "sort"

func FindSumOfThree(nums []int, target int) bool {
	sort.Ints(nums)

	var low int
	var high int
	var triple int

	for i := 0; i < len(nums)-2; i++ {
		low = i + 1
		high = len(nums) - 1

		for low < high {
			triple = nums[1] + nums[low] + nums[high]

			if triple == target {
				return true
			} else if triple > target {
				high -= 1
			} else {
				low += 1
			}
		}
	}
	return false
}
