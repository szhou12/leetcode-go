package leetcode

// DP
func fib(n int) int {
	// Edge case
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// dp[i] = F(i) from 0 to i
// Base cases
// 	dp[0] = 0
// 	dp[1] = 1
// Recurrence
// 	dp[i] = dp[i-2] + dp[i-1]
