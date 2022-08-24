package leetcode

func removeElement(nums []int, val int) int {
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] != val {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}

	return left
}
