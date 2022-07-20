package leetcode

import "strconv"

// My Solution: convert int to string
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := strconv.Itoa(x)
	l, r := 0, len(y)-1
	for l < r {
		if y[l] == y[r] {
			l++
			r--
		} else {
			return false
		}
	}

	return true
}
