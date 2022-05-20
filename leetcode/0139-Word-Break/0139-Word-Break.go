package leetcode

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	s = "#" + s

	for i := 1; i <= n; i++ {
		dp[i] = false
		for _, w := range wordDict {
			t := len(w)
			if t <= i && dp[i-t] && s[i-t+1:i+1] == w {
				dp[i] = true
			}
		}
	}

	return dp[n]
}
