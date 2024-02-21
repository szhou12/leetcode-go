package leetcode

import "sort"

func maxSelectedElements(nums []int) int {
	n := len(nums)

	// STEP 1: Sort in ascending order
	sort.Ints(nums)

	// STEP 2: DP
	dp := make([][]int, n)
	for i := 0; i < n ; i++ {
		dp[i] = make([]int, 2)
	}

	// Base case
	dp[0][0] = 1
	dp[0][1] = 1

	// Recurrence
	for i := 1; i < n; i++ {
		// 另起炉灶
		dp[i][0] = 1
		dp[i][1] = 1
		if nums[i] - nums[i-1] == 0 {
			dp[i][1] = max(dp[i][1], dp[i-1][0] + 1)
			dp[i][1] = max(dp[i][1], dp[i-1][1])
			dp[i][0] = max(dp[i][0], dp[i-1][0])
		} else if nums[i]-nums[i-1] == 1 {
			dp[i][0] = max(dp[i][0], dp[i-1][0] + 1)
			dp[i][1] = max(dp[i][1], dp[i-1][1] + 1)
		} else if nums[i]-nums[i-1] == 2 {
			dp[i][0] = max(dp[i][0], dp[i-1][1]+1)
		}
	}

	res := 1
	for i := 0; i < n; i++ {
		res = max(res, max(dp[i][0], dp[i][1]))
	}

	return res
}

// key 1: sort first to make elements with similar values group together

// dp[i][0] := len of longest consecutive elements in nums[:i] (inclusive) ending at i with nums[i]+0
// dp[i][1] := len of longest consecutive elements in nums[:i] (inclusive) ending at i with nums[i]+1
// do you understand why we need to annotate status 0 and 1?

// X X X i-1, i X X X
//        a   b
// 1. b-a == 2 => a+1, b+0
// dp[i][0] = dp[i-1][1] + 1
// 2. b-a == 1 => a+0, b+0 OR a+1, b+1
// dp[i][0] = dp[i-1][0] + 1
// dp[i][1] = dp[i-1][1] + 1
// 3. b-a == 0 => a+0, b+1 OR throw out b, directly inherits len ending at a or a+1
// dp[i][1] = dp[i-1][0] + 1
// dp[i][1] = dp[i-1][1]
// dp[i][0] = dp[i-1][0]
// 4. b 另起炉灶
// dp[i][0] = 1
// dp[i][1] = 1
