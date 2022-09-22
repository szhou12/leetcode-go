package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	// Step 1: sort in ascending order
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var res [][]int

	// Step 2: 对每个元素，找符合的two-sum
	// 1,   2,   2,   2,   3,   4
	// i-->
	//    left-->           <--right
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 为防止duplicate triplet的产生，skip all duplicate item
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right { // 对第i元素，找到所有可能与它配对的 two-sum
			twoSum := nums[left] + nums[right]
			if nums[i]+twoSum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++

				for left < right && nums[left] == nums[left-1] { // 为防止duplicate triplet的产生，skip all duplicate item
					left++
				}
			} else if nums[i]+twoSum < 0 { // 现在的3sum小了，而right指向元素是尽可能最大的了，那只能增加left
				left++
			} else { // 现在的3sum大了，而left指向元素是尽可能最小的了，那只能减小right
				right--
			}
		}
	}

	return res
}
