package leetcode

func isPowerOfThree(n int) bool {
	for n > 1 {
		if n%3 == 0 {
			n /= 3
		} else {
			return false
		}
	}

	return n == 1
}
