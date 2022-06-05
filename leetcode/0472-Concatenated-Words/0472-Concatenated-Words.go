package leetcode

func findAllConcatenatedWordsInADict(words []string) []string {
	wordDict := make(map[string]bool)
	for _, word := range words {
		wordDict[word] = true
	}
	var res []string
	for _, word := range words {
		// corner case
		if word == "" {
			continue
		}
		wordDict[word] = false
		if canBreak(word, wordDict) {
			res = append(res, word)
		}
		wordDict[word] = true
	}
	return res
}

func canBreak(word string, wordDict map[string]bool) bool {
	n := len(word)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDict[word[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]

}
