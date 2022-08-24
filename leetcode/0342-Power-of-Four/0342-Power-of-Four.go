package leetcode

func isPowerOfFour(n int) bool {
	for n >= 4 {
		if n%4 == 0 {
			n /= 4
		} else {
			return false
		}
	}

	return n == 1
}
