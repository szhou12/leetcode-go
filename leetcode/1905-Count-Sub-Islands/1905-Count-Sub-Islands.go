package leetcode

var dirs = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	// 突破点: 提前沉没掉不符合要求的岛屿 (i.e. grid2中是陆地，但grid1中已经是水的岛屿)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 && !visited[i][j] {
				dfs(grid2, &visited, i, j)
			}
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && !visited[i][j] {
				dfs(grid2, &visited, i, j)
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
	if grid[x][y] == 0 {
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
