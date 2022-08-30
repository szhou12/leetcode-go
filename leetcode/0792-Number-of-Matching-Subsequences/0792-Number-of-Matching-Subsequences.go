package leetcode

// Solution - Brute force
// Use 0392 as helper function
func numMatchingSubseq_BF(s string, words []string) int {
	count := 0
	for _, t := range words {
		if isSubseq(s, t) {
			count++
		}
	}

	return count
}

// check if t is subseq of s
func isSubseq(s string, t string) bool {
	// edge cases
	// if s empty, t empty => true
	// if s empty, t not => false
	// if s not, t empty => true
	if len(t) == 0 {
		return true
	}
	if len(s) == 0 {
		return false
	}

	// pointer for t
	left := 0
	for right := 0; right < len(s); right++ {
		if len(t) == left {
			return true
		}
		if t[left] == s[right] {
			left++
		}
	}

	return len(t) == left
}
