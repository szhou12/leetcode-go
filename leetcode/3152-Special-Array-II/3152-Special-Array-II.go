package leetcode

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1]
		if nums[i-1]%2 == nums[i]%2 {
			dp[i]++
		}
	}

	res := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		a, b := queries[i][0], queries[i][1]
		if dp[b] - dp[a] > 0 {
			res[i] = false
		} else {
			res[i] = true
		}
	}

	return res
}