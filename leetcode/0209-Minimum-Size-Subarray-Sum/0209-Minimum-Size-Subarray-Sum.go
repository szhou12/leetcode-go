package leetcode

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	res := len(nums) + 1

	for right, val := range nums {
		sum += val
		for sum >= target {
			res = min(res, right-left+1)

			left++
			sum -= nums[left]
		}
	}

	// no such subarray exists
	if res == len(nums)+1 {
		return 0
	}
	return res

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
