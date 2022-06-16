package leetcode

import "math"

// My solution: 2D space O(n*m)
func minimumTotal(triangle [][]int) int {
	// corner case
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	n := len(triangle)
	dp := make([][]int, n)
	for row := 0; row < n; row++ {
		m := len(triangle[row])
		dp[row] = make([]int, m)
		if row == 0 {
			dp[row][0] = triangle[0][0]
		}
	}

	for row := 1; row < n; row++ {
		for col := 0; col < len(triangle[row]); col++ {
			if col == 0 {
				dp[row][col] = dp[row-1][col] + triangle[row][col]
			} else if col == len(triangle[row])-1 {
				dp[row][col] = dp[row-1][col-1] + triangle[row][col]
			} else {
				dp[row][col] = min(dp[row-1][col], dp[row-1][col-1]) + triangle[row][col]
			}
		}
	}

	res := math.MaxInt
	lastRow := len(dp) - 1
	for col := 0; col < len(dp[lastRow]); col++ {
		res = min(res, dp[lastRow][col])
	}

	return res

}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// Optimal Solution: space = O(n)
// dp[] 只需要记录最后一行，对应index上叠加value即可
func minimumTotal_opt(triangle [][]int) int {
	// corner case
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	dp := make([]int, len(triangle[len(triangle)-1]))
	res := math.MaxInt
	dp[0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == 0 {
				// 最左边
				dp[j] += triangle[i][0]
			} else if j == len(triangle[i])-1 {
				// 最右边
				dp[j] += dp[j-1] + triangle[i][j]
			} else {
				// 中间
				dp[j] += min(dp[j-1], dp[j]) + triangle[i][j]
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		res = min(res, dp[i])
	}
	return res

}
