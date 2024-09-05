package fastslowpointers

func pow(digit int, power int) int {
	res := 1

	for i := 0; i < power; i++ {
		res = res * digit
	}

	return res
}

func sumOfSquaredDigits(number int) int {
	totalSum := 0

	for number > 0 {
		digit := number % 10
		number = number / 10

		totalSum += pow(digit, 2)
	}

	return totalSum
}

func IsHappy(num int) bool {
	slowPointer := num
	fastPointer := sumOfSquaredDigits(num)

	for fastPointer != 1 && slowPointer != fastPointer {
		slowPointer = sumOfSquaredDigits(slowPointer)
		fastPointer = sumOfSquaredDigits(sumOfSquaredDigits(fastPointer))
	}

	return fastPointer == 1
}
