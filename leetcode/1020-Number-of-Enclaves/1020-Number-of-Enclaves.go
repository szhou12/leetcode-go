package leetcode

// DFS
func numEnclaves_DFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// left, right
	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			dfs(&grid, i, 0)
		}
		if grid[i][n-1] == 1 {
			dfs(&grid, i, n-1)
		}
	}
	// top, bottom
	for j := 0; j < n; j++ {
		if grid[0][j] == 1 {
			dfs(&grid, 0, j)
		}
		if grid[m-1][j] == 1 {
			dfs(&grid, m-1, j)
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}

	return count
}

func dfs(grid *[][]int, x int, y int) {
	// Base Cases
	if !inBoard(grid, x, y) {
		return
	}
	if (*grid)[x][y] == 0 {
		return
	}

	// 当前层淹没陆地
	(*grid)[x][y] = 0

	// Recursion
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
	dfs(grid, x, y+1)
	dfs(grid, x, y-1)

}

func inBoard(grid *[][]int, x int, y int) bool {
	return x >= 0 && x < len(*grid) && y >= 0 && y < len((*grid)[0])
}
