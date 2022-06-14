package leetcode

import "math"

func find132pattern(nums []int) bool {
	// corner case
	if len(nums) < 3 {
		return false
	}

	n := len(nums)
	// for each A[j], find the min element in A[0...j-1]
	leftMin := make([]int, n)
	leftMin[0] = math.MaxInt
	for j := 1; j < n; j++ {
		leftMin[j] = min(leftMin[j-1], nums[j-1])
	}

	// for each A[j], find the largest element < A[j] in A[j+1...n]
	stack := []int{}
	Ak := math.MinInt
	for j := n - 1; j >= 0; j-- {
		// monotonic stack: increasing order from top to bottom of stack
		for len(stack) != 0 && nums[j] > stack[len(stack)-1] {
			Ak = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		// check: A[i] < A[k] < A[j]
		if leftMin[j] < Ak {
			return true
		}

		stack = append(stack, nums[j])
	}
	return false
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
