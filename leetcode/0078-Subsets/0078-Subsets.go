package leetcode

func subsets(nums []int) [][]int {
	var result [][]int
	var subset []int

	dfs(nums, 0, subset, &result)
	return result

}

// 第一种写法： # of branches = 2 + Base Case添加合法子集
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

// 第一种写法： # of branches = 从index到最后 每一个元素 + 沿途添加合法子集
func dfs2(nums []int, index int, subset []int, result *[][]int) {
	// 沿recursion tree添加合法子集
	s := make([]int, len(subset))
	copy(s, subset)
	*result = append(*result, s)

	for i := index; i < len(nums); i++ {
		subset = append(subset, nums[i])
		dfs2(nums, i+1, subset, result)
		subset = subset[:len(subset)-1]
	}

}
