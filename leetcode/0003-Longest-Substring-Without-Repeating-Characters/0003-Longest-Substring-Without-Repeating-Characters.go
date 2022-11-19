package leetcode

// 写法一: 模版Flex
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]bool)
	right := 0
	res := 0

	for left := 0; left < len(s); left++ {
		for right < len(s) && !window[s[right]] {
			window[s[right]] = true
			right++
		}
		res = max(res, right-left)

		window[s[left]] = false

	}
	return res
}

// 写法二: 模版Fix
func lengthOfLongestSubstring_fix(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0

	res := 0
	for right < len(s) {
		rightElement := s[right] // 吃
		right++
		window[rightElement]++

		for window[rightElement] > 1 {
			leftElement := s[left] // 吐
			left++
			window[leftElement]--
		}

		// update result
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
