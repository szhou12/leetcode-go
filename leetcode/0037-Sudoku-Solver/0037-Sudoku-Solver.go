package leetcode

func solveSudoku(board [][]byte) {
	dfs(&board, 0, 0)
}

// return true if there exists a soln; false o/w
func dfs(board *[][]byte, i int, j int) bool {
	m, n := 9, 9

	// 所有行都走完了，说明找到了一个解，触发Base Case
	if i == m {
		return true
	}

	// 当前行的9列都走完了, 则, 走到下一行第一个格子重新开始
	if j == n {
		return dfs(board, i+1, 0)
	}

	// 如果当前格子有预设数字，不用再找，直接跳到下一个格子
	if (*board)[i][j] != '.' {
		return dfs(board, i, j+1)
	}

	// 9 branches: 每一个branch试一个数字
	for ch := '1'; ch <= '9'; ch++ {
		// 如果当前试的数字不满足三大条件，尝试下一个数字
		if !isValid(board, i, j, byte(ch)) {
			continue
		}

		(*board)[i][j] = byte(ch)
		if dfs(board, i, j+1) { // 如果找到一个可行解，立即结束
			return true
		}
		(*board)[i][j] = '.'

	}

	// 穷举完 1~9，依然没有找到可行解，此路不通
	return false

}

func isValid(board *[][]byte, row int, col int, ch byte) bool {
	// 检查所在 行 是否有重复数字
	for j := 0; j < 9; j++ {
		if (*board)[row][j] == ch {
			return false
		}
	}

	// 检查所在 列 是否有重复数字
	for i := 0; i < 9; i++ {
		if (*board)[i][col] == ch {
			return false
		}
	}

	// 检查所在 3X3 box 是否有重复数字
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if (*board)[i][j] == ch {
				return false
			}
		}
	}
	return true
}
