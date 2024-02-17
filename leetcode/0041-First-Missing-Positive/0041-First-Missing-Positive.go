package leetcode

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		if (1 <= nums[i] && nums[i] <= n) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i-- // // do NOT move forward yet, re-examine the new element swapped to cur position
		}

		
	}

	// 检查是否'众神归位'
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i+1
		}
	}

	return n+1
}