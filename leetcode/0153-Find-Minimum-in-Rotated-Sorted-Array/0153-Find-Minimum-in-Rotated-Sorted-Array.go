package leetcode

// 性质：rotate以后，一定能分成两个区间，左边区间所有值 > 右边区间所有值
// 每次mid 可以和right比较，如果 nums[mid] > nums[right]， 说明mid在左边区间，right在右边区间
// 如果 nums[mid] < nums[right]， 说明mid和right同在右边区间，mid可能是答案
func findMin(nums []int) int {
    left, right := 0, len(nums) - 1

    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else { // nums[mid] < nums[right]
            right = mid
        }
    }

    return nums[left]
}