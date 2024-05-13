package leetcode

import "math"

func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// dp[i][j] := minimun value in the rectangle of grid[0...i][0...j]
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt
		}
	}

	res := math.MinInt

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if 0 <= i-1 && i-1 < m {
				dp[i][j] = min(dp[i][j], dp[i-1][j])
			}
			if 0 <= j-1 && j-1 < n {
				dp[i][j] = min(dp[i][j], dp[i][j-1])
			}
			res = max(res, grid[i][j]-dp[i][j])
			dp[i][j] = min(dp[i][j], grid[i][j])

		}
	}

	return res

}
