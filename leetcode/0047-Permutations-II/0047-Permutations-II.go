package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	// result := dfs1(nums)
	result := dfs2(nums)
	return result
}

func dfs1(nums []int) [][]int {
	used := make([]bool, len(nums))
	cur := []int{}
	result := [][]int{}

	// KEY: MUST sort all elements first in order to use this method!
	sort.Ints(nums)
	dfs1Helper(nums, cur, &used, &result)
	return result
}

func dfs1Helper(nums []int, cur []int, used *[]bool, result *[][]int) {
	// base case
	if len(cur) == len(nums) {
		res := make([]int, len(cur))
		copy(res, cur)
		*result = append(*result, res)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*used)[i] {
			continue
		}
		// KEY: pruning - deduplication
		// MUST keep same elements in order e.g. 2 -> 2' -> 2''
		// in other words, 2' will be added ONLY when 2 has been added
		if i > 0 && nums[i-1] == nums[i] && !(*used)[i-1] {
			continue
		}
		cur = append(cur, nums[i])
		(*used)[i] = true
		dfs1Helper(nums, cur, used, result)
		cur = cur[:len(cur)-1]
		(*used)[i] = false
	}
}

func dfs2(nums []int) [][]int {
	result := [][]int{}

	dfs2Helper(nums, 0, &result)

	return result
}

func dfs2Helper(nums []int, index int, result *[][]int) {
	// base case
	if index == len(nums) {
		res := make([]int, len(nums))
		copy(res, nums)
		*result = append(*result, res)
		return
	}

	// KEY: at each level, if one element already used, then all its following dups won't be used
	used := make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if !used[nums[i]] {
			used[nums[i]] = true
			nums[i], nums[index] = nums[index], nums[i]
			dfs2Helper(nums, index+1, result)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
}
