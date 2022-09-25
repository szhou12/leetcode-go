package leetcode

// Optimal Solution - DP
// dp[i][j] = # unique paths from grid[0][0] to grid[i-1][j-1]
// base cases:
// dp[0][0] = 1
// dp[0][j] = 1
// dp[i][0] = 1
// recurrence:
// dp[i][j] = dp[i-1][j] + dp[i][j-1]
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// Base Cases
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		dp[i][0] = 1 //只有一列时，只有向下走一种方法
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1 //只有一行时，只有向右走一种方法
	}

	// Recurrence
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j]
		}
	}

	return dp[m-1][n-1]
}

// DFS暴力解 - 会超时!!!
// # levels = (m-1) + (n-1)
// # branches = go Right / go Down
// Must have (m-1) Downs & (n-1) Rights
func uniquePaths_DFS(m int, n int) int {
	return dfs(0, 0, m-1, n-1)
}

// return # of unique paths
func dfs(down int, right int, m int, n int) int {
	// Base case 1: 任意一个方向越界了，能发生越界说明，这个方向走多了，而另一个方向还没用完, 显然这不是一个解
	if down > m || right > n {
		return 0
	}
	// Base case 2: 两个方向同时用完了对应步数，找到了一个解
	if down == m && right == n {
		return 1
	}

	// branch 1: go right at current step
	r := dfs(down, right+1, m, n)
	// branch 2: go down at current step
	d := dfs(down+1, right, m, n)

	return r + d
}
