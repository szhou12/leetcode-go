package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// corner case
	if m == 0 || n == 0 {
		return false
	}

	left, right := 0, m*n-1

	for left < right {
		mid := left + (right - left) / 2
		pivot := matrix[mid/n][mid%n]
		if pivot == target {
			return true
		} else {
			if pivot < target {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	// post-processing: may not going into the loop
	if matrix[left/n][left%n] == target {
		return true
	}
	return false

}