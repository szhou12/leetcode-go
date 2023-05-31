package leetcode

func longestNiceSubarray(nums []int) int {
	count := 0
	right := 0
	res := 0
	for left := 0; left < len(nums); left++ {
		// 吃
		for right < len(nums) && (count&nums[right] == 0) {
			count += nums[right]
			right++
		}

		// update result
		// case 1: right出界 但是 右边界可以延伸 => [left, right-1]
		// case 2: right没出界 但是 右边界不能延伸 => [left, right-1]
		// case 3: right出界 且 右边界不能延伸 => [left, right-1]
		res = max(res, (right-1)-left+1)

		// 吐
		count -= nums[left]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
