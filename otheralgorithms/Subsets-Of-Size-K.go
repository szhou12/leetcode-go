package otheralgorithms

/*
 * All Subsets of Size K
 *
 * Given a []int,
 * return a list containing all subsets of size K.
 *
 * Assumptions
 * There are no duplicate elements in the original set.
 *
 */
func subsetsOfSizeK(nums []int, k int) [][]int {
	var result [][]int
	var subset []int

	findSubsetsOfSizeK(nums, 0, subset, &result, k)
	return result
}

// DFS
func findSubsetsOfSizeK(nums []int, index int, subset []int, result *[][]int, k int) {
	// base case
	if k == len(subset) {
		subsetCopy := make([]int, len(subset))
		copy(subsetCopy, subset)
		*result = append(*result, subsetCopy)
		return
	}

	// Early stopping!!!
	// It's possible that index is already out of bound (checked all letters) but still len(subset) < k
	// In this case, we should terminate recursion as long as index == len(subset)
	if index == len(subset) {
		return
	}

	// branch 1: add current element
	subset = append(subset, index+1)
	findSubsetsOfSizeK(nums, index+1, subset, result, k)
	// backtracking
	subset = subset[:len(subset)-1]
	// branch 2: not add current element
	findSubsetsOfSizeK(nums, index+1, subset, result, k)
}
