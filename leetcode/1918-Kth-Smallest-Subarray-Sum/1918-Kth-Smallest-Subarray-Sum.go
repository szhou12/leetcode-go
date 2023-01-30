package leetcode

import "math"

func kthSmallestSubarraySum(nums []int, k int) int {
	left, right := math.MaxInt, 0
	for _, num := range nums {
		left = min(left, num)
		right += num
	}

	// 为什么要有额外的 prefixSum array ? 没有行不行？ Ans: 不行
	// 为什么 prefixSum array 长度为 len(nums) + 1? 长度为 len(nums) 行不行? Ans: 不行
	prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	for left < right { // O(logD)
		mid := left + (right-left)/2
		count := countLessOrEqual(mid, prefixSum) // O(n)
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// two pointers
// 物理意义:
// count # of prefix-sum pair diffs (subarray sums) that are <= mid
// if returned # >= k, mid may be the answer; if returned # < k, mid def not answer
func countLessOrEqual(mid int, prefixSum []int) int {
	count := 0

	right := 0
	for left := 0; left < len(prefixSum); left++ {
		for right < len(prefixSum) && prefixSum[right]-prefixSum[left] <= mid {
			// Note: prefixSum[right]-prefixSum[left] = subarray sum for nums[left+1:right+1] ([left+1, right] 左闭右闭)
			right++
		}

		// [left+1, right-1] 左闭右闭 这部分的长度 = # of subarray sums
		count += right - (left + 1)
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
