package leetcode

// My Solution
func longestSemiRepetitiveSubstring(s string) int {
	n := len(s)
	res := 0

	pairFound := false // if a consecutive pair exists in Sliding Window
	right := 0
	for left := 0; left < n; left++ {
		// 吃
		for right < n {
			if pairFound && s[right] == s[right-1] {
				break
			}
			if right > 0 && s[right] == s[right-1] {
				pairFound = true
			}
			right++
		}

		// update result: [left, right)
		res = max(res, right-left)
		// 吐
		if left + 1 < n && s[left] == s[left+1] {
			pairFound = false
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 参考答案: 适用度高, 能处理允许多对重复的情况
func longestSemiRepetitiveSubstring_standard(s string) int {
	n := len(s)
	res := 0

	count := 0 // # of consecutive pairs within Sliding Window
	right := 0
	for left := 0; left < n; left++ {
		// 吃
		for right < n && count + isPair(s, left, right) < 2 {
			count += isPair(s, left, right)
			right++
		}

		// update result
		res = max(res, right-left)
		// 吐
		if left + 1 < n && s[left] == s[left+1] {
			count--
		}
	}

	return res

}

func isPair(s string, left int, right int) int {
	// NOTE: 这里的条件写成下面形式也能过
	// (right > 0 && s[right] == s[right-1])
	if left < right && s[right] == s[right-1] {
		return 1
	}
	return 0
}