package leetcode

// Method 1: Divide and Conquer
func longestSubstring(s string, k int) int {
	return findLongest(s, 0, len(s), k)
}

func findLongest(s string, start int, end int, k int) int {
	// base case: if length of (end - start) < k, already cannot satisfy condition, return 0
	if end-start < k {
		return 0
	}

	// create freq map for each character between s[start : end]
	count := make([]int, 26)
	for i := start; i < end; i++ { // O(n)
		idx := int(s[i] - 'a')
		count[idx]++
	}

	// find spliter char in s[start : end]
	for i := 0; i < 26; i++ { // O(26)
		if 0 < count[i] && count[i] < k {
			for j := start; j < end; j++ { // O(n)
				if int(s[j]) == i+'a' {
					left := findLongest(s, start, j, k)  // O(n)
					right := findLongest(s, j+1, end, k) // O(n)
					return max(left, right)
				}
			}
		}
	}

	// if jump out of loop, it means all chars >= k
	return end - start
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Method 2: Sliding Window
