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

/************* Union Find ****************/
func solveUF(board [][]byte) {
	uf := UnionFind{}
	uf.Init()

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}
			idx := i*n + j
			uf.father[idx] = idx
		}
	}

	// Union
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}
			cur := i*n + j
			for _, d := range dir {
				dx, dy := i+d[0], j+d[1]
				if !(0 <= dx && dx < m && 0 <= dy && dy < n) {
					continue
				}
				if board[dx][dy] == 'X' {
					continue
				}
				nei := dx*n + dy
				if uf.Find(cur) != uf.Find(nei) {
					uf.Union(cur, nei)
				}
			}
		}
	}

	// Record each family and mark families on borders
	family := make(map[int][]int)
	border := make(map[int]bool) // record family root who has border members
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}
			root := uf.Find(i*n + j)
			if _, ok := family[root]; !ok {
				family[root] = make([]int, 0)
			}
			family[root] = append(family[root], i*n+j)

			// check if (i, j) on border
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				border[root] = true
			}

		}
	}

	// Flip families that have no border members
	for root, members := range family {
		if border[root] {
			continue
		}
		for _, member := range members {
			i, j := member/n, member%n
			board[i][j] = 'X'
		}
	}
}

type UnionFind struct {
	father map[int]int
}

func (uf *UnionFind) Init() {
	uf.father = make(map[int]int)
}

func (uf *UnionFind) Union(x int, y int) {
	fx := uf.father[x]
	fy := uf.father[y]
	if fx < fy {
		uf.father[fy] = fx
	} else {
		uf.father[fx] = fy
	}
}

func (uf *UnionFind) Find(x int) int {
	if x != uf.father[x] {
		uf.father[x] = uf.Find(uf.father[x])
	}
	return uf.father[x]
}
