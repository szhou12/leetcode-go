package leetcode

// dp[i] := max books can be taken from the subarray ending at i
func maximumBooks(books []int) int64 {
	n := len(books)
	dp := make([]int, n)

	// base case
	dp[0] = books[0]

	stack := make([]int, 0) // stores index

	// recurrence
	for i := 1; i < n; i++ {
		for len(stack) > 0 && books[stack[len(stack)-1]] > books[i]-(i-stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			j := stack[len(stack)-1]
			L := i - j
			dp[i] = dp[j] + ((books[i]-L+1)+books[i])*L/2
		} else {
			L := min(i + 1, books[i])
			dp[i] = ((books[i]-L+1)+books[i])*L/2
		}

		stack = append(stack, i)
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return int64(res)

}
