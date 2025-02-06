package leetcode

func maxAscendingSum(nums []int) int {
	n := len(nums)

	res := nums[0]

	cur := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			cur += nums[i]
		} else {
			cur = nums[i]
		}

		res = max(res, cur)
	}

	return res
}