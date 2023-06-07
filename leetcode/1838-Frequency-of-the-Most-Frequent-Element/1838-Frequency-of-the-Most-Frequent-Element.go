package leetcode

import "sort"

func maxFrequency(nums []int, k int) int {
	// Step 1: sort in ascending order so that we can use sliding window
	sort.Ints(nums)

	// Step 2: Sliding window
	// 固定右边界，收缩左边界
	// nums[right] is the target number that we want to add prev nums to the same frequency
	res := 1

	count := 0 // # of operations
	left := 0
	for right := 1; right < len(nums); right++ {
		// [left : right-1] -> [left : right]
		count += (right - left) * (nums[right] - nums[right-1])
		for count > k {
			// 吐掉最左边的增量
			count -= nums[right] - nums[left]
			left++
		}
		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
