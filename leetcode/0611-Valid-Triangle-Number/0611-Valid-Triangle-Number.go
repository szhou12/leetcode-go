package leetcode

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	result := 0
	for i := 2; i < len(nums); i++ {
		// how many triangles can be formed with the largest edge being nums[i]
		result += allpairs(i, nums)
	}
	return result
}

func allpairs(i int, nums []int) int {
	start := 0
	end := i - 1
	counts := 0
	for start < end {
		if nums[start]+nums[end] > nums[i] {
			// since nums sorted,
			// if nums[start]+nums[end] > nums[i],
			// it means each element > nums[start] can form a triangle with other two edges: nums[end] & nums[i]
			counts += end - start
			end--
		} else {
			start++
		}
	}
	return counts
}
