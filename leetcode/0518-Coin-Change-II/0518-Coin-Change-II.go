package leetcode

// DP - Optimal Solution
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	// Base Case
	dp[0] = 1
	// Recurrence
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// DFS - 会超时
// # of branches = # of coins[i]
// # levels = len(coins)
// each level = coins[i]
// 注意: 这道题只要求满足要求的个数, 不需要把解都找出来, 所以curCombo, res [][]int 可以省略
func change_DFS(amount int, coins []int) int {
	var res int
	// var res [][]int
	// curCombo := make([]int, 0)

	dfs(amount, coins, 0, &res)
	return res
}

func dfs(amount int, coins []int, index int, res *int) {
	// Base case
	if index == len(coins) {
		if amount == 0 {
			// c := make([]int, len(curCombo))
			// copy(c, curCombo)
			*res++
		}
		return
	}

	for n := 0; n*coins[index] <= amount; n++ {
		// curCombo = append(curCombo, n)
		// dfs(amount-n*coins[index], coins, index+1, curCombo, res)
		// curCombo = curCombo[:len(curCombo) - 1]
		dfs(amount-n*coins[index], coins, index+1, res)
	}
}
