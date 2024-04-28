package leetcode

import (
	"math"
	"sort"
)

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	res := math.MaxInt
	for i := 0; i < 3; i++ {
		diff := nums2[0] - nums1[i]
		if valid(nums1, nums2, diff) {
			res = min(res, diff)
		}
	}

	return res
}

func valid(nums1 []int, nums2 []int, diff int) bool {
	errors := 0 // 容错率
	j := 0
	for i := 0; i < len(nums1) && j < len(nums2); i++ {
		if nums1[i]+diff != nums2[j] {
			errors++
		} else {
			j++
		}
	}
	return errors <= 2
}
