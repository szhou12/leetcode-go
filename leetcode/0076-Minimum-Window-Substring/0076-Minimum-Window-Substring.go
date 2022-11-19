package leetcode

import "math"

// 写法一: 模版Fix
func minWindow(s string, t string) string {
	need := make(map[byte]int) // hashmap t
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	window := make(map[byte]int)
	left, right := 0, 0
	count := 0
	start, length := 0, math.MaxInt // result
	for right < len(s) {
		rightElement := s[right] // 吃
		right++
		if _, ok := need[rightElement]; ok {
			window[rightElement]++
			if window[rightElement] == need[rightElement] {
				count++
			}
		}

		for count == len(need) { // 开始收缩左边界的条件: 需要的字符及相应的次数都包括了
			// update result
			if right-left < length {
				length = right - left
				start = left
			}

			leftElement := s[left] // 吐
			left++
			if _, ok := need[leftElement]; ok {
				if window[leftElement] == need[leftElement] {
					count--
				}
				window[leftElement]--
			}
		}
	}

	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}

// 写法二: 模版Flex
func minWindow_flex(s string, t string) string {
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	window := make(map[byte]int)
	right := 0
	count := 0
	start, length := 0, math.MaxInt
	for left := 0; left < len(s); left++ {
		// 固定左边界，延伸右边界
		for right < len(s) && count < len(need) {
			rightElement := s[right]
			if _, ok := need[rightElement]; ok {
				window[rightElement]++
				if window[rightElement] == need[rightElement] {
					count++
				}
			}
			right++
		}

		// update result
		if count == len(need) {
			if right-left < length {
				length = right - left
				start = left
			}
		}

		// 吐
		leftElement := s[left]
		if _, ok := need[leftElement]; ok {
			if window[leftElement] == need[leftElement] {
				count--
			}
			window[leftElement]--
		}
	}

	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}
