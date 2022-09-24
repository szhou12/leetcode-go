package leetcode

func searchRange(nums []int, target int) []int {

	firstOccur := findFirst(nums, target)
	lastOccur := findLast(nums, target)
	return []int{firstOccur, lastOccur}
}

func findFirst(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2 // mid 在两个元素中靠左 b/c line 20
		if nums[mid] == target {     // A[mid] 有可能是答案
			right = mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 查看边界条件：1. target过于小，在左边界以外，2. target在数组范围内，只是没有
	if left < 0 || nums[left] != target {
		return -1
	}
	return left
}

func findLast(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := right - (right-left)/2 // mid 在两个元素中靠右 bc line 44
		if nums[mid] == target {      // A[mid] 有可能是答案
			left = mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 查看边界条件：1. target过于大，在右边界以外，2. target在数组范围内，只是没有
	if left > len(nums)-1 || nums[left] != target {
		return -1
	}
	return left

}
