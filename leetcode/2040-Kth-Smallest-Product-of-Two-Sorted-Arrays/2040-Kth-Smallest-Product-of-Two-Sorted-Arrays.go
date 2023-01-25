package leetcode

import "math"

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	left := int(-1e5 * 1e5)
	right := int(1e5 * 1e5)

	for left < right { // O(logN)
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

/*
nums1[i] * nums2[j] <= mid
if nums1[i] > 0:
	find largest j s.t. nums2[j] <= floor(mid/nums1[i]) --> [0, j] // bc nums2 sorted
	count += (j-0+1)
if nums1[i] == 0:
    if mid >= 0: count += len(nums2) // all nums2 elements satisfy the requirement
    if mid < 0: count += 0 // no nums2 element satisfy the requirement
if nums1[i] < 0:
	find smallest j s.t. nums2[j] >= ceil(mid/nums1[i]) --> [j, n-1]
	count += (n-1-j+1)
*/
func countLessOrEqual(nums1 []int, nums2 []int, mid int) int {
	var res int

	// O(NlogN)
	for i := 0; i < len(nums1); i++ {
		if nums1[i] > 0 {
			y := math.Floor(float64(mid) / float64(nums1[i]))
			j := upperBound(nums2, int(y))
			res += j
		} else if nums1[i] == 0 {
			if mid < 0 {
				res += 0
			} else {
				res += len(nums2)
			}
		} else { // nums1[i] < 0
			y := math.Ceil(float64(mid) / float64(nums1[i]))
			j := lowerBound(nums2, int(y))
			res += (len(nums2) - 1 - j + 1)
		}
	}

	return res
}

// find first index j s.t. nums[j] >= target
func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if nums[left] >= target {
		return left
	} else {
		return left + 1
	}

}

// find first index j s.t. nums[j] > target
func upperBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] > target {
		return left
	} else {
		return left + 1
	}
}
