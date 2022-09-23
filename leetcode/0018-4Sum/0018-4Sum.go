package leetcode

import "sort"

// Solution: one more pointer + three sum
// sort first
// i = 0 -> len(nums) - 3
// left := i+1; right := len(nums) - 1
func fourSum(nums []int, target int) [][]int {
	// Sort first
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		triplets := threeSum(nums[i+1:], target-nums[i])

		for _, v := range triplets {
			v = append([]int{nums[i]}, v...)
			res = append(res, v)
		}
	}

	return res

}

func threeSum(nums []int, target int) [][]int {

	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			twoSum := nums[left] + nums[right]
			if nums[i]+twoSum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++

				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if nums[i]+twoSum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
