package leetcode

// DP Solution
func longestPalindrome_DP(s string) string {
	// Edge Case
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(s); j++ {
			if i == j { // Base Case
				dp[i][j] = true
				// Update res
				if j-i+1 > len(res) {
					res = s[i : i+1]
				}
				continue
			} else if i > j { // Base Case
				continue
			}

			// Recurrence
			if s[i] == s[j] {
				dp[i][j] = j-i+1 == 2 || dp[i+1][j-1]
			}
			// Update res
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

// Two-pointer Solution: 从中心向两端扩散的双指针
func longestPalindrome_TP(s string) string {
	// Edge Case
	if len(s) < 2 {
		return s
	}

	res := ""
	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心的最长回文子串
		p1 := palindrome(s, i, i)
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		p2 := palindrome(s, i, i+1)
		res = max(res, p1)
		res = max(res, p2)
	}
	return res
}

func palindrome(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	// while loop终止条件 使得我们要在返回substring时 排除掉 l, r 指向的字母
	return s[l+1 : r]
}

func max(s1 string, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}
