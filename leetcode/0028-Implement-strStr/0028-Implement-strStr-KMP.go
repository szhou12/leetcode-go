package leetcode

// Step 1: the longest prefix suffix (lsp) array of p
//
//	lsp[i] = the max length k s.t. p[0:k-1] = p[i-k+1:i]
//
// Step 2: pattern matching
//
//	dp[i] = the max length x s.t. s[i-x+1:i] = p[0:x-1]
//	when dp[i] = length(p) => match!
func strStr_KMP(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	// edge cases
	// try to find an empty pattern in a non-empty target
	if m == 0 {
		return 0
	}
	// try to find a non-empty pattern in an empty target
	if n == 0 {
		return -1
	}

	// build lsp for pattern string = needle
	lsp := preprocess(needle)

	// dp[i] = max length x of suffix in haystack[:i] s.t. prefix needle[0:x-1] == suffix haystack[i-x+1:i] (双闭区间)
	dp := make([]int, n)

	// Base case
	if needle[0] == haystack[0] {
		dp[0] = 1
	}

	// Early return
	// Check first char: if length(pattern) == 1 & dp[0] == 1
	// MUST have this bc the rest of code is for i > 0
	if m == 1 && dp[0] == 1 {
		return 0
	}

	// Recurrence
	for i := 1; i < n; i++ {
		x := dp[i-1]
		for x > 0 && needle[x] != haystack[i] {
			x = lsp[x-1]
		}
		if needle[x] == haystack[i] {
			dp[i] = x + 1
		} else {
			dp[i] = x
		}

		// check if haystack[i-j+1:i] (i.e. dp[i]) is a needle
		if dp[i] == m {
			return i - m + 1
		}
	}

	return -1

}

func preprocess(needle string) []int {
	n := len(needle)

	// dp[i] = max length j s.t. prefix needle[0:j-1] == suffix needle[i-j+1:i] (双闭区间)
	dp := make([]int, n)

	// base case: 因为我们要求真前缀和真后缀，只有一个字符的时候不存在真前缀和真后缀，所以dp[0] = 0
	dp[0] = 0

	// recurrence
	for i := 1; i < n; i++ {
		j := dp[i-1]
		// j 是上一轮匹配的长度, 恰恰好的是, neddle[j]也表示needle的前缀中这一轮即将要匹配的字符 i.e. needle[0:j-1]的下一个字符
		for j > 0 && needle[i] != needle[j] {
			// j-1 这里表示index
			// j 这里表示新长度j'
			j = dp[j-1]
		}
		if needle[i] == needle[j] {
			dp[i] = j + 1
		} else {
			// 表示没有匹配到任何前缀，此时前面会跳出只能因为 j==0, 也就是匹配长度=0
			dp[i] = j
		}
	}

	return dp
}