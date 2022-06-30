package leetcode

func countHousePlacements(n int) int {
	modulo := int64(1e9 + 7)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 1
	dp[0][1] = 0

	for i := 1; i <= n; i++ {
		dp[i][0] = int(int64(dp[i-1][1]+dp[i-1][0]) % modulo) // NOTE: addition here could cause overflow, thus modulus
		dp[i][1] = dp[i-1][0]
	}

	res := int64(dp[n][0]+dp[n][1]) % modulo

	return int(res * res % modulo)

}
