package leetcode

func climbStairs(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// dp[i] := # of ways to climb to i
// recurrence:
// dp[i] = dp[i-1] + dp[i-2]
// base case:
// dp[0] = 1, dp[1] = 1

// dp[2] = dp[2-1] + dp[2-2] = 1 + 1 = 2
// dp[3] = dp[3-1] + dp[3-2] = dp[2] + dp[1] = 2 + 1 = 3
// dp[4] = dp[3] + dp[2] = 3 + 2 = 5
// 4
// 1 1 + 2
// 2 + 2
// 1 1 1 + 1
// 1 2 + 1
// 2 1 + 1
