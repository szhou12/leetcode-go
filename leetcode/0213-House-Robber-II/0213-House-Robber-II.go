package leetcode

func rob(nums []int) int {
	// Edge Cases
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	n := len(nums)
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange(nums []int, start int, end int) int {
	dp := make([]int, end-start+1)
	// Base Cases
	dp[0] = nums[start]
	dp[1] = max(nums[start], nums[start+1])

	// Recurrence
	for i := start + 2; i <= end; i++ {
		// 注意！这里dp的index 要 -start
		dp[i-start] = max(dp[i-start-1], dp[i-start-2]+nums[i])
	}

	return dp[len(dp)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
