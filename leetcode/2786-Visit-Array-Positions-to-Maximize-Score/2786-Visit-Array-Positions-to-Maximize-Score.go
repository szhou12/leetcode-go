package leetcode

import "math"

// dp[i][0] := max score in nums[0:i] when the last selected number is even
// dp[i][1] := max score in nums[0:i] when the last selected number is odd
func maxScore(nums []int, x int) int64 {
	n := len(nums)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			dp[i][j] = math.MinInt / 2 // divide by 2 to avoid overflow bc we may have -x
		}
	}

	// base case
	if nums[0]%2 == 0 {
		dp[0][0] = nums[0]
	} else {
		dp[0][1] = nums[0]
	}

	// recurrence
	for i := 1; i < n; i++ {
		if nums[i]%2 == 0 { // EVEN
			dp[i][1] = dp[i-1][1] // NO WAY to select nums[i] as odd, so we have to skip i-th number

			dp[i][0] = max(dp[i][0], dp[i-1][0])           // skip i-th number
			dp[i][0] = max(dp[i][0], dp[i-1][0]+nums[i])   // add i-th number with no price for same parity
			dp[i][0] = max(dp[i][0], dp[i-1][1]+nums[i]-x) // add i-th number with price X for different parity
		} else { // ODD
			dp[i][0] = dp[i-1][0] // NO WAY to select nums[i] as even, so we have to skip i-th number

			dp[i][1] = max(dp[i][1], dp[i-1][1])         // skip i-th number
			dp[i][1] = max(dp[i][1], dp[i-1][1]+nums[i]) // add i-th number with no price for same parity
			dp[i][1] = max(dp[i][1], dp[i-1][0]+nums[i]-x)
		}
	}

	return int64(max(dp[n-1][0], dp[n-1][1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
