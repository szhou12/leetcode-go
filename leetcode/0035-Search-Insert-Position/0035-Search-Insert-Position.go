package leetcode

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := right - (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid - 1
		}
	}

	if nums[left] >= target {
		return left
	}

	return left + 1

}
