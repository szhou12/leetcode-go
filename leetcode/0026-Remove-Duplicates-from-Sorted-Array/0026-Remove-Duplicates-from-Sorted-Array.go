package leetcode

func removeDuplicates(nums []int) int {
	l := 0

	for r := 0; r < len(nums); r++ {
		if nums[l] != nums[r] {
			l++
			nums[l] = nums[r]
		}
	}

	return l + 1
}
