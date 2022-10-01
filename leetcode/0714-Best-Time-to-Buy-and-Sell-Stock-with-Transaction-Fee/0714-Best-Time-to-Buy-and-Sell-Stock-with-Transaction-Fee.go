package leetcode

func maxProfit(prices []int, fee int) int {
	n := len(prices)

	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		// Base Cases: i == 0
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i] - fee
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)

	}

	return dp[n-1][0]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
