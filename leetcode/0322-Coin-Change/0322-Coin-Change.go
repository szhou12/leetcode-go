package leetcode

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ { // 遍历背包容量
		for j := 0; j < len(coins); j++ { // 遍历物品 - 每一种coin
			if i >= coins[j] {
				dp[i] = min(dp[i], 1+dp[i-coins[j]])
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
