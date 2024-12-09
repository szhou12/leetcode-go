package leetcode

import "math"

func maxSubArray(nums []int) int {
	n := len(nums)

	// dp[i] := max subarray sum ending at i
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = math.MinInt
	}

	// base case
	dp[0] = nums[0]

	// recurrence
	res := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
		res = max(res, dp[i])
	}

	return res
}
