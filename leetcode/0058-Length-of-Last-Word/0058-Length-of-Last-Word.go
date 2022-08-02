package leetcode

// start: empty space right before last word
// end: last char of last word
func lengthOfLastWord(s string) int {
	end := len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}

	// edge case: no word found
	if end < 0 {
		return 0
	}

	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}

	return end - start
}
