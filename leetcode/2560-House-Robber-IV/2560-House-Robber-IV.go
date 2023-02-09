package leetcode

import "math"

func minCapability(nums []int, k int) int {
	left, right := 0, maxVal(nums)

	for left < right {
		mid := left + (right-left)/2

		if atLeastK(nums, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func atLeastK(nums []int, k int, cap int) bool {
	n := len(nums)

	// dp[i][0]: # of houses we can rob in nums[:i] if we DO NOT rob i-th house (nums[i])
	// dp[i][1]: # of houses we can rob in nums[:i] if we DO rob i-th house (nums[i])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	// Base case
	dp[0][0] = 0       // if we DO NOT rob 0-th house
	if nums[0] > cap { // we CAN NOT rob 0-th house. dp[0][1] (Rob 0-th house) 是无意义的, 故赋值无限小
		dp[0][1] = math.MinInt
	} else {
		dp[0][1] = 1
	}

	// Recurrrence
	for i := 1; i < n; i++ {
		if nums[i] > cap { // we CAN NOT rob i-th house
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
			dp[i][1] = math.MinInt
		} else { // we CAN rob i-th house
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
			dp[i][1] = dp[i-1][0] + 1
		}
	}

	maxHouseRobbed := max(dp[n-1][0], dp[n-1][1])
	if maxHouseRobbed < k {
		return false
	} else {
		return true
	}
}

// find biggest element in nums
func maxVal(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = max(res, nums[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
