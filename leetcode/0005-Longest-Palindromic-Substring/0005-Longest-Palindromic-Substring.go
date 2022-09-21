package leetcode

func longestPalindrome(s string) string {
	// Edge Case
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(s); j++ {
			if i == j { // Base Case
				dp[i][j] = true
				// Update res
				if j-i+1 > len(res) {
					res = s[i : i+1]
				}
				continue
			} else if i > j { // Base Case
				continue
			}

			// Recurrence
			if s[i] == s[j] {
				dp[i][j] = j-i+1 == 2 || dp[i+1][j-1]
			}
			// Update res
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
