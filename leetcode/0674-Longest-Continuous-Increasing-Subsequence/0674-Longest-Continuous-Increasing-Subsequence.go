package leetcode

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	globalMax := dp[0]

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		globalMax = max(globalMax, dp[i])
	}

	return globalMax
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
