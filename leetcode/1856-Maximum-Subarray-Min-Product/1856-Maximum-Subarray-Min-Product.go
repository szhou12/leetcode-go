package leetcode

import "math"

func maxSumMinProduct(nums []int) int {
	n := len(nums)

	prefixSum := make([]int, n)
	prefixSum[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	prevSmaller, nextSmaller := findSmaller(nums)

	res := math.MinInt
	for i := 0; i < n; i++ {
		var l int
		if prevSmaller[i] == -1 {
			l = 0
		} else {
			l = prefixSum[prevSmaller[i]]
		}

		r := prefixSum[nextSmaller[i] - 1]

		cur := nums[i] * (r - l)

		res = max(res, cur)

	}

	return res % int(1e9+7)
}

func findSmaller(nums []int) ([]int, []int) {
	n := len(nums)

	// increasing stack
	stack := make([]int, 0)

	nextSmaller := make([]int, n)
	for i := 0; i < n; i++ {
		nextSmaller[i] = n
	}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			nextSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)

	prevSmaller := make([]int, n)

	for i := 0; i < n; i++ {
		prevSmaller[i] = -1
	}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack) - 1]] > nums[i] {
			prevSmaller[stack[len(stack) - 1]] = i
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	return prevSmaller, nextSmaller
}
