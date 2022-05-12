package leetcode

func permute(nums []int) [][]int {
	var result [][]int

	dfs(nums, 0, &result)

	return result

}

func dfs(nums []int, index int, result *[][]int) {
	// base case
	if index == len(nums) {
		res := make([]int, len(nums))
		copy(res, nums)
		*result = append(*result, res)
		return
	}

	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, index+1, result)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
