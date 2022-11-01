package leetcode

// DP
// dp[i][j] = max length of longest common subarray for nums1[1...i] ending at i and nums2[1...j] ending at j
// base cases
// dp[0][0] = 0
// dp[i][0] = 0
// dp[0][j] = 0
// recurrence
// dp[i][j] = dp[i-1][j-1] + 1 if nums1[i] == nums2[j]
//          = 0 o/w
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	nums1 = append([]int{-1}, nums1...)
	nums2 = append([]int{-1}, nums2...)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}

	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			res = max(res, dp[i][j])
		}
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
