package amz

import "math"

/*
Given an array nums[].
Find the first index where its prefix sum array becomes <= 0. Return -1 if no such index exists.

Example:
nums = [1, 2]
prefix-sum = [1, 3]
Output: -1

nums = [2, -4, 1]
prefix-sum = [2, -2, -1]
Output: 2 (bc -2 < -1, -1 is more close to 0)

nums = [1,2,3,-6]
prefix-sum = [1,3,6,0]
Output: 3
*/

func getFirstNonPosPresum(nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n)


	res := -1
	temp := math.MinInt
	prefixSum[0] = nums[0]
	if prefixSum[0] <= 0 {
		res = 0
		temp = prefixSum[0]
	}
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
		if prefixSum[i] <= 0 && temp < prefixSum[i] {
			res = i
			temp = prefixSum[i]
		}
	}

	return res

}
