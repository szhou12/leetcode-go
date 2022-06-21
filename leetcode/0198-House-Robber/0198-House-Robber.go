package leetcode

// My solution
func rob(nums []int) int {
	n := len(nums)
	// edge cases
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		for j := 0; j < i-1; j++ {
			dp[i] = max(dp[i], dp[j]+nums[i])
		}
		dp[i] = max(dp[i-1], dp[i])
	}
	return dp[n-1]

}

// Optimizing Solution
func rob_opt(nums []int) int {
	n := len(nums)
	// edge cases
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]

}

// Optimizing Solution Again for Space
func rob_opt2(nums []int) int {
	n := len(nums)
	// edge cases
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	pvpv := nums[0]
	pv := max(nums[0], nums[1])
	res := pv
	for i := 2; i < n; i++ {
		res = max(pv, pvpv+nums[i])
		pvpv = pv
		pv = res
	}
	return res

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
