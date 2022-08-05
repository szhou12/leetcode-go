package leetcode

// My solution: Loop Backwards
func maxProfit_reverse(prices []int) int {
	profit := 0

	futureMaxPrice := prices[len(prices)-1]
	for day := len(prices) - 1; day >= 0; day-- {
		curProfit := futureMaxPrice - prices[day]
		profit = max(profit, curProfit)
		futureMaxPrice = max(prices[day], futureMaxPrice)
	}

	return profit
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
