package leetcode

import "sort"

func maxSumDistinctTriplet(x []int, y []int) int {
	// Step 1: for each unique value in x, find its largest y value
	// O(n)
	// map x : largest y
	dict := make(map[int]int)
	for i, v := range x {
		if _, ok := dict[v]; !ok {
			dict[v] = y[i]
		} else {
			dict[v] = max(dict[v], y[i])
		}
	}

	if len(dict) < 3 {
		return -1
	}

	// Step 2: sort all largest y values in descending order and adds the top 3
	// O(nlogn)
	nums := make([]int, 0)
	for _, largestY := range dict {
		nums = append(nums, largestY)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	res := 0
	for i := 0; i < 3; i++ {
		res += nums[i]
	}

	return res

}
