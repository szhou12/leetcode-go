package leetcode

// DP[i][k] := min operations to make nums[:i] beautiful if changing nums[i] = k (k=1,2,3)
// Base Case:
// DP[0][k] = 0
// Recurrence:
// DP[i][1] = DP[i-1][1] + (nums[i] != 1)
// DP[i][2] = min(DP[i-1][1], DP[i-1][2]) + (nums[i] != 2)
// DP[i][3] = min(DP[i-1][1], DP[i-1][2], DP[i-1][3]) + (nums[i] != 3)
func minimumOperations(nums []int) int {
	n := len(nums)

	nums = append([]int{0}, nums...)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 4)
	}

	for i := 1; i <= n; i++ {
		if nums[i] != 1 {
			dp[i][1] = dp[i-1][1] + 1
		} else {
			dp[i][1] = dp[i-1][1]
		}

		if nums[i] != 2 {
			dp[i][2] = min(dp[i-1][1], dp[i-1][2]) + 1
		} else {
			dp[i][2] = min(dp[i-1][1], dp[i-1][2])
		}

		if nums[i] != 3 {
			dp[i][3] = min(min(dp[i-1][1], dp[i-1][2]), dp[i-1][3]) + 1
		} else {
			dp[i][3] = min(min(dp[i-1][1], dp[i-1][2]), dp[i-1][3])
		}
	}

	return min(min(dp[n][1], dp[n][2]), dp[n][3])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
