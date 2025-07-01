package leetcode

// Trick: 如果nums[mid] == nums[right]， right指向的元素可以“剥掉”，直到不同的元素出现
func findMin(nums []int) int {
    left, right := 0, len(nums) - 1
    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else if nums[mid] < nums[right] {
            right = mid
        } else { // nums[mid] == nums[right]
            right--
        }
    }

    return nums[left]
}