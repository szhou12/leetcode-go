package leetcode

import "math"

// Top-down method: Recursion + DP memoization
func getMoneyAmount_Memo(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	return dfs(&dp, 1, n)
}

func dfs(dp *[][]int, left int, right int) int {
	// base case
	if left >= right { // only one number left to guess, it MUST be the final answer
		return 0
	}
	// pruning: if [left...right] already computed, no need to compute again
	if (*dp)[left][right] != 0 {
		return (*dp)[left][right]
	}

	res := math.MaxInt
	for i := left; i <= right; i++ { // O(n)
		cost := i + max(dfs(dp, left, i-1), dfs(dp, i+1, right))
		res = min(res, cost)
	}

	(*dp)[left][right] = res
	return res
}

// Bottom-up method: DP
func getMoneyAmount(n int) int {
	// NOTE: length defined as n+2
	// bc at min step, both lower bound and upper bound will be out of bound
	// dp[1][0]; dp[n+1][n]; they should be set to default value 0
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	// NOTE: 这里的两层for-loop, 外层是 loop length, 内层是 loop start position
	// NOTE: length 最小值从 2 开始, 因为length=1时只有一个值, 已经猜到答案, cost是默认值0, 无需再DP
	for length := 2; length <= n; length++ {
		for i := 1; i+length-1 <= n; i++ {
			j := i + length - 1
			dp[i][j] = math.MaxInt

			for k := i; k <= j; k++ {
				cost := k + max(dp[i][k-1], dp[k+1][j])
				dp[i][j] = min(dp[i][j], cost)
			}
		}
	}

	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
