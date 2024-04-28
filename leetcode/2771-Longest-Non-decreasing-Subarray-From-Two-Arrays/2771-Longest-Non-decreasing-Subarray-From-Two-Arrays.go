package leetcode

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	// dp[i][0/1] := max subarray length ending at i for taking min/max(nums1[i], nums2[i])
	// dp[i][0] = 1 if min([i], [i]) < min([i-1], [i-1])
	//          = max(dp[i-1][0], dp[i-1][1]) + 1 if min([i], [i]) >= max([i-1], [i-1]) && min([i], [i]) >= min([i-1], [i-1])
	// dp[i][1] = 1 if max([i], [i]) < min([i-1], [i-1])
	//          = max(dp[i-1][0], dp[i-1][1]) + 1 if max([i], [i]) >= max([i-1], [i-1]) && max([i], [i]) >= min([i-1], [i-1])
	n := len(nums1)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	// base case
	dp[0][0] = 1
	dp[0][1] = 1
	// recurrence
	for i := 1; i < n; i++ {
		// when take min at i
		if min(nums1[i], nums2[i]) < min(nums1[i-1], nums2[i-1]) { // 如果比前i-1的最小值还小，接不上
			dp[i][0] = 1
		} else { // 比i-1的最小值要大，起码可以接上i-1取min的case
			dp[i][0] = dp[i-1][0] + 1
			if min(nums1[i], nums2[i]) >= max(nums1[i-1], nums2[i-1]) { // 比i-1的最大值要大，可以再试试接上i-1取max的case
				dp[i][0] = max(dp[i][0], dp[i-1][1]+1)
			}
		}

		// when take max at i
		if max(nums1[i], nums2[i]) < min(nums1[i-1], nums2[i-1]) {
			dp[i][1] = 1
		} else {
			dp[i][1] = dp[i-1][0] + 1
			if max(nums1[i], nums2[i]) >= max(nums1[i-1], nums2[i-1]) {
				dp[i][1] = max(dp[i][1], dp[i-1][1]+1)
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, max(dp[i][0], dp[i][1]))
	}
	return res

}
