package leetcode

func findAnagrams(s string, p string) []int {
	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	window := make(map[byte]int)
	left, right := 0, 0
	res := make([]int, 0)
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

		for right-left >= len(p) { // 开始收缩左边界的条件: [left, right) 定长 len(p) (ie. 上限不能超过len(p))
			// update result
			if count == len(need) {
				res = append(res, left)
			}

			leftElement := s[left] // 吐
			left++
			if _, ok := window[leftElement]; ok {
				if window[leftElement] == need[leftElement] {
					count--
				}
				window[leftElement]--
			}
		}
	}
	return res
}
