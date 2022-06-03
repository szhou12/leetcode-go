package leetcode

// dp[i][j] := length of longest padlindrome subseq of s[i...j]
func longestPalindromeSubseq(s string) int {
	n := len(s)

	// prepend s with a placeholder
	s = "#" + s

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := n; i >= 1; i-- {
		for j := 1; j <= n; j++ {
			if i > j {
				dp[i][j] = 0
			} else if i == j {
				dp[i][j] = 1
			} else {
				if s[i] == s[j] {
					dp[i][j] = 2 + dp[i+1][j-1]
				} else {
					dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				}
			}
		}
	}
	return dp[1][n]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
