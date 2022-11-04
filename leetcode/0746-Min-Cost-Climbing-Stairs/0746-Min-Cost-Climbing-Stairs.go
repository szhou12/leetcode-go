package leetcode

// DP
// 根据题目, 2 <= len(cost) <= n, 故无需处理edge cases
// dp[i] = min cost to reach i-th cell ending at i
func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	dp := make([]int, n)
	// base cases
	dp[0] = cost[0]
	dp[1] = cost[1]

	// recurrence
	for i := 2; i < n; i++ {
		// By def of dp[i], it ends at i, so it has to add i-th cell's cost
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[n-1], dp[n-2])
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
