package leetcode

import "math"

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)

	prefixSum := make([]int, n)
	prefixSum[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	// dp[i] := max k-size subarray sum ending at i
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = math.MinInt
	}

	// base case
	dp[k-1] = prefixSum[k-1]

	// recurrence
	res := dp[k-1]
	for i := k; i < n; i++ {
		dp[i] = max(dp[i-k], 0) + (prefixSum[i] - prefixSum[i-k])
		res = max(res, dp[i])
	}

	return int64(res)
}
