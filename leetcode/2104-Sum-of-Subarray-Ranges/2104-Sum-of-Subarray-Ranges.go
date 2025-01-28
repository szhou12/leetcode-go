package leetcode

func subArrayRanges(nums []int) int64 {
	n := len(nums)

	nextSmaller, prevSmallerOrEqual := findSmaller(nums)
	nextGreater, prevGreaterOrEqual := findGreater(nums)

	res := 0

	// sum of subarray maxs
	for i := 0; i < n; i++ {
		counts := (i - prevGreaterOrEqual[i]) * (nextGreater[i] - i)
		res += nums[i] * counts
	}
	
	// sum of subarray mins
	for i := 0; i < n; i++ {
		counts := (i - prevSmallerOrEqual[i]) * (nextSmaller[i] - i)
		res -= nums[i] * counts
	}

	return int64(res)

}

func findSmaller(nums []int) ([]int, []int) {
	n := len(nums)

	// monotonic increasing stack
	stack := make([]int, 0)

	// nextSmaller[i] = index of i-th element's next strictly smaller element
	nextSmaller := make([]int, n)
	for i := 0; i < n; i++ {
		nextSmaller[i] = n // dummy nextSmaller in case no valid nextSmaller
	}

	// from left to right
	for i := 0; i < n; i++ {
		// 守门员守不住了，标记谁踢掉的守门员，踢掉守门员
		for len(stack) > 0 && nums[stack[len(stack) - 1]] > nums[i] {
			nextSmaller[stack[len(stack) - 1]] = i
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	// clear stack
	stack = make([]int, 0)

	// prevSmallerOrEqual[i] = index of i-th element's previous smaller or equal element
	prevSmallerOrEqual := make([]int, n)
	for i := 0; i < n; i++ {
		prevSmallerOrEqual[i] = -1 // dummy prevSmallerOrEqual in case no valid prevSmallerOrEqual
	}

	// from right to left
	for i := n-1; i >= 0; i-- {
		// 守门员守不住了，标记谁踢掉的守门员，踢掉守门员
		for len(stack) > 0 && nums[stack[len(stack) - 1]] >= nums[i] {
			prevSmallerOrEqual[stack[len(stack) - 1]] = i
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	return nextSmaller, prevSmallerOrEqual

}

func findGreater(nums []int) ([]int, []int) {
	n := len(nums)

	// monotonic decreasing stack
	stack := make([]int, 0)

	// nextGreater[i] = index of i-th element's next strictly greater element
	nextGreater := make([]int, n)
	for i := 0; i < n; i++ {
		nextGreater[i] = n // dummy nextGreater in case no valid nextGreater
	}

	// from left to right
	for i := 0; i < n; i++ {
		// 守门员守不住了，标记谁踢掉的守门员，踢掉守门员
		for len(stack) > 0 && nums[stack[len(stack) - 1]] < nums[i] {
			nextGreater[stack[len(stack) - 1]] = i
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	// clear stack
	stack = make([]int, 0)

	// prevGreaterOrEqual[i] = index of i-th element's previous greater or equal element
	prevGreaterOrEqual := make([]int, n)
	for i := 0; i < n; i++ {
		prevGreaterOrEqual[i] = -1 // dummy prevGreaterOrEqual in case no valid prevGreaterOrEqual
	}

	// from right to left
	for i := n-1; i >= 0; i-- {
		// 守门员守不住了，标记谁踢掉的守门员，踢掉守门员
		for len(stack) > 0 && nums[stack[len(stack) - 1]] <= nums[i] {
			prevGreaterOrEqual[stack[len(stack) - 1]] = i
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	return nextGreater, prevGreaterOrEqual

}