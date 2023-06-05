package leetcode

/*
 * Naive Solution
 * loop through all letters A-Z to be the target letter within the window
 */
func characterReplacement(s string, k int) int {
	res := 0

	for i := 0; i < 26; i++ {
		target := byte('A' + i)
		curLongest := slidingWindow(s, k, target)
		res = max(res, curLongest)
	}

	return res
}

func slidingWindow(s string, k int, target byte) int {
	n := len(s)
	size := 0

	right := 0
	for left := 0; left < n; left++ { // 左闭右闭 [left, right]
		// 吃
		for right < n && (k > 0 || s[right] == target) {
			if s[right] != target {
				k--
			}
			right++
		}

		// update result
		size = max(size, right-left)

		// 吐
		leftElement := s[left]
		if leftElement != target {
			k++
		}
	}

	return size
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 * Elegant Solution
 * loop through all letters A-Z to be the target letter within the window
 */
func characterReplacement_elegant(s string, k int) int {
	n := len(s)

	count := make([]int, 26)
	res := 0
	right := 0
	for left := 0; left < n; left++ {
		// 吃
		for right < n && canExpand(s, k, left, right, count) {
			count[int(s[right]-'A')]++
			right++
		}

		// update result
		res = max(res, right-left)

		// 吐
		count[int(s[left]-'A')]--

	}
	return res
}

func canExpand(s string, k int, left int, right int, count []int) bool {
	count[int(s[right]-'A')]++
	majority := getMax(count)
	totalSize := right - left + 1
	count[int(s[right]-'A')]--
	return totalSize-majority <= k
}

func getMax(count []int) int {
	res := 0
	for _, cnt := range count {
		if cnt > res {
			res = cnt
		}
	}
	return res
}
