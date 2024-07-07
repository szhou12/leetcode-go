package leetcode

func numberOfSubmatrices(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	prefixSumX := make([][]int, m+1)
	prefixSumY := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		prefixSumX[i] = make([]int, n+1)
		prefixSumY[i] = make([]int, n+1)
	}

	res := 0

	// prefixSum[0][0:n+1] = 0
	// prefixSum[0:m+1][0] = 0
	// prefixSum[i+1][j+1] = grid[i][j]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// inclusion-exclusion principle
			prefixSumX[i+1][j+1] = prefixSumX[i][j+1] + prefixSumX[i+1][j] - prefixSumX[i][j]
			prefixSumY[i+1][j+1] = prefixSumY[i][j+1] + prefixSumY[i+1][j] - prefixSumY[i][j]

			// prefix sum in the rectangle from grid[0][0] to grid[i][j]
			switch grid[i][j] {
			case 'X':
				prefixSumX[i+1][j+1] += 1
			case 'Y':
				prefixSumY[i+1][j+1] += 1
			}

			// count valid submatrices
			if prefixSumX[i+1][j+1] > 0 && prefixSumX[i+1][j+1] == prefixSumY[i+1][j+1] {
				res++
			}
		}
	}

	return res

}
