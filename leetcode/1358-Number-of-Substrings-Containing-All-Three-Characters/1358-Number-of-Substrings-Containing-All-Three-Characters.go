package leetcode

func numberOfSubstrings(s string) int {
	n := len(s)

	window := make(map[byte]int)
	window['a'] = 0
	window['b'] = 0
	window['c'] = 0

	right := 0
	res := 0
	for left := 0; left < n; left++ {
		// 吃
		for right < n && (window['a'] == 0 || window['b'] == 0 || window['c'] == 0) {
			rightElement := s[right]
			window[rightElement]++
			right++
		}

		// update result
		// NOTE: only update result when a, b, c all exist in window
		// DO NOT update result is right is simply out of bound
		// Check if a, b, c all exist in window
		if !(window['a'] == 0 || window['b'] == 0 || window['c'] == 0) {
			// NOTE: right物理意义表示第一个满足滑窗内满足都至少包含一个a,b,c的位置，
			// 也就是说，从这个right到最末尾这段区间内，都是满足条件的
			res += (n - 1) - (right - 1) + 1
		}

		// 吐
		leftElement := s[left]
		window[leftElement]--
	}

	return res
}
