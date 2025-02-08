package leetcode

// dp[i] := max books can be taken from the subarray ending at i
func maximumBooks(books []int) int64 {
	n := len(books)
	dp := make([]int, n)


	stack := make([]int, 0) // stores index

	// recurrence
	for i := 0; i < n; i++ {
		for len(stack) > 0 && books[stack[len(stack)-1]] > books[i]-(i-stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			j := stack[len(stack)-1]
			L := i - j
			dp[i] = dp[j] + ((books[i]-L+1)+books[i])*L/2
		} else {
			L := min(i+1, books[i])
			dp[i] = ((books[i] - L + 1) + books[i]) * L / 2
		}

		stack = append(stack, i)
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return int64(res)

}

func maximumBooks2(books []int) int64 {
	n := len(books)

	prevSmaller := findPrevSmaller(books)

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		j := prevSmaller[i]
		L := min(books[i], i-j)
		arithmeticSum := (books[i] + (books[i] - L + 1)) * L / 2

		dp[i] = arithmeticSum
		if j != -1 {
			dp[i] += dp[j]
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return int64(res)

}

func findPrevSmaller(books []int) []int {
	n := len(books)

	// preprocessing
	// 0  1  2
	//[5, 3, 7]
	//[5, 2, 5]
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = books[i] - i
	}

	// monotonic increasing stack
	stack := make([]int, 0)

	prevSmaller := make([]int, n)
	for i := 0; i < n; i++ {
		prevSmaller[i] = -1
	}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			prevSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return prevSmaller
}
