package leetcode

func maximumLength(nums []int, k int) int {
	n := len(nums)

	// dp[i][t] := maximum possible length of a good subsequence of nums[0:i] including i-th element && there costs t times of seq[i] != seq[i + 1]
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
	}

	// base case: dp[0][0...k] = 1 all covered in recurrence

	// recurrence
	res := 1
	for i := 0; i < n; i++ {
		for t := 0; t <= k; t++ {

			cur := 1 // trick: for every i, 自立门户，最小可能是i号元素本身，所以长度为1

			for j := 0; j < i; j++ {
				if nums[j] == nums[i] {
					cur = max(cur, dp[j][t]+1)
				} else if t-1 >= 0 {
					cur = max(cur, dp[j][t-1]+1)
				}
			}

			dp[i][t] = cur
			res = max(res, dp[i][t]) // res在每一次i都要更新, 因为不一定i==n-1时是最大值
		}
	}

	return res
}