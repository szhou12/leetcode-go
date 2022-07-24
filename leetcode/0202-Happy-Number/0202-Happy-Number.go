package leetcode

func isHappy(n int) bool {
	// base cases
	if n == 1 {
		return true
	}
	if n < 10 && n != 1 {
		return false
	}

	newNum := 0

	for n > 0 {
		digit := n % 10
		newNum += square(digit)
		n /= 10
	}

	return isHappy(newNum)
}

func square(n int) int {
	return n * n
}
