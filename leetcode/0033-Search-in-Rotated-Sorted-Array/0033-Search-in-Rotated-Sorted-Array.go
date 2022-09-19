package leetcode

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// 判别 nums[mid], target分别落在左右哪个区间
		if (nums[mid] >= nums[0] && target >= nums[0]) || (nums[mid] < nums[0] && target < nums[0]) { // nums[mid], target在同一区间
			if nums[mid] > target {
				right = mid // 写成 right = mid - 1 也对
			} else {
				left = mid + 1
			}
		} else if nums[mid] >= nums[0] && target < nums[0] { // nums[mid]在 左区间，target在 右区间
			left = mid + 1
		} else if nums[mid] < nums[0] && target >= nums[0] { // nums[mid]在 右区间，target在 左区间
			right = mid // 写成 right = mid - 1 也对
		}
	}

	// post-processing
	if nums[left] == target {
		return left
	}
	return -1
}
