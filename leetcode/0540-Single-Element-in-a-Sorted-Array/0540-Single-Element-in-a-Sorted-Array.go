package leetcode

func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums) - 1

	for left < right {
		mid := left + (right - left) / 2

		// ask yourself: why (mid+1) and (mid-1) will not cause out-of-bound error in this problem?
		// hint: nums[] will always be odd-length
		if (mid % 2 == 0 && nums[mid] == nums[mid+1]) || (mid % 2 == 1 && nums[mid] == nums[mid-1]) {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}