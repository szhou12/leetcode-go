package leetcode

import "math"

func maxProfit(k int, prices []int) int {
	n := len(prices)

	// 如果 允许交易的次数 k > n/2, 相当于 k = +inf. k, 不再起作用, 直接用 LC-122的算法
	if k > n/2 {
		return maxProfitINF(prices)
	}

	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
	}

	// Base Cases: k==0
	for i := 0; i < n; i++ {
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt
	}

	for i := 0; i < n; i++ {
		for j := k; j >= 1; j-- { // 正反写j++/j--都可以，j--更符合语义
			// Base Cases: i==0
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			// Recurrence
			// 第i天不持有 (卖出) 股票
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			// 第i天持有 (买入) 股票
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

func maxProfitINF(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		// base cases
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		//recurrence
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}
