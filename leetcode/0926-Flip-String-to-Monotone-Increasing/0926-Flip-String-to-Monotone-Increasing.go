package leetcode

func minFlipsMonoIncr(s string) int {
	n := len(s)
	s = "#" + s

	prevOnes := 0
	dp := make([]int, n+1)

	// base case
	dp[0] = 0

	for i := 1; i <= n; i++ {
		if string(s[i]) == "1" {
			dp[i] = dp[i-1]
			prevOnes += 1
		} else {
			dp[i] = min(dp[i-1]+1, prevOnes)
		}
	}
	return dp[n]

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
