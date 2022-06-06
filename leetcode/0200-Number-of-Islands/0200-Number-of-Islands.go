package leetcode

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	// corner case
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
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(grid, &visited, i, j)
				res++
			}
		}
	}

	return res
}

func dfs(grid [][]byte, visited *[][]bool, x int, y int) {
	// base cases
	// case 1: out of board
	if !inBoard(grid, x, y) {
		return
	}
	// case 2: (x, y) is '0'/ocean
	if grid[x][y] == '0' {
		return
	}
	// case 3: (x, y) already visited
	if (*visited)[x][y] {
		return
	}

	// DFS at four directions
	// mark (x,y) visited
	(*visited)[x][y] = true
	for i := 0; i < 4; i++ {
		dx := x + dir[i][0]
		dy := y + dir[i][1]
		dfs(grid, visited, dx, dy)
	}
}

func inBoard(grid [][]byte, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
}
