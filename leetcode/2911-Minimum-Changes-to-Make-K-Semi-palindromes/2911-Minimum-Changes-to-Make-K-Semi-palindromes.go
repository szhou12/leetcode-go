package leetcode

import "math"

func minimumChanges(s string, k int) int {
	n := len(s)

	// Step 1: Calculate # of changes to make semi-palindromes
	// substr[i][j] = # of changes to make any substring s[i:j] (inclusive) a semi-palindrome
	substr := make([][]int, n)
	for i := 0; i < n; i++ {
		substr[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			length := j-i+1
			substr[i][j] = math.MaxInt / 2 // initial state: maybe there is no way to make substring a semi-palindrome
			// find a proper divisor d
			for d := 1; d < length; d++ {
				if length % d != 0 {
					continue
				}
				totalChanges := 0
				// for each possible reminder r by mod d
				for r := 0; r < d; r++ {
					totalChanges += helper(s, i, j, d, r)
				}
				substr[i][j] = min(substr[i][j], totalChanges)
			}
		}
	}

	// Step 2: DP
	// dp[i][p] = min # of changes in make s[0:i] (inclusive) so it has p semi-palindrome substrings
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		// base case
		dp[i][1] = substr[0][i]
	}

	for i := 0; i < n; i++ {
		for p := 2; p <= k; p++ { // Note: p starts from 2 instead of 1 bc p-1=0 (0 substrings) has no meaning
			dp[i][p] = math.MaxInt / 2
			for j := 0; j < i; j++ {
				dp[i][p] = min(dp[i][p], dp[j][p-1]+substr[j+1][i])
			}
		}
	}

	return dp[n-1][k]

}

// calculate # of changes to make substring s[start:end] (inclusive) a semi-palindrome
func helper(s string, start, end, d, r int) int {
	count := 0

	left := start + r // first element in the substring whose reminder = r  by mod d
	right := end - (d-r) + 1 // last element in the substring whose reminder = r by mod d

	// squeeze from both ends towards the center
	for left < right {
		if s[left] != s[right] {
			count++
		}
		left+=d
		right-=d
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
