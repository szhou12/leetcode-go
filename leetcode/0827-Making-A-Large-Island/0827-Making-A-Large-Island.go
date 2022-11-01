package leetcode

var dirs = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func largestIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	allLand := true
	islands := make(map[int]int) // key=岛屿编号, value=该岛屿面积
	mark := 2                    // 编号从2开始因为原grid由 0,1组成
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				allLand = false
			}
			if grid[i][j] == 1 && !visited[i][j] {
				area := 0
				dfs(&grid, &visited, &area, i, j, mark)
				islands[mark] = area
				mark++ // 每找完一个岛屿, 编号累加
			}
		}
	}

	// early stopping: 如果全是陆地, 压根儿不用找了, 最大面积就是整个grid
	if allLand {
		return m * n
	}

	// check all possible postions to turn 0 to 1
	largest := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 { // 尝试每个0的空位, 看它能不能连接上、下、左、右可能的岛屿
				visitedIsland := make(map[int]bool) // 用来记录哪些邻居岛屿已经连接过, 避免重复连接
				grid[i][j] = 1
				count := 1
				for _, dir := range dirs {
					x := i + dir[0]
					y := j + dir[1]
					if !inBoard(&grid, x, y) || grid[x][y] == 0 {
						continue
					}
					if visitedIsland[grid[x][y]] {
						continue
					}
					count += islands[grid[x][y]]     // 把相邻四面的岛屿数量加起来
					visitedIsland[grid[x][y]] = true // 标记该岛屿已经连接过

				}
				largest = max(count, largest)
			}
		}
	}
	return largest

}

func dfs(grid *[][]int, visited *[][]bool, area *int, x int, y int, mark int) {
	// base cases
	if !inBoard(grid, x, y) {
		return
	}
	if (*visited)[x][y] {
		return
	}
	if (*grid)[x][y] == 0 {
		return
	}

	*area++
	(*visited)[x][y] = true
	(*grid)[x][y] = mark // 染色成当前岛屿的编号

	for _, dir := range dirs {
		dx := x + dir[0]
		dy := y + dir[1]
		dfs(grid, visited, area, dx, dy, mark)
	}
}

func inBoard(grid *[][]int, x int, y int) bool {
	m, n := len(*grid), len((*grid)[0])
	return (0 <= x && x < m) && (0 <= y && y < n)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
