package leetcode

var dirs = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	// 突破点: 提前沉没掉不符合要求的岛屿 (i.e. 贴着边界的岛屿)
	// 沉没掉贴着上、下边界的岛屿
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 && !visited[i][0] {
			dfs(grid, &visited, i, 0)
		}
		if grid[i][n-1] == 0 && !visited[i][n-1] {
			dfs(grid, &visited, i, n-1)
		}
	}
	// 沉没掉贴着左、右边界的岛屿
	for j := 0; j < n; j++ {
		if grid[0][j] == 0 && !visited[0][j] {
			dfs(grid, &visited, 0, j)
		}
		if grid[m-1][j] == 0 && !visited[m-1][j] {
			dfs(grid, &visited, m-1, j)
		}
	}

	// 再数合法的岛屿个数
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && !visited[i][j] {
				dfs(grid, &visited, i, j)
				res++
			}
		}
	}
	return res

}

func dfs(grid [][]int, visited *[][]bool, x int, y int) {
	// base cases
	if !inBoard(grid, x, y) {
		return
	}
	if (*visited)[x][y] {
		return
	}
	if grid[x][y] == 1 {
		return
	}

	(*visited)[x][y] = true

	for _, dir := range dirs {
		dx := x + dir[0]
		dy := y + dir[1]
		dfs(grid, visited, dx, dy)
	}
}

func inBoard(grid [][]int, x int, y int) bool {
	m, n := len(grid), len(grid[0])
	return (0 <= x && x < m) && (0 <= y && y < n)
}
