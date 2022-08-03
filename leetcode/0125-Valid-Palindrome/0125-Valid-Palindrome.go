package leetcode

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	l := 0
	r := len(s) - 1

	for l < r {
		// skip spaces
		for l < r && !isChar(s[l]) {
			l++
		}
		for l < r && !isChar(s[r]) {
			r--
		}
		// check if palindrome
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isChar(char byte) bool {
	if (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') {
		return true
	}

	return false
}
