package leetcode

// My solution: Loop Backwards 每一轮保留最大的Price
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

// Solution 2: Loop forward 每一轮保留回头看过的最小的Price
func maxProfit_forward(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for day := 1; day < len(prices); day++ {
		if prices[day]-minPrice > maxProfit {
			maxProfit = prices[day] - minPrice
		}
		if prices[day] < minPrice {
			minPrice = prices[day]
		}
	}
	return maxProfit
}
