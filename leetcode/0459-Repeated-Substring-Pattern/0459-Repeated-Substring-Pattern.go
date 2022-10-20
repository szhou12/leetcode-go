package leetcode

// My Solution
func repeatedSubstringPattern(s string) bool {
	// Edge case
	if len(s) == 2 && s[0] == s[1] {
		return true
	}

	for right := 1; right < len(s)-1; right++ {
		substr := s[:right]
		if len(s)%len(substr) == 0 && canForm(s, substr) {
			return true
		}
	}

	return false
}

func canForm(s string, substr string) bool {
	n := len(substr)
	l := 0
	for l < len(s) {
		if s[l:l+n] != substr {
			return false
		}
		l += n
	}
	return true
}
