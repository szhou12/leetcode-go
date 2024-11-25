package leetcode

import (
	"math"
	"sort"
)

/*
Key Idea:
Sort nums[] in ascending order first.
Outer: Binary Search - given a difference, how many pairs <= difference?

	Inner: Binary Search (upper bound) - count that given nums[i], how many j s.t. nums[j] - nums[i] <= difference
*/
func smallestDistancePair(nums []int, k int) int {
	n := len(nums)

	sort.Ints(nums)

	// largest possible difference is between the first and last element
	right := nums[n-1] - nums[0]
	// smallest possible difference is found among the adjacent
	left := math.MaxInt
	for i := 1; i < n; i++ {
		left = min(left, nums[i]-nums[i-1])
	}

	// Outer Binary Search
	for left < right {
		mid := left + (right-left)/2
		if moreOrEqualK(nums, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left

}

func moreOrEqualK(nums []int, k int, targetDiff int) bool {
	n := len(nums)

	count := 0

	// targetDiff = 3
	// 1  2  3  4  4  5  6
	//    i       j-1 j
	for i := 0; i < n; i++ {
		// nums[j] - nums[i] <= targetDiff
		// nums[j] <= targetDiff + nums[i]
		// find 1st index j s.t. nums[j] > targetDiff + nums[i]
		// so # pais <= targetDiff = (j-1) - i

		// Upper Bound
		j := upperBound(nums, targetDiff+nums[i])
		count += (j - 1) - i

	}

	return count >= k

}

func upperBound(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool { return nums[i] > target })
}
