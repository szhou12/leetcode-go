package otheralgorithms

func allPermutationsOfSubsets(nums []int) [][]int {
	var result [][]int
	findAllPermutations(nums, 0, &result)
	return result
}

func findAllPermutations(nums []int, index int, result *[][]int) {
	// base case
	if index == len(nums) {
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		*result = append(*result, numsCopy)
		return
	}

	// add subset at each level
	*result = append(*result, nums[0:index])

	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		findAllPermutations(nums, index+1, result)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
