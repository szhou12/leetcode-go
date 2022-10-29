package leetcode

// DP
// dp[i][j] = longest common subseq of text1[1..i] and text2[1...j]
// base cases:
// dp[0][0] = 0
// dp[i][0] = 0
// dp[0][j] = 0
// recurrence
// dp[i][j] = dp[i-1][j-1] + 1 if text1[i] == text2[j]
//          = max(dp[i-1][j], dp[i][j-1]) o/w
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	text1 = "#" + text1
	text2 = "#" + text2

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// base case
	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
