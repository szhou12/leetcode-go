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
		mid := left + (right-left)/2 // mid 在两个元素中靠左 否则，下数两行的 right = mid 会在[0, 1]情况下死循环
		if nums[mid] >= target {     // A[mid] 有可能是答案
			right = mid
		} else { // A[0...mid] 都小于 target, 一定没有答案, 移动左边界
			left = mid + 1
		}
	}

	// 查看边界条件：1. target过于小，在左边界以外，2. target在数组范围内，只是没有
	// if left < 0 || nums[left] != target {
	// 	return -1
	// }
	// return left

	// 简化写法：
	// 因为mid的计算方式，以及，left和right的收敛方式
	// 所以, left 不会出界, 选择left为返回对象
	if nums[left] != target { // 两种情况都包括了 1. target过于小，在左边界以外，2. target在数组范围内，只是没有
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
		mid := right - (right-left)/2 // mid 在两个元素中靠右 否则, 下数两行的 left = mid 会在[0, 1]情况下死循环
		if nums[mid] <= target {      // A[mid] 有可能是答案
			left = mid
		} else { // A[mid...n-1] 都大于 target, 一定没有答案, 移动右边界
			right = mid - 1
		}
	}

	// 查看边界条件：1. target过于大，在右边界以外，2. target在数组范围内，只是没有
	// if left > len(nums)-1 || nums[left] != target {
	// 	return -1
	// }
	// return left

	// 简化写法：
	// 因为mid的计算方式，以及，left和right的收敛方式
	// 所以, right 不会出界, 选择right为返回对象
	if nums[right] != target { // 两种情况都包括了 1. target过于小，在左边界以外，2. target在数组范围内，只是没有
		return -1
	}
	return right

}
