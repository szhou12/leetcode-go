package leetcode

// My Solution
func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	res := 0
	for i := 0; i < m; i++ { // O(m)
		for j := 0; j < n; j++ { // O(n)
			if grid[i][j] == 1 {
				zeros := 0
				for _, dir := range dirs { // O(4)
					dx := i + dir[0]
					dy := j + dir[1]
					if outOfBound(grid, dx, dy) || grid[dx][dy] == 0 {
						zeros++
					}

				}
				res += zeros
			}
		}
	}
	return res
}

func outOfBound(grid [][]int, x int, y int) bool {
	m, n := len(grid), len(grid[0])
	if 0 <= x && x < m && 0 <= y && y < n {
		return false
	}
	return true
}
