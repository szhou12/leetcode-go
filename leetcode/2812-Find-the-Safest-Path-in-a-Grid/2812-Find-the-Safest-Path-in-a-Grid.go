package leetcode

var dir = [][]int {
	{1, 0},
	{0,1},
	{-1,0},
	{0,-1},
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	
    // STEP 1: (BFS) count distance from each cell to nearest thief
	// grid[i][j] := distance from cell (i, j) to nearest thief + 1
	countDistance(grid)

	// STEP 2: (Binary Search) guess max safe distance + (BFS) check if valid
	left, right := 0, 2*n
	for left < right {
		mid := right - (right-left)/2
		if validDistance(grid, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left

}

func validDistance(grid [][]int, d int) bool {
	n := len(grid)

	// early stopping: if start node already can't take d
	if grid[0][0] - 1 < d {
		return false
	}

	queue := make([][]int, 0)

	visited :=make([][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, n)
	}

	// start node: (0, 0)
	queue = append(queue, []int{0, 0})

	// loop
	for len(queue) > 0 {
		// Actually no need to track level-order
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for k := 0; k < 4; k++ {
				dx := cur[0] + dir[k][0]
				dy := cur[1] + dir[k][1]

				// check if out of bound
				if !(0 <= dx && dx < n && 0 <= dy && dy < n) {
					continue
				}
				// check if already visited
				if visited[dx][dy] != 0 {
					continue
				}
				// check if given distance is too much
				if grid[dx][dy] - 1 < d {
					continue
				}

				// check if arrive at destination
				if dx == n-1 && dy == n-1 {
					return true
				}

				visited[dx][dy] = 1
				queue = append(queue, []int{dx, dy})
			}
		}
	}

	return false
}

func countDistance(grid [][]int) {
	n := len(grid)

	queue := make([][]int, 0)

	// start nodes: thief node
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	// loop
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			x, y := cur[0], cur[1]
			queue = queue[1:]

			for k := 0; k < 4; k++ {
				dx := cur[0] + dir[k][0]
				dy := cur[1] + dir[k][1]

				// check if out of bound
				if !(0 <= dx && dx < n && 0 <= dy && dy < n) {
					continue
				}
				// check if already visited
				if grid[dx][dy] != 0 {
					continue
				}

				grid[dx][dy] = grid[x][y] + 1
				queue = append(queue, []int{dx, dy})
			}
		}
	}
}