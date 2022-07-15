package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var subset []int

	// O(nlogn)
	sort.Ints(candidates) // KEY: deduplication

	// O(n + n(n-1) + ... + n!) = O(n!)
	dfs(candidates, target, 0, subset, &result)
	return result
}

func dfs(candidates []int, targetLeft int, index int, subset []int, result *[][]int) {
	// base case
	if targetLeft == 0 {
		res := make([]int, len(subset))
		copy(res, subset)
		*result = append(*result, res)
		return
	}

	// recurrence
	// i = branch i
	for i := index; i < len(candidates); i++ {
		// pruning
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		// proper candidate
		if candidates[i] <= targetLeft {
			subset = append(subset, candidates[i])
			dfs(candidates, targetLeft-candidates[i], i+1, subset, result) // NOTE: i + 1 instead of index + 1!!!
			subset = subset[:len(subset)-1]
		}
	}
}

// Recurrence tree
// candidates = [1,2,2,2,5]
//      {}
//   / | | | \
//  1  2 2 2  5
