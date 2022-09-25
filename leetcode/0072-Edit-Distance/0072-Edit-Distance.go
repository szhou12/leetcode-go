package leetcode

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	word1 = "#" + word1
	word2 = "#" + word2
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// Base Cases
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = i // transform first i chars of word1 --> word2 that is an empty string => delete i chars
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j // transform empty string word1 --> word2 of first j chars => insert j chars
	}

	// Recurrence
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				replaceCost := 1 + dp[i-1][j-1]
				deleteCost := 1 + dp[i-1][j]
				insertCost := 1 + dp[i][j-1]
				dp[i][j] = min(replaceCost, deleteCost)
				dp[i][j] = min(dp[i][j], insertCost)
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
