package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {

	m := len(s1)
	n := len(s2)

	// corner case
	if m+n != len(s3) {
		return false
	}

	// MUST prepend a placeholder (占位符) bc index of string starts from 1 NOT 0
	s1 = "#" + s1
	s2 = "#" + s2
	s3 = "#" + s3

	// init a slice of slices of type boolean
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for i := 1; i <= m; i++ {
		dp[i][0] = (dp[i-1][0] && s1[i] == s3[i])
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = (dp[0][j-1] && s2[j] == s3[j])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if dp[i-1][j] && s3[i+j] == s1[i] {
				dp[i][j] = true
			} else if dp[i][j-1] && s3[i+j] == s2[j] {
				dp[i][j] = true
			}
		}
	}

	return dp[m][n]
}
