package leetcode

import "math"

var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// person[i][j]: earliest time/shoretest path for the person to arrive at (i,j)
	person := bfs(grid, 0)
	// fire[i][j]: earliest time/shoretest path for the fire to arrive at (i,j)
	fire := bfs(grid, 1)

	// person can't arrive at destination
	if person[m-1][n-1] == math.MaxInt {
		return -1
	}
	// person always arrives later than fire
	if person[m-1][n-1] > fire[m-1][n-1] {
		return -1
	}
	// fire can't arrive at destination
	if fire[m-1][n-1] == math.MaxInt {
		return 1e9
	}

	// max # of days that person can delay before starting
	d := fire[m-1][n-1] - person[m-1][n-1]
	// look left cell and top cell to destination
	// case 1: if fire arrives at left cell and top cell at the same time,
	//			person will meet fire before destination
	if fire[m-1][n-2] == fire[m-2][n-1] {
		return d - 1
	}
	// case 2: if fire arrives at LEFT cell earlier than top cell,
	// 			i) person can wait d days and arrive at destination from TOP cell if TOP cell is on person's shortest path
	//			ii) person waits at most d-1 days, otherwise
	if fire[m-2][n-1] > fire[m-1][n-2] { // fire go from left cell
		if person[m-1][n-1]-person[m-2][n-1] == 1 { // person go from top cell
			return d
		}
	}
	// case 3: if fire arrives at TOP cell earlier than left cell,
	//			i) person can wait d days and arrive at destination from LEFT cell if LEFT cell is on person's shortest path
	//			Ii) person waits at most d-1 days, otherwise
	if fire[m-1][n-2] > fire[m-2][n-1] { // fire go from top cell
		if person[m-1][n-1]-person[m-1][n-2] == 1 { // person go from left cell
			return d
		}
	}

	return d - 1

}

func bfs(grid [][]int, role int) [][]int {
	m, n := len(grid), len(grid[0])

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = math.MaxInt
		}
	}

	queue := make([][]int, 0)

	// start nodes
	// if bfs on person, start from (0,0)
	// if bfs on fire, start from all fire cells
	if role == 0 {
		queue = append(queue, []int{0, 0})
		res[0][0] = 0
	} else {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 1 {
					queue = append(queue, []int{i, j})
					res[i][j] = 0
				}
			}
		}
	}

	// loop
	// no need for level order traversal
	for len(queue) > 0 {
		// cur
		cur := queue[0]
		queue = queue[1:]
		x, y := cur[0], cur[1]
		// if visits destination at first time, end loop
		if x == m-1 && y == n-1 {
			break
		}

		// make the next move
		for _, d := range dir {
			dx, dy := x+d[0], y+d[1]
			// check if out of bound
			if !(0 <= dx && dx < m && 0 <= dy && dy < n) {
				continue
			}
			// check if hit wall
			if grid[dx][dy] == 2 {
				continue
			}
			// check if already visited
			if res[dx][dy] != math.MaxInt {
				continue
			}
			queue = append(queue, []int{dx, dy})
			res[dx][dy] = res[x][y] + 1
		}
	}
	return res
}
