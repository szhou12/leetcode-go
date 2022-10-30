package leetcode

var dirs = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] { // find a new island
				count := 0
				dfs(grid, &visited, &count, i, j) // count how big this new island
				res = max(res, count)
			}
		}
	}

	return res
}

func dfs(grid [][]int, visited *[][]bool, count *int, x int, y int) {
	// Base Cases
	if !inBoard(grid, x, y) {
		return
	}
	if (*visited)[x][y] {
		return
	}
	if grid[x][y] == 0 {
		return
	}

	*count++
	(*visited)[x][y] = true
	for _, dir := range dirs {
		dx := x + dir[0]
		dy := y + dir[1]
		dfs(grid, visited, count, dx, dy)
	}
}

func inBoard(grid [][]int, x int, y int) bool {
	m, n := len(grid), len(grid[0])
	return (0 <= x && x < m) && (0 <= y && y < n)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
