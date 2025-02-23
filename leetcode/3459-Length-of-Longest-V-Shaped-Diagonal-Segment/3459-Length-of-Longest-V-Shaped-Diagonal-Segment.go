package leetcode

// need to maintain the order cuz only clockwise 90-degree turn is allowed
var dirs = [][]int{
	{-1, -1}, // up left
	{-1, 1}, // up right
	{1, 1}, // down right
	{1, -1}, // down left
}

// need to declare memo globally otherwise will TLE
var memo [500][500][4][2]int

func lenOfVDiagonal(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	memo = [500][500][4][2]int{}

	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// start at anywhere with 1
			if grid[i][j] == 1 {
				res = max(res, dfs(grid, i, j, 0, 1))
				res = max(res, dfs(grid, i, j, 1, 1))
				res = max(res, dfs(grid, i, j, 2, 1))
				res = max(res, dfs(grid, i, j, 3, 1))
			}
		}
	}

	return res
}

func dfs(grid [][]int, i int, j int, d int, turn int) int {
	// early return
	if memo[i][j][d][turn] != 0 {
		return memo[i][j][d][turn]
	}

	res := 1 // arriving at current cell adds 1 point

	// Two branches to make the next move: 
	// 1. keep the same direction
	di, dj := i + dirs[d][0], j + dirs[d][1]
	if inBound(grid, di, dj) && canContinue(grid[i][j], grid[di][dj]) {
		res = max(res, 1 + dfs(grid, di, dj, d, turn)) // 1 point for arriving at cur cell + points continuing the path
	}

	// 2. make a 90-degree clockwise turn
	if turn == 1 {
		newD := (d + 1) % 4
		di, dj := i + dirs[newD][0], j + dirs[newD][1]
		if inBound(grid, di, dj) && canContinue(grid[i][j], grid[di][dj]) {
			res = max(res, 1 + dfs(grid, di, dj, newD, 0)) // 1 point for arriving at cur cell + points continuing the path
		}
	}

	// udpate memo
	memo[i][j][d][turn] = res

	return res
}

func inBound(grid [][]int, i, j int) bool {
	m, n := len(grid), len(grid[0])
	return 0 <= i && i < m && 0 <= j && j < n
}

func canContinue(a, b int) bool {
	if a == 1 {
		return b == 2
	}

	if a == 2 {
		return b == 0
	}

	if a == 0 {
		return b == 2
	}

	return false
}


