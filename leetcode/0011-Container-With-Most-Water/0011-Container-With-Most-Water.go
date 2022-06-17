package leetcode

func maxArea(height []int) int {
	n := len(height)
	res := 0
	left := 0
	right := n - 1

	for left < right {
		// calculate cur area
		curArea := min(height[left], height[right]) * (right - left)
		res = max(res, curArea)

		// move pointers
		// WHY move lower-height side?
		// bc width strictly goes down,
		// height is min(h[left], r[right])
		// the only chance that Area increases when width decreases is to move lower-height side
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
