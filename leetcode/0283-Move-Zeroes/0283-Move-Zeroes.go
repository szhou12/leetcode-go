package leetcode

// My Solution：快慢指针同向而行，快指针向慢指针扔非零元素
// 物理意义:
// left = left 挡板的左边都是非零元素
// [left, right) = 区间都是 0
func moveZeroes(nums []int) {

	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}

	for left < len(nums) {
		nums[left] = 0
		left++
	}

}

// Optimal Solution: 快慢指针同向而行，当快指针非零时，直接交换快慢指针的元素
func moveZeroes_optimal(nums []int) {
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
