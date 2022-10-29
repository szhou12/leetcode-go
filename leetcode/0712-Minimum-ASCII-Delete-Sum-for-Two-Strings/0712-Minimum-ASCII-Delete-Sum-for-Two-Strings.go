package leetcode

// DP
// dp[i][j] = min sum of deleted chars to make s1[1...i] == s2[1...j]
// base cases
// dp[0][0] = 0
// dp[i][0] = sum(s1[i]) for all i
// dp[0][j] = sum(s2[j]) for all j
// recurrence
// dp[i][j] = dp[i-1][j-1] if s1[i] == s2[j]
//          = min(dp[i-1][j]+s1[i], dp[i][j-1]+s2[j], dp[i-1][j-1]+s1[i]+s2[j])
func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	s1 = "#" + s1
	s2 = "#" + s2
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i])
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i]), dp[i][j-1]+int(s2[j]))
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+int(s1[i])+int(s2[j]))
			}
		}
	}
	return dp[m][n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
