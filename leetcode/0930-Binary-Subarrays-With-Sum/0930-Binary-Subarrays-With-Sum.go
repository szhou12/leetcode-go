package leetcode

// 模版Flex
func numSubarraysWithSum(nums []int, goal int) int {
	postZeros := countPostZeros(nums)

	res := 0
	right := 0
	sum := 0
	for left := 0; left < len(nums); left++ {
		// left >= right 防止 left 超过 right
		// 超过的情况: goal=0, 因为 sum 初始值=0, 此时不会进入loop, right 此时不会更新
		for left >= right || (right < len(nums) && sum < goal) {
			sum += nums[right]
			right++
		}

		if sum == goal {
			res += 1 + postZeros[right-1]
		}

		sum -= nums[left]

	}
	return res

}

func countPostZeros(nums []int) []int {
	res := make([]int, len(nums))

	count := 0
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = count
		if nums[i] == 0 {
			count++
		} else {
			count = 0
		}
	}

	return res
}
