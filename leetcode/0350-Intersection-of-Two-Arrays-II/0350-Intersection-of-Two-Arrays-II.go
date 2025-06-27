package leetcode

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	return solution2(nums1, nums2)
}

// Hash Map: "collision"
func solution1(nums1, nums2 []int) []int {
	res := make([]int, 0)
	freq := make(map[int]int)
	for _, num := range nums1 {
		freq[num]++
	}
	for _, num := range nums2 {
		if _, ok := freq[num]; ok {
			if freq[num] > 0 {
				res = append(res, num)
				freq[num]--
			}

		}
	}

	return res
}

// sorting + two pointers
func solution2(nums1, nums2 []int) []int {
	res := make([]int, 0)

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}
