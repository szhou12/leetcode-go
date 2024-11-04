package leetcode

import (
	"math/rand"
)

// Solution: Quick Select
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, left int, right int, k int) int {
	// get pivot
	pIndex := getPivotIndex(left, right)
	pivot := nums[pIndex]

	// Ranbow Sort / Sort Colors
	i := left  // all elements BEFORE i-th < pivot
	t := left  // 挡板 that points to first unknown; [i, t) are elements == pivot
	j := right // all elements AFTER j-th > pivot

	for t <= j { // Stops at t == j + 1
		if nums[t] < pivot {
			// swap i, t
			nums[t], nums[i] = nums[i], nums[t]
			// increment both i, t
			i++
			t++
		} else if nums[t] > pivot {
			// swap j, t
			nums[t], nums[j] = nums[j], nums[t]
			// decrement ONLY j bc t receives another unknown
			j--
		} else { // nums[t] == pivot
			t++
		}
	}

	// After while loop, we will have:
	// S S S O O O O L L L L L L
	// l     i     j t         r

	// Recursion
	if right-j >= k {
		return quickSelect(nums, j+1, right, k)
	} else if right-i+1 >= k {
		return pivot
	} else {
		return quickSelect(nums, left, i-1, k-(right-i+1))
	}

}

func getPivotIndex(left int, right int) int {
	return left + rand.Intn(right-left+1)
}
