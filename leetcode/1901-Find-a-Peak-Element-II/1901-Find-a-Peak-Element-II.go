package leetcode

import "math"

func findPeakGrid(mat [][]int) []int {
	m := len(mat)

	leftX, rightX := 0, m-1
	for leftX < rightX {
		midX := leftX + (rightX - leftX) / 2
		midY, midMax := maxElement(mat[midX])
		leftMax, rightMax := -1, -1
		if midX > 0 {
			_, leftMax = maxElement(mat[midX-1])
		}
		if midX < m - 1 {
			_, rightMax = maxElement(mat[midX+1])
		}

		largest := max(midMax, max(leftMax, rightMax))

		if largest == midMax {
			return []int{midX, midY}
		} else if largest == leftMax {
			rightX = midX - 1
		} else {
			leftX = midX + 1
		}

	}

	leftY, _ := maxElement(mat[leftX])
	return []int{leftX, leftY}
}

func maxElement(nums []int) (int, int) {
	idx, res := -1, math.MinInt
	for i, x := range nums {
		if x > res {
			idx, res = i, x
		}
	}

	return idx, res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}