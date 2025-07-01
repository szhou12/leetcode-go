package leetcode

// dp[i] := the minimum number of coins to make up amount i
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	// NOTE：这里内外层循环对调是没有问题的
	
	// for i := 1; i <= amount; i++ { // 遍历背包容量
	// 	for j := 0; j < len(coins); j++ { // 遍历物品 - 每一种coin
	// 		if i >= coins[j] {
	// 			dp[i] = min(dp[i], 1+dp[i-coins[j]])
	// 		}
	// 	}
	// }

	// 外层遍历物品
	for j := 0; j < len(coins);j++ {
        for i := 1; i <= amount; i++ { // 遍历背包容量
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
