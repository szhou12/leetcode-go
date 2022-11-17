package leetcode

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0

	res := 0
	for right < len(s) {
		rightElement := s[right]
		right++
		window[rightElement]++

		for window[rightElement] > 1 {
			leftElement := s[left]
			left++
			window[leftElement]--
		}

		res = max(res, right-left)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
