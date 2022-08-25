package leetcode

func isSubsequence(s string, t string) bool {
	// edge cases
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	l := 0

	for r := 0; r < len(t); r++ {
		if l == len(s) {
			return true
		}
		if t[r] == s[l] {
			l++
		}
	}

	return l == len(s)
}

// Same solution - better coding style
func isSubsequence_standard(s string, t string) bool {
	for len(s) > 0 && len(t) > 0 {
		if s[0] == t[0] {
			s = s[1:]
		}
		t = t[1:]
	}
	return len(s) == 0
}
