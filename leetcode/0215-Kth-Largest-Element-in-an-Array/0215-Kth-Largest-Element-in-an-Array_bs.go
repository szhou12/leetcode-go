package leetcode

import "math"

// Solution: Binary Search by Value
func findKthLargest_BS(nums []int, k int) int {
	left := math.MinInt32
	right := math.MaxInt32

	for left < right { // stop at left == right
		mid := right - (right-left)/2
		if count(nums, mid) >= k {
			left = mid // this can lead to dead loop when left=0, right=1, then mid = left + (right-left)/2 = 0. Hence, mid = right - (right-left)/2
		} else {
			right = mid - 1
		}
	}
	// return either left or right bc while loop stops at left == right
	return left
}

func count(nums []int, t int) int {
	// count the # of elements in nums that are >= t
	res := 0
	for _, num := range nums {
		if num >= t {
			res++
		}
	}

	return res
}
