package leetcode

import "math"

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	// base case: at starting position, need 0 jumps
	dp[0] = 0

	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] >= i-j {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// testing: [2,3,1,1,4]
// dp: [0, max, max, max, max]
// i: 1 j: 0
// dp: [0, 1, max, max, max]
// i: 2 j: 0, 1
// dp: [0, 1, min(2,1)=1, max, max]
// i: 3 j: 0, 1, 2
// dp: [0, 1, 1, min(2,2)=2, max]
// i: 4 j: 0, 1, 2, 3
// dp: [0, 1, 1, 2, min(2, 3)=2]
