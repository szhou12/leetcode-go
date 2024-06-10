package leetcode

var M = int(1e9 + 7)

func valueAfterKSeconds(n int, k int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for t := 0; t < k; t++ {
		for i := 1; i < n; i++ {
			dp[i] = (dp[i-1] + dp[i]) % M
		}
	}
	return dp[n-1]

}
