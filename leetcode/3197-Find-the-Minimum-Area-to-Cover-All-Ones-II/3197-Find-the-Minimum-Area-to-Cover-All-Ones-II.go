package leetcode

import "math"

func minimumSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := math.MaxInt

	var area1, area2, area3 int

	// case 1
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			area1 = minimumArea(grid, 0, 0, i-1, n-1)
			area2 = minimumArea(grid, i, 0, j-1, n-1)
			area3 = minimumArea(grid, j, 0, m-1, n-1)
			res = min(res, area1+area2+area3)
		}
	}

	// case 2
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			area1 = minimumArea(grid, 0, 0, m-1, i-1)
			area2 = minimumArea(grid, 0, i, m-1, j-1)
			area3 = minimumArea(grid, 0, j, m-1, n-1)
			res = min(res, area1+area2+area3)
		}
	}

	// case 3
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area1 = minimumArea(grid, 0, 0, i-1, j-1)
			area2 = minimumArea(grid, 0, j, i-1, n-1)
			area3 = minimumArea(grid, i, 0, m-1, n-1)
			res = min(res, area1+area2+area3)
		}
	}

	// case 4
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area1 = minimumArea(grid, 0, 0, i-1, j-1)
			area2 = minimumArea(grid, i, 0, m-1, j-1)
			area3 = minimumArea(grid, 0, j, m-1, n-1)
			res = min(res, area1+area2+area3)
		}
	}

	// case 5
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area1 = minimumArea(grid, 0, 0, i-1, n-1)
			area2 = minimumArea(grid, i, 0, m-1, j-1)
			area3 = minimumArea(grid, i, j, m-1, n-1)
			res = min(res, area1+area2+area3)
		}
	}

	// case 6
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area1 = minimumArea(grid, 0, 0, m-1, j-1)
			area2 = minimumArea(grid, 0, j, i-1, n-1)
			area3 = minimumArea(grid, i, j, m-1, n-1)
			res = min(res, area1+area2+area3)
		}
	}

	return res
}

/*
Modified version of 3195
top-left: (a, b)
bottom-right: (c, d)
*/
func minimumArea(grid [][]int, a, b, c, d int) int {
	// check if a normal rectangle
	if !(a <= c && b <= d) {
		return math.MaxInt / 3 // divide 3 to avoid overflow in (area1+area2+area3)
	}

	m, n := len(grid), len(grid[0])

	minRow, minCol := m, n
	maxRow, maxCol := -1, -1

	ones := 0
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			if grid[i][j] == 1 {
				minRow, maxRow = min(minRow, i), max(maxRow, i)
				minCol, maxCol = min(minCol, j), max(maxCol, j)
				ones++
			}
		}
	}

	// check if there is no 1
	if ones == 0 {
		return math.MaxInt / 3 // divide 3 to avoid overflow in (area1+area2+area3)
	}

	return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}
