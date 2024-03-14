package leetcode

func minimumTimeToInitialState(word string, k int) int {
	n := len(word)
	lsp := preprocess(word)

	keptLen := lsp[n-1]
	for keptLen > 0 && (n-keptLen)%k != 0 {
		// find next longest suffix prefix
		keptLen = lsp[keptLen-1]
	}

	if keptLen == 0 { // all been cut off
		if n%k == 0 {
			return n / k
		} else { // 切不完整, 多切一刀, 切掉剩余
			return n/k + 1
		}
	} else {
		return (n - keptLen) / k
	}
}

// KMP STEP 1
func preprocess(needle string) []int {
	n := len(needle)

	// dp[i] := max length j s.t. needle[0:j-1] == needle[i-j+1:i] (inclusive)
	dp := make([]int, n)

	// base case: require true prefix and suffix (i.e. needle itself as a whole can't be a prefix/suffix)
	dp[0] = 0

	for i := 1; i < n; i++ {
		j := dp[i-1]
		for j > 0 && needle[j] != needle[i] {
			j = dp[j-1]
		}
		if needle[j] == needle[i] {
			dp[i] = j + 1
		} else {
			dp[i] = j
		}
	}

	return dp
}
