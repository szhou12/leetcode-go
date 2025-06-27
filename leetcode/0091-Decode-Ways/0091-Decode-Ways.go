package leetcode

// dp[i] = number of ways to decode s[:i] (inclusive)
func numDecodings(s string) int {
	n := len(s)

	s = "#" + s // prepend for easy indexing

	dp := make([]int, n+1)

	// base cases
	dp[0] = 1 // empty digit string has 1 way to decode - picking nothing
	if s[1] != '0' {
		dp[1] = 1
	}

	// recurrence
	for i := 2; i <= n; i++ {
		// case 1: i-th digit is not 0, it can be added to every decoding ways in dp[i-1]
		if s[i] != '0' {
			dp[i] += dp[i-1]
		}

		// case 2: i-th digit and i-1-th digit form a valid two-digit number (10-19, 20-26)
		// it can be added to every decoding ways in dp[i-2]
		if s[i-1] == '1' && ('0' <= s[i] && s[i] <= '9') {
			dp[i] += dp[i-2]
		} else if s[i-1] == '2' && ('0' <= s[i] && s[i] <= '6') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}