package leetcode

import "sort"

func findPairs(nums []int, k int) int {
	// edge case
	if len(nums) <= 1 {
		return 0
	}

	// Step 1: Sort in increasing order!!!
	sort.Ints(nums)

	res := 0
	for left := 0; left < len(nums); left++ {
		// skip same element
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}

		// reset right
		right := left + 1
		// grow right
		for right < len(nums) && nums[right]-nums[left] < k {
			right++
		}

		if right < len(nums) && nums[right]-nums[left] == k {
			res++
		}

	}

	return res
}
