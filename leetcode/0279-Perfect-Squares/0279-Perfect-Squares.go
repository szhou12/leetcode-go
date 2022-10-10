package leetcode

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)

	// Base Cases
	dp[0] = 0
	dp[1] = 1

	// Recurrence
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt
		if isPerfectSquare(i) {
			dp[i] = 1
		} else {
			for j := 1; j*j <= i; j++ {
				dp[i] = min(dp[i-j*j]+1, dp[i])
			}
		}
	}

	return dp[n]
}

func isPerfectSquare(x int) bool {
	if x == 1 {
		return true
	}
	for i := 1; i <= x/2; i++ {
		if i*i == x {
			return true
		}
	}
	return false
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
