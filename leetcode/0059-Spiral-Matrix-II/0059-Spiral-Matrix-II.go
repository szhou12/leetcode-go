package leetcode

func generateMatrix(n int) [][]int {
	num := 1
	target := n * n

	top, bottom := 0, n-1 // top/bottom row index
	left, right := 0, n-1 // left/right col index

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for num <= target {
		// Fill top row
		for j := left; j <= right; j++ { // 注意终止条件是 <=
			matrix[top][j] = num
			num++
		}
		top++

		// Fill right col
		for i := top; i <= bottom; i++ { // 注意终止条件是 <=
			matrix[i][right] = num
			num++
		}
		right--

		// Fill bottom row
		for j := right; j >= left; j-- { // 注意终止条件是 <=
			matrix[bottom][j] = num
			num++
		}
		bottom--

		// Fill left col
		for i := bottom; i >= top; i-- { // 注意终止条件是 <=
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}
