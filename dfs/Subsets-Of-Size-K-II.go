package dfs

import "sort"

/*
 * All Subsets of Size K II
 *
 * Given a []int,
 * return a list containing all subsets of size K.
 *
 * Assumptions
 * There ARE duplicate elements in the original set.
 *
 */
func subsetsOfSizeK2(nums []int, k int) [][]int {
	var result [][]int
	var subset []int

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	findSubsetsOfSizeK2(nums, 0, subset, &result, k)
	return result
}

// DFS
func findSubsetsOfSizeK2(nums []int, index int, subset []int, result *[][]int, k int) {
	// base case
	if len(subset) == k {
		subsetCopy := make([]int, len(subset))
		copy(subsetCopy, subset)
		*result = append(*result, subsetCopy)
		return
	}

	// early stopping
	if index == len(nums) {
		return
	}

	// branch 1: add current element
	subset = append(subset, nums[index])
	findSubsetsOfSizeK2(nums, index+1, subset, result, k)
	// backtracking
	subset = subset[:len(subset)-1]
	// skip all duplicates
	for index < len(nums)-1 && nums[index+1] == nums[index] {
		index++
	}
	findSubsetsOfSizeK2(nums, index+1, subset, result, k)
}
