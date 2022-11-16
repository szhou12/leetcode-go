package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// Step 1: sort in increasing order
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	diff := math.MaxInt
	res := 0
	// Step 2: 对每个元素i，找符合的two-sum
	for i := 0; i < len(nums); i++ {
		// skip duplicated nums[i]
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			threeSum := nums[i] + nums[left] + nums[right]
			if abs(threeSum-target) < diff { // find a closer three-sum
				diff = abs(threeSum - target)
				res = threeSum
			}
			if threeSum == target {
				return threeSum
			} else if threeSum < target {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
