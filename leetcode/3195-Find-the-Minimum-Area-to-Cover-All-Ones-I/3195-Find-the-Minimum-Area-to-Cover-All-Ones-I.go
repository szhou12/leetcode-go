package leetcode

func minimumArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	minRow, minCol := m, n
	maxRow, maxCol := -1, -1

	ones := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ones += 1
				minRow = min(minRow, i)
				minCol = min(minCol, j)
				maxRow = max(maxRow, i)
				maxCol = max(maxCol, j)
			}
		}
	}

	// check if no 1's
	if ones == 0 {
		return 0
	}

	return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}
