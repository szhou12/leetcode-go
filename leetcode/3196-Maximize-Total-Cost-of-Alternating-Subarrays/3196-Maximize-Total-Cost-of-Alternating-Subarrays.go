package leetcode

import "math"

func maximumTotalCost(nums []int) int64 {
	n := len(nums)

	// dp[i][0] := the maximum sum of subarrays in nums[0:i] (inclusive) without sign change at i-th element
	// dp[i][1] := the maximum sum of subarrays in nums[0:i] (inclusive) with sign change at i-th element
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	// base case
	dp[0][0] = nums[0]
	dp[0][1] = math.MinInt // a sign-changed element must follow a non-sign-changed element, therefore first element cannot change its sign. Give it a "no-meaning" value

	// recurrence
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]) + nums[i]
		dp[i][1] = dp[i-1][0] - nums[i]
	}

	return int64(max(dp[n-1][0], dp[n-1][1]))
}
