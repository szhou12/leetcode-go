package leetcode

func generate(numRows int) [][]int {
	var result [][]int

	// i = ith row
	for i := 0; i < numRows; i++ {
		row := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, result[i-1][j-1]+result[i-1][j])
			}
		}
		result = append(result, row)
	}
	return result
}
