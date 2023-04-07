package leetcode

import "math"

func preimageSizeFZF(k int) int {
	low := lowerBound(k)
	high := upperBound(k)

	return high - low - 1
}

// find the last number < target (trailing zeroes < k)
func lowerBound(k int) int {
	left, right := 0, math.MaxInt

	for left < right {
		mid := right - (right-left)/2
		if trailingZeroes(mid) < k {
			left = mid
		} else {
			right = mid - 1
		}
	}

	// post-processing
	if trailingZeroes(right) < k {
		return right
	}
	return right - 1

}

// find the first number > target (trailing zeroes > k)
func upperBound(k int) int {
	left, right := 0, math.MaxInt

	for left < right {
		mid := left + (right-left)/2
		if trailingZeroes(mid) <= k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// post-processing
	if trailingZeroes(left) > k {
		return left
	}
	return left + 1
}

func trailingZeroes(n int) int {
	res := 0
	divisor := 5
	for divisor <= n {
		res += n / divisor
		divisor *= 5
	}
	return res
}
