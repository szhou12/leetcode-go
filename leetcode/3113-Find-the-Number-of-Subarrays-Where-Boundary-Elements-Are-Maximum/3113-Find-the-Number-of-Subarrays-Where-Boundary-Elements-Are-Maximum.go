package leetcode

import "sort"

func numberOfSubarrays(nums []int) int64 {
	n := len(nums)

	// Step 1: prevGreaterOrEqual - for a given i as the last element, find its previous greater or equal element
	stack := make([]int, 0)

	prevGreaterOrEqual := make([]int, n)
	for i := 0; i < n; i++ {
		prevGreaterOrEqual[i] = -1
	}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			prevGreaterOrEqual[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	res := 0
	// Step 2: binary search + hashmap
	// mp := {nums[i] : [indices of same values]}
	mp := make(map[int][]int)
	for i := 0; i < n; i++ {
		if _, ok := mp[nums[i]]; !ok {
			mp[nums[i]] = make([]int, 0)
		}
		mp[nums[i]] = append(mp[nums[i]], i)

		// lower bound on target = prevGreaterOrEqual[i]
		lb := lowerBound(mp[nums[i]], prevGreaterOrEqual[i])

		res += i - lb

	}

	return int64(res)

}

func lowerBound(nums []int, target int) int {
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	return nums[i]
}
