package leetcode

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	top, bottom := 0, m-1
	left, right := 0, n-1
	num, target := 1, m*n
	var res []int

	for num <= target {
		// Fill top
		for j := left; j <= right && num <= target; j++ {
			res = append(res, matrix[top][j])
			num++
		}
		top++

		// Fill right
		for i := top; i <= bottom && num <= target; i++ {
			res = append(res, matrix[i][right])
			num++
		}
		right--

		// Fill bottom
		for j := right; j >= left && num <= target; j-- {
			res = append(res, matrix[bottom][j])
			num++
		}
		bottom--

		// Fill left
		for i := bottom; i >= top && num <= target; i-- {
			res = append(res, matrix[i][left])
			num++
		}
		left++
	}

	return res
}
