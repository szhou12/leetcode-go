package leetcode

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, len(s)+1)
	dp[0] = true

	// convert wordDict to a set
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

// the inner loop may exceed time limit!
func wordBreak_deprecated(s string, wordDict []string) bool {
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
