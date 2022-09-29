package leetcode

// DP
func largestRectangleArea(heights []int) int {
	n := len(heights)
	lmin := make([]int, n) // lmin[i]: index of first bar [i-1...0] shorter than bar i
	rmin := make([]int, n) // rmin[i]: index of firt bar [i+1...n-1] shorter than bar i

	lmin[0] = -1
	for i := 1; i < n; i++ {
		tempIdx := i - 1
		for tempIdx >= 0 && heights[tempIdx] >= heights[i] {
			tempIdx = lmin[tempIdx]
		}
		lmin[i] = tempIdx
	}

	rmin[n-1] = n
	for i := n - 1; i >= 0; i-- {
		tempIdx := i + 1
		for tempIdx < n && heights[tempIdx] >= heights[i] {
			tempIdx = rmin[tempIdx]
		}
		rmin[i] = tempIdx
	}

	res := 0
	for i := 0; i < n; i++ {
		curArea := heights[i] * (rmin[i] - lmin[i] - 1) // 因为宽要去掉lmin和rmin
		res = max(res, curArea)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
