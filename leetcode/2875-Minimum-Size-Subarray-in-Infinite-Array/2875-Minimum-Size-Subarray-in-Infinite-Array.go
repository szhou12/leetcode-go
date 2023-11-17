package leetcode

import (
	"math"
)

func minSizeSubarray(nums []int, target int) int {
	n := len(nums)

	// double the nums
	nums2 := make([]int, n)
	copy(nums2, nums)
	nums2 = append(nums2, nums...)

	// nums sum
	sum := 0
	for _, num := range nums {
		sum += num
	}

	k := target / sum
	leftTarget := target % sum

	l := slidingWindow(nums2, leftTarget)

	if l == -1 {
		return -1
	}

	return l + k*n
}

// modified sliding window: find the shortest subarray with sum == target
func slidingWindow(nums []int, target int) int {
	n := len(nums)

	minLen := math.MaxInt
	sum := 0
	right := 0
	for left := 0; left < n; left++ {
		// 吃
		for right < n && sum < target {
			sum += nums[right]
			right++
		}

		// update minLen
		if sum == target {
			minLen = min(minLen, right-left)
		}

		// 吐
		sum -= nums[left]
	}

	if minLen > n {
		return -1
	}
	return minLen
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
