package leetcode

import "sort"

func numberOfSubarrays(nums []int) int64 {
	n := len(nums)

	// Step 1: prevGreater- for a given i as the last element, find its previous greater element
	stack := make([]int, 0)

	prevGreater := make([]int, n)
	for i := 0; i < n; i++ {
		prevGreater[i] = -1
	}

	// Compute prevGreater array
	// 写法一: right --> left
	for i := n - 1; i >= 0; i-- {
		// 注意：这里不是 <=
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			prevGreater[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	// 写法二: left --> right
	// for i := 0; i < n; i++ {
	// 	// 注意：这里是 <=
	// 	for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
	// 		stack = stack[:len(stack)-1]
	// 	}
	// 	if len(stack) > 0 {
	// 		prevGreater[i] = stack[len(stack)-1]
	// 	}
	// 	stack = append(stack, i)
	// }

	res := 0
	// Step 2: binary search + hashmap
	// mp := {nums[i] : [indices of same values in ascending order]}
	mp := make(map[int][]int)
	for i := 0; i < n; i++ {
		mp[nums[i]] = append(mp[nums[i]], i)

		// upper bound: find first index of nums[i] appears after prevGreater[i]'s index
		head := upperBound(mp[nums[i]], prevGreater[i])

		res += len(mp[nums[i]]) - head

	}

	return int64(res)

}

func upperBound(nums []int, target int) int {
	idx := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})
	return idx
}
