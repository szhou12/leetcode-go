package leetcode

import "math"

// DP: 写法与 LC-188 完全一致
func maxProfit(prices []int) int {
	k := 2

	n := len(prices)
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
	}

	// Base Cases: k == 0
	for i := 0; i < n; i++ {
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt
	}

	for i := 0; i < n; i++ {
		for j := k; j >= 1; j-- {
			// Base Cases: i == 0
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[n-1][k][0]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
