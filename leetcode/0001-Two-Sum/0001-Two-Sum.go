package leetcode

import "sort"

// Optimal Solution - map
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int) // key = num, value = num's index

	for idx, val := range nums {
		if matchIdx, ok := dict[target-val]; ok {
			return []int{idx, matchIdx}
		}
		dict[val] = idx
	}
	return nil
}

// Solution - Two Pointers (通法)
// 注意！！！此法只适用于输出是 value 而不是index 的情况，因为一开始的sort就打乱了原来的index
func twoSum_2p(nums []int, target int) []int {
	// sort first
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{nums[left], nums[right]}
		}
	}

	return nil
}
