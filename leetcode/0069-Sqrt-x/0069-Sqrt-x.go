package leetcode

// Binary Search
func mySqrt(x int) int {
	left := 0
	right := x

	for left < right {
		mid := right - (right-left)/2
		midVal := mid * mid
		if midVal == x {
			return mid
		} else if midVal < x {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}
