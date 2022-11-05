package leetcode

func pivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	leftSum := 0
	rightSum := sum(nums)

	for i := 0; i < len(nums); i++ {
		rightSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}

func sum(nums []int) int {
	res := 0
	for _, val := range nums {
		res += val
	}
	return res
}
