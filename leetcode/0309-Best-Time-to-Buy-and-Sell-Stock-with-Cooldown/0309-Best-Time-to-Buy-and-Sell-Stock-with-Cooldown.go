package leetcode

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)

	for i := 0; i < n; i++ {
		// Base Cases: i==0
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		// Base Cases: i == 1 due to cooldown
		if i == 1 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], -prices[i])
			// dp[1][1] = max(dp[0][1], dp[-1][0]-prices[i])
			//			= max(dp[0][1], 0-prices[i]) 还没开始i==-1就不持有股票显然是利润是0
			//			= max(dp[0][1], prices[i])
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i]) // i-2 due to cooldown
	}

	return dp[n-1][0]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
