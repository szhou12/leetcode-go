package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// Base Cases
	i, j := 0, 0
	for i < m {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
		i++
	}
	for j < n {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
		j++
	}

	// Reccurrence
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
