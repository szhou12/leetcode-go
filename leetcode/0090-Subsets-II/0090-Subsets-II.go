package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	var subset []int

	// Sort first so that same elements groupted together
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	dfs(nums, 0, subset, &result)
	return result
}

func dfs(nums []int, index int, subset []int, result *[][]int) {
	// base case
	if index == len(nums) {
		subsetCopy := make([]int, len(subset))
		copy(subsetCopy, subset)
		*result = append(*result, subsetCopy)
		return
	}

	// branch 1: add current element
	subset = append(subset, nums[index])
	dfs(nums, index+1, subset, result)
	subset = subset[:len(subset)-1]

	// KEY: Skip all duplicated elements
	for index < len(nums)-1 && nums[index+1] == nums[index] {
		index++
	}

	// branch 2: not add current element
	dfs(nums, index+1, subset, result)
}
