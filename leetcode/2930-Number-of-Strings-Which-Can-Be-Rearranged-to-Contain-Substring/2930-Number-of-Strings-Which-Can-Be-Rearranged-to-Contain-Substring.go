package leetcode

var M = int(1e9+7)
func stringCount(n int) int {
	dp := make([][][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][][]int, 2)
		for a := 0; a <= 1; a++ {
			dp[i][a] = make([][]int, 3)
			for b := 0; b <= 2; b++ {
				dp[i][a][b] = make([]int, 2)
			}
		}
	}

	// base case 
	dp[0][0][0][0] = 1

	// recurrence
	for i := 1; i <= n; i++ {
		for a := 0; a <= 1; a++ {
			for b := 0; b <= 2; b++ {
				for c := 0; c <= 1; c++ {
					// compute dp[i][a][b][c]
					for ch := 'a'; ch <= 'z'; ch++ {
						if ch == 'l' && a == 1 {
							dp[i][1][b][c] += dp[i-1][0][b][c]
						} else if ch == 'e' && b == 1 {
							dp[i][a][1][c] += dp[i-1][a][0][c]
						} else if ch == 'e' && b == 2 {
							dp[i][a][2][c] += dp[i-1][a][1][c]
						} else if ch == 't' && c == 1 {
							dp[i][a][b][1] += dp[i-1][a][b][0]
						} else {
							dp[i][a][b][c] += dp[i-1][a][b][c]
						}
						dp[i][a][b][c] %= M
					}
				}
			}
		}
	}

	return dp[n][1][2][1]
}