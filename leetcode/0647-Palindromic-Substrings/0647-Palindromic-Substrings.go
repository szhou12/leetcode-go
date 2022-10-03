package leetcode

// DP
// Time = O(n^2), Space = O(n^2)
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	res := 0
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			// Base Cases
			if i > j {
				continue
			} else if i == j {
				dp[i][j] = true
				res++ // 找到一个合法的回文串
				continue
			}
			// Recurrence
			if s[i] == s[j] {
				if j-i+1 == 2 || dp[i+1][j-1] {
					res++ // 找到一个合法的回文串
					dp[i][j] = true
				}
			}

		}
	}

	return res
}

// Two Pointers
// Time = O(n^2), Space = O(1)
func countSubstrings_TP(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += extend(s, i, i, len(s))
		res += extend(s, i, i+1, len(s))
	}

	return res
}

func extend(s string, l int, r int, n int) int {
	res := 0
	for l >= 0 && r < n && s[l] == s[r] {
		res++
		l--
		r++
	}
	return res
}
