package leetcode

func strStr_KMP(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	// edge cases
	if m == 0 {
		return 0
	}
	if n == 0 {
		return -1
	}

	// build lsp for pattern string = needle
	lsp := preprocess(needle)

	// dp[i] = max length j of suffix haystack[:i] s.t. prefix needle[0:j-1] == suffix haystack[i-j+1:i] (双闭区间)
	dp := make([]int, n)
	// base case
	dp[0] = 0

}

func preprocess(needle string) []int {
	n := len(needle)

	// dp[i] = max length j s.t. prefix needle[0:j-1] == suffix needle[i-j+1:i] (双闭区间)
	dp := make([]int, n)

	// base case
	dp[0] = 0

	// recurrence
	for i := 1; i < n; i++ {
		j := dp[i-1]
		for j>=1 && needle[i] != needle[j] {
			j = dp[j-1]
		}
		if needle[i] == needle[j] {
			dp[i] = j + 1
		} else {
			dp[i] = j
		}
	}

	return dp
}