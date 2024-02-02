package leetcode

var M = int(1e9 + 7)

func sumCounts(nums []int) int {
	n := len(nums)

	// STEP 1: for each i, find last seen position of nums[i]
	position := make(map[int]int) // {nums[i] : recently seen index}
	prev := make([]int, n)        // prev[i] = last seen position of nums[i]
	// initialize prev: MUST init to -1 bc 0 is a valid position
	for i := 0; i < n; i++ {
		prev[i] = -1
	}
	for i := 0; i < n; i++ {
		if k, ok := position[nums[i]]; ok {
			prev[i] = k
		}
		position[nums[i]] = i
	}

	// STEP 2: DP + Segment Tree
	// dp[i] := sum of squares of distinct counts of all subarrays in num[0:i] ending at i
	root := NewSegTreeNode(0, n-1, 0)
	dp := make([]int, n)
	// recurrence loop MUST start at i=0. No need to addtionally write a base case
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] += 0
		} else {
			dp[i] += dp[i-1]
		}
		k := prev[i]
		dp[i] += 2 * root.queryRange(k+1, i-1)
		dp[i] += i - 1 - k
		dp[i] += 1
		dp[i] %= M

		// update segment tree nums[i] contributes 1 count to range [k+1:i]
		root.updateRangeBy(k+1, i, 1)
	}

	// STEP 3: collect all subarrays ending at each i
	res := 0
	for i := 0; i < n; i++ {
		res += dp[i]
		res %= M
	}
	return res

}

// Segment Tree Implementation
