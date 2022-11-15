package leetcode

import "sort"

// Optimal Solution - Sort + HashMap
func smallerNumbersThanCurrent(nums []int) []int {
	numsOrigin := make([]int, len(nums))
	copy(numsOrigin, nums)
	// Step 1: sort in increasing order
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// Step 2: map (k=elemment, v=index); loop backwards
	dict := make(map[int]int)
	for i := len(nums) - 1; i >= 0; i-- {
		dict[nums[i]] = i
	}

	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = dict[numsOrigin[i]]
	}
	return res
}

// Brute force solution
func smallerNumbersThanCurrent_bruteforce(nums []int) []int {
	counter := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[j] < nums[i] {
				counter[i]++
			}
		}
	}

	return counter
}
