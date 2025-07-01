package leetcode

// DP
// dp[i][0] := max profit in [0:i] without holding a stock
// dp[i][1] := max profit in [0:i] holding a stock
func maxProfit(prices []int) int {
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


// Greedy Algorithm
func maxProfit_Greedy(prices []int) int {
	sum := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			sum += prices[i] - prices[i-1]
		}
	}

	return sum
}
