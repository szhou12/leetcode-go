package leetcode

func maxProfit(prices []int, fee int) int {
	n := len(prices)

	// dp[i][0] := max profit for NOT holding the stock on i-th day
	// dp[i][1] := max profit for holding the stock on i-th day
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	// Base case
	dp[0][0] = 0
	dp[0][1] = -prices[0] - fee // 贷款买入

	// Recurrence
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]-prices[i]-fee, dp[i-1][1])
	}

	return dp[n-1][0]
}

