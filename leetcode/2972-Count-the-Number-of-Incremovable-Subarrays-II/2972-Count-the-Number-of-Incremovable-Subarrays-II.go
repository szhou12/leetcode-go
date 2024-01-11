package leetcode

import "math"

func incremovableSubarrayCount(nums []int) int64 {
	n := len(nums)
	// Step 1: Prepend min and append max
	nums = append([]int{math.MinInt}, nums...)
	nums = append(nums, math.MaxInt)

	// Step 2: find longest ascending [1...L] and ascending [R...n]
	L := 1
	for L < n {
		if nums[L] < nums[L+1] { // MUST write this way so that L is included in [1...L]. DO NOT use nums[L-1] < nums[L]
			L++
		} else {
			break
		}
	}

	R := n
	for R >= 1 {
		if nums[R-1] < nums[R] {
			R--
		} else {
			break
		}
	}

	// Step 3: No middle section - whole array is ascending
	// arbitrarily select two end points (no overlap) divided by 2 + two end points at same index
	if R < L {
		return int64(n*(n-1)/2 + n)
	}

	// Step 4: for i in [0 ... L], find smallest j in [R ... n+1] s.t. nums[i] < nums[j]
	// Binary Search upper bound
	res := 0
	for i := 0; i <= L; i++ {
		j := upperBound(nums[R:], nums[i])
		res += (n + 1) - (R + j) + 1 // NOTE: j is 0-index in nums[R:], so need to add R back
	}

	return int64(res)

}

// find first index i s.t. nums[i] > target
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
