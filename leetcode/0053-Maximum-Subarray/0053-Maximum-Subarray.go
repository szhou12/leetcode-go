package leetcode

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	gmax := nums[0]
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		dp[i] = max(dp[i]+nums[i], nums[i])
		gmax = max(gmax, dp[i])
	}
	return gmax
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
