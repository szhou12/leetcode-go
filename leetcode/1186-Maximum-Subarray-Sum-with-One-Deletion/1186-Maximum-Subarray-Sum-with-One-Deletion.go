package leetcodd

// dp[i][0] := maximum sum for a non-empty subarray ending at i, without a deletion
// dp[i][1] := maximum sum for a non-empty subarray ending at i, with a deletion
func maximumSum(arr []int) int {
	n := len(arr)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	// base cases
	dp[0][0] = arr[0]
	dp[0][1] = 0 // it's ok to set this way as it's only used in dp[i-1][1] + arr[i]. But we need to be careful about the starting value of res should be dp[0][0] instead of MinInt in case of only one element in the array

	res := dp[0][0]

	// recurrence
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0] + arr[i], arr[i])
		dp[i][1] = max(dp[i-1][1] + arr[i], dp[i-1][0])
		res = max(res, max(dp[i][0], dp[i][1]))
	}

	return res
}