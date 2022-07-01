package leetcode

func sortColors(nums []int) {
	i := 0
	t := 0
	j := len(nums) - 1

	for t <= j {
		if nums[t] == 0 {
			//swap i, t
			nums[i], nums[t] = nums[t], nums[i]
			i++
			t++
		} else if nums[t] == 2 {
			//swap j, t
			nums[j], nums[t] = nums[t], nums[j]
			j--
		} else {
			t++
		}
	}
}
