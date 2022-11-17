package leetcode

import "math"

func minWindow(s string, t string) string {
	need := make(map[byte]int) // hashmap t
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	window := make(map[byte]int)
	left, right := 0, 0
	start, length := 0, math.MaxInt
	count := 0
	for right < len(s) {
		rightElement := s[right] // 吃
		right++
		if _, ok := need[rightElement]; ok {
			window[rightElement]++
			if window[rightElement] == need[rightElement] {
				count++
			}
		}

		for count == len(need) {
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
