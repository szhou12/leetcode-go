package leetcode

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	s = "#" + s
	t = "#" + t
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// Base Cases
	dp[0][0] = 1
	// dp[i][0] = 1
	for i := 1; i <= m; i++ {
		dp[i][0] = 1
	}
	// dp[0][j] = 0
	for j := 1; j <= n; j++ {
		dp[0][j] = 0
	}

	// Recurrence
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i] == t[j] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}
