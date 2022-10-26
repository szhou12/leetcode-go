package leetcode

// DP Method 1
// dp[i][j] := min # of deletions needed to make word1[1...i] identical to word2[1...j]
// 			   (word1[0], word2[0] placed with a placeholder)
// Base cases:
//	dp[0][j] = j, dp[i][0] = i, dp[0][0] = 0
// Recurrence:
//  dp[i][j] = dp[i-1][j-1] if word1[i] == word2[j]
//  dp[i][j] = min(dp[i-1][j-1]+2, dp[i-1][j]+1, dp[i][j-1]+1) o/w
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	word1 = "#" + word1
	word2 = "#" + word2
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// Base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Recurrence
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+2)
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
