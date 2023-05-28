package leetcode

func takeCharacters(s string, k int) int {
	n := len(s)
	count := make(map[byte]int)
	count['a'] = 0
	count['b'] = 0
	count['c'] = 0
	for i := 0; i < n; i++ {
		count[s[i]]++
	}
	if count['a'] < k || count['b'] < k || count['c'] < k {
		return -1
	}

	window := make(map[byte]int)
	window['a'], window['b'], window['c'] = 0, 0, 0
	windowSize := 0
	right := 0
	for left := 0; left < n; left++ {
		// 吃
		for right < n && (count['a']-window['a'] >= k && count['b']-window['b'] >= k && count['c']-window['c'] >= k) {
			rightElement := s[right]
			window[rightElement]++
			right++
		}

		// update window size based on different senarios
		// 1. if right out of bound => [left, right-1]
		// 2. if cannot expand right => [left, right-2]
		// 3. if right out of bound & cannot expand right => [left, right-2]
		if !(right < n) && (count['a']-window['a'] >= k && count['b']-window['b'] >= k && count['c']-window['c'] >= k) {
			windowSize = max(windowSize, right-1-left+1)
		} else if right < n && !(count['a']-window['a'] >= k && count['b']-window['b'] >= k && count['c']-window['c'] >= k) {
			windowSize = max(windowSize, right-2-left+1)
		} else if !(right < n) && !(count['a']-window['a'] >= k && count['b']-window['b'] >= k && count['c']-window['c'] >= k) {
			windowSize = max(windowSize, right-2-left+1)
		}

		// 吐
		leftElement := s[left]
		window[leftElement]--

	}

	return n - windowSize

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
