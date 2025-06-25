package leetcode

var dir = [][]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	inBound := func(x, y int) bool {
		return (0 <= x && x < m) && (0 <= y && y < n)
	}

	var dfs func(x, y, idx int) bool
	dfs = func(x, y, idx int) bool {
		// base cases
		// out of bound
		if !inBound(x, y) {
			return false
		}
		// already visited
		if visited[x][y] {
			return false
		}
		// letter doesn't match
		if word[idx] != board[x][y] {
			return false
		}
		// letter matches and it is the last letter
		if idx == len(word)-1 {
			return true
		}

		// recursion
		visited[x][y] = true

		canReach := false
		for _, d := range dir {
			dx, dy := x+d[0], y+d[1]

			// if out of bound
			if !inBound(dx, dy) {
				continue
			}

			// if already visited
			if visited[dx][dy] {
				continue
			}

			canReach = canReach || dfs(dx, dy, idx+1)
		}

		// backtracking!!!!
		visited[x][y] = false

		return canReach
	}

	// iterate every cell, every time we've found a cell whose letter matches the first character of word, we conduct dfs to see if it can form the word
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}
