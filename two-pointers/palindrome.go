package twopointers

func IsPalindrome(inputString string) bool {
	left := 0
	right := len(inputString) - 1
	for left < right {
		if inputString[left] != inputString[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}
