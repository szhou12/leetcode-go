package leetcode

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	// Step 1: sort nums in ascending order
	sort.Ints(nums)

	res := 0

	// Step 2: for each element, find its pair
	for _, x := range nums {
		// [lower - x, upper - x]
		right := upperBound(nums, upper-x)
		left := lowerBound(nums, lower-x)
		count := right - left

		// itself cannot be its pair
		if lower <= x+x && x+x <= upper {
			count--
		}
		res += count
	}

	// divide by 2 because each pair is counted twice as i, j are interchangeable
	return int64(res / 2)
}

// first index i s.t. nums[i] > target
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

	if nums[left] <= target {
		return left + 1
	} else {
		return left
	}
}

// first index i s.t. nums[i] >= target
func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] < target {
		return left + 1
	} else {
		return left
	}
}
