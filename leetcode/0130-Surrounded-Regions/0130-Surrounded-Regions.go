package leetcode

var dir = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

/************* DFS ****************/
func solveDFS(board [][]byte) {
	m, n := len(board), len(board[0])

	// Step 1: 找到在边界的O区域
	// 1a) From left, right
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(&board, i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(&board, i, n-1)
		}
	}
	// 1b) From top, bottom
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(&board, 0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(&board, m-1, j)
		}
	}

	// Step 2: convert 'O' areas to 'X', 'A' areas to 'O'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board *[][]byte, x int, y int) {
	// Base Cases
	// 1. out of bound
	if !inBoard(board, x, y) {
		return
	}
	// 2. Not 'O'
	if (*board)[x][y] != 'O' {
		return
	}
	// 当前层干点什么事儿
	(*board)[x][y] = 'A'
	// Recursion
	for i := 0; i < len(dir); i++ {
		dx := x + dir[i][0]
		dy := y + dir[i][1]
		dfs(board, dx, dy)
	}
}

func inBoard(board *[][]byte, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(*board) && y < len((*board)[0])
}

/************* BFS ****************/
func solveBFS(board [][]byte) {
	m, n := len(board), len(board[0])

	// Step 1: 找到在边界的O区域
	// 1a) From left, right
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			bfs(&board, i, 0)
		}
		if board[i][n-1] == 'O' {
			bfs(&board, i, n-1)
		}
	}
	// 1b) From top, bottom
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			bfs(&board, 0, j)
		}
		if board[m-1][j] == 'O' {
			bfs(&board, m-1, j)
		}
	}

	// Step 2: convert 'O' areas to 'X', 'A' areas to 'O'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}

func bfs(board *[][]byte, x int, y int) {
	queue := [][]int{{x, y}}
	(*board)[x][y] = 'A' // 只要加入队列，立刻标记

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for i := 0; i < 4; i++ {
				dx := cur[0] + dir[i][0]
				dy := cur[1] + dir[i][1]
				// Base Cases
				if !inBoard(board, dx, dy) {
					continue
				}
				if (*board)[dx][dy] != 'O' {
					continue
				}

				queue = append(queue, []int{dx, dy})
				(*board)[dx][dy] = 'A' // 只要加入队列，立刻标记
			}
		}
	}
}
