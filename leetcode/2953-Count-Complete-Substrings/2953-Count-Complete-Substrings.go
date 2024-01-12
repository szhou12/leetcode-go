package leetcode

// A complete substring
// Condition 1: Each character in s occurs exactly k times.
// Condition 2: For any two adjacent characters c1 and c2 in s, the absolute difference in their positions in the alphabet is at most 2
func countCompleteSubstrings(word string, k int) int {
	n := len(word)

	res := 0
	// Flex sliding window on Condition 2
	for left := 0; left < n; { // 注意这种loop的写法! 每一iter, left直接跳到right
		right := left + 1
		for right < n && abs(int(word[right])-int(word[right-1])) <= 2 {
			right++
		}
		// valid section = [left, right-1]
		res += charOccurs(word[left:right], k)
		left = right
	}

	return res
}

// count # of substrings in s that meets Condition 1
func charOccurs(s string, k int) int {
	count := 0
	// get unique chars
	set := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		set[s[i]] = true
	}
	// iterate total kinds of unique chars
	// if T unique chars, window size = T*k to meet Condition 1
	for T := 1; T <= len(set); T++ { // O(26)
		windowSize := T * k
		freq := make([]int, 26)

		// Fixed sliding window
		r := 0
		for l := 0; l+windowSize-1 < len(s); l++ { // 注意这种loop的写法! 结束条件是右边界 (l+windowSize-1) 不超过len(s)
			// 吃
			for r <= l+windowSize-1 {
				freq[int(s[r] - 'a')]++
				r++
			}
			// valid section = [l, r-1]
			if isValid(freq, k) {
				count++
			}
			// 吐
			freq[int(s[l] - 'a')]--
		}
	}
	return count

}

// check if each existing char has eactly k occurrences
// O(26)
func isValid(freq []int, k int) bool {
	for _, f := range freq {
		// if there exists a char that doesn't have k occurrences, failed
		if f != 0 && f != k {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
