package leetcode

func lengthOfLIS(nums []int) int {
	n := len(nums)
	// prepend nums with a placeholder
	nums = append([]int{-1}, nums...)

	dp := make([]int, n+1)
	dp[0] = 0
	// init every element to be 1 bc one num itself is a subseq
	for i := 1; i <= n; i++ {
		dp[i] = 1
	}

	res := dp[1]
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
