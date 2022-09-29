package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	dp := make([]int, len(nums))
	if nums[0] == 1 {
		dp[0] = 1
	}

	res := dp[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 0
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp[i] := max # of consecutive 1 ending at i
// Base case:
// dp[0] = 1 if a[0] == 1
//       = 0 o/w
// Recurrence:
// dp[i] = dp[i-1] + 1 if a[i] == 1
//       = 0           o/w
