package leetcode

import "math"

func minimumMoves(grid [][]int) int {
	res := math.MaxInt
	dfs(0, 0, grid, &res)
	return res
}

func dfs(cur int, moves int, grid [][]int, res *int) {
	// early stopping
	if moves >= *res {
		return
	}

	// base case: all cells are checked
	if cur == 9 {
		*res = min(*res, moves)
		return
	}

	// current cell (i, j)
	i, j := cur/3, cur%3
	if grid[i][j] != 0 {
		dfs(cur+1, moves, grid, res)
		return
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] <= 1 {
				continue
			}
			grid[x][y]--
			grid[i][j]++
			dfs(cur+1, moves+abs(x-i)+abs(y-j), grid, res)
			grid[x][y]++
			grid[i][j]--
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
