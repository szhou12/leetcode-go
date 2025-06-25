package leetcode

var dir = [][]int {
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	res := 0

	queue := make([][]int, 0) // [[x, y, shortest time to reach]]

	// start nodes
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == 2 {
				queue = append(queue, []int{x, y, 0})
			}
		}
	}

	// Loop
	for len(queue) > 0 {
		// pop the current
		cur := queue[0]
		queue = queue[1:]
		x, y, t := cur[0], cur[1], cur[2]
		res = max(res, t)

		// make the next move
		for _, d := range dir {
			dx, dy := x + d[0], y + d[1]
			// check if out of bound
			if !inBoard(dx, dy, grid) {
				continue
			}
			// check if rotten or empty
			if grid[dx][dy] == 2 || grid[dx][dy] == 0 {
				continue
			}

			queue = append(queue, []int{dx, dy, t + 1})
			grid[dx][dy] = 2 // mark as 'visited' immediately once enqueued
		}

	}

	// post-processing: check if there is any fresh left
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == 1 {
				return -1
			}
		}
	}

	return res
}

func inBoard(x, y int, grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	return 0 <= x && x < m && 0 <= y && y < n
}