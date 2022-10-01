package leetcode

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

/************* DFS ****************/
// 逆向思维: 从低往高处回流
func pacificAtlanticDFS(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	pacific, atlantis := make([][]bool, rows), make([][]bool, rows)
	for i := 0; i < rows; i++ {
		pacific[i] = make([]bool, cols)
		atlantis[i] = make([]bool, cols)
	}

	// pacific遍历O(rows*cols), atlantis遍历O(rows*cols)
	for i := 0; i < rows; i++ {
		dfs(heights, i, 0, &pacific, heights[i][0])            // 左边界回流
		dfs(heights, i, cols-1, &atlantis, heights[i][cols-1]) // 右边界回流
	}
	for j := 0; j < cols; j++ {
		dfs(heights, 0, j, &pacific, heights[0][j])            // 上边界回流
		dfs(heights, rows-1, j, &atlantis, heights[rows-1][j]) // 下边界回流
	}

	res := make([][]int, 0)
	// O(rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if pacific[i][j] && atlantis[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func dfs(heights [][]int, x int, y int, visited *[][]bool, prevHeight int) {
	// 3 Cases of Early Stopping
	// Case 1: out of bound
	if !inBoard(heights, x, y) {
		return
	}
	// Case 2: 上一层height > 当前height，无法回流至当前cell
	if prevHeight > heights[x][y] {
		return
	}
	// Case 3: 当前cell已经标记为可行解
	if (*visited)[x][y] {
		return
	}

	(*visited)[x][y] = true

	// recursion: 分别向四个方向走
	dfs(heights, x+1, y, visited, heights[x][y])
	dfs(heights, x-1, y, visited, heights[x][y])
	dfs(heights, x, y+1, visited, heights[x][y])
	dfs(heights, x, y-1, visited, heights[x][y])
}

func inBoard(heights [][]int, x int, y int) bool {
	return x >= 0 && x < len(heights) && y >= 0 && y < len(heights[0])
}

/************* BFS ****************/
func pacificAtlanticBFS(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	pacific, atlantis := make([][]bool, rows), make([][]bool, rows)
	for i := 0; i < rows; i++ {
		pacific[i] = make([]bool, cols)
		atlantis[i] = make([]bool, cols)
	}

	// pacific遍历O(rows*cols), atlantis遍历O(rows*cols)
	for i := 0; i < rows; i++ {
		bfs(heights, i, 0, &pacific)       // 左边界回流
		bfs(heights, i, cols-1, &atlantis) // 右边界回流
	}
	for j := 0; j < cols; j++ {
		bfs(heights, 0, j, &pacific)       // 上边界回流
		bfs(heights, rows-1, j, &atlantis) // 下边界回流
	}

	res := make([][]int, 0)
	// O(rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if pacific[i][j] && atlantis[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func bfs(heights [][]int, x int, y int, visited *[][]bool) {
	queue := [][]int{{x, y}}
	(*visited)[x][y] = true

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for i := 0; i < 4; i++ {
				dx := cur[0] + dir[i][0]
				dy := cur[1] + dir[i][1]
				if !inBoard(heights, dx, dy) {
					continue
				}
				if (*visited)[dx][dy] {
					continue
				}
				if heights[cur[0]][cur[1]] > heights[dx][dy] {
					continue
				}
				queue = append(queue, []int{dx, dy})
				(*visited)[dx][dy] = true
			}
		}
	}
}
