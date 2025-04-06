package leetcode

func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)

	// dp[i][v] := true if i-th element of nums[] can be reduced to 0 from current value v
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, 1001)
	}

	// base case
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}

	// edge case: nums[] = [0, 0, ..., 0], no need to transform
	if isOk(nums, dp) {
		return 0
	}

	// recurrence
	for j, q := range queries {
		l, r, val := q[0], q[1], q[2]
		for i := l; i <= r; i++ {
			for v := 1000; v >= 0; v-- {
				if v-val >= 0 && dp[i][v-val] {
					dp[i][v] = true
				}
			}
		}

		if isOk(nums, dp) {
			return j+1
		}
	}

	return -1
}

func isOk(nums []int, dp [][]bool) bool {
	for i := 0; i < len(nums); i++ {
		if !dp[i][nums[i]] {
			return false
		}
	}
	return true
}