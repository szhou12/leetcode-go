package leetcode

import "strings"

// The 'shift' operation can be regarded as a 'sliding' operation on A+A
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	A := s + s
	return strings.Contains(A, goal)
}
