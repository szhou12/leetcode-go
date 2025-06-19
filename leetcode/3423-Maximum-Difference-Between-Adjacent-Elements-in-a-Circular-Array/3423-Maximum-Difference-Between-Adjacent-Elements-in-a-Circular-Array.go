package leetcode

func maxAdjacentDistance(nums []int) int {
	// double the array
	nums = append(nums, nums...)

	res := 0
	for i := 0; i+1 < len(nums); i++ {
		res = max(res, abs(nums[i]-nums[i+1]))
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}