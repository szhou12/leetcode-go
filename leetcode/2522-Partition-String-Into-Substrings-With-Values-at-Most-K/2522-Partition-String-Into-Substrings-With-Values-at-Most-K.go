package leetcode

import "strconv"

// DP
// dp[i]: minimum number of substrings in a good partition of s[1...i]
func minimumPartition(s string, k int) int {
	n := len(s)
	kLen := len(strconv.Itoa(k))

	for i := 0; i < len(s); i++ {
		if int(s[i]-'0') > k {
			return -1
		}
	}

	dp := make([]int, n+1)
	dp[0] = 0

	s = "#" + s

	for i := 1; i <= n; i++ {
		// find starting index j: i - j + 1 = kLen
		j := i - kLen + 1
		partition, _ := strconv.Atoi(s[max(0, j) : i+1])
		if j >= 1 && partition <= k {
			dp[i] = dp[j-1] + 1
		} else {
			dp[i] = dp[max(0, j)] + 1
		}
	}

	return dp[n]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
