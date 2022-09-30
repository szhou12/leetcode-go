package leetcode

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
