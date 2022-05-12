package leetcode

func subsets(nums []int) [][]int {
	var result [][]int
	var subset []int

	dfs(nums, 0, subset, &result)
	return result

}

func dfs(nums []int, index int, subset []int, result *[][]int) {
	// base case
	if index == len(nums) {
		s := make([]int, len(subset))
		copy(s, subset)
		*result = append(*result, s)
		return
	}

	// add to subset
	subset = append(subset, nums[index])
	dfs(nums, index+1, subset, result)
	subset = subset[:len(subset)-1]

	dfs(nums, index+1, subset, result)
}
