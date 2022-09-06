package leetcode

import "math"

// My Solution - Easy to come up, BUT will cause TLE!!!
func longestIdealString_TLE(s string, k int) int {
	s = "#" + s
	dp := make([]string, len(s))
	dp[0] = ""
	dp[1] = s[1:2]
	resLen := math.MinInt

	for i := 2; i < len(s); i++ {
		dp[i] = s[i : i+1]

		curLen := 0
		for j := 1; j < i; j++ {
			if abs(int(s[i]), int(dp[j][len(dp[j])-1])) <= k {
				if curLen < len(dp[j]) {
					dp[i] = dp[j] + s[i:i+1]
					curLen = len(dp[j])
				}
			}
		}
		resLen = max(resLen, len(dp[i]))
	}

	return resLen
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func abs(a int, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}
