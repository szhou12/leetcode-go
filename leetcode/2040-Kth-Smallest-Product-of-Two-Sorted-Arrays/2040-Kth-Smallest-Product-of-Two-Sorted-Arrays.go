package leetcode

import "math"

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	left := math.MinInt
	right := math.MaxInt

	for left < right {
		mid := left + (right-left)/2

		count := int64(countLessOrEqual(nums1, nums2, mid))
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return int64(left)
}

func countLessOrEqual(nums1 []int, nums2 []int, mid int) int {}

// find smallest index j s.t. nums[j] >= target
func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := right - (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right

}

// find largest index j s.t. nums[j] <= target
func upperBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
