package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	result := dfs1(nums)
	return result
}

func dfs1(nums []int) [][]int {
	used := make([]bool, len(nums))
	cur := []int{}
	result := [][]int{}

	// KEY: have to sort all elements first in order to use method!
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
