package leetcode

func predictTheWinner(nums []int) bool {
	n := len(nums)
	// edge case: only one number, 1st player wins
	if n == 1 {
		return true
	}

	nums = append([]int{0}, nums...)
	presum := make([]int, len(nums))
	for i := 1; i <= n; i++ {
		presum[i] = presum[i-1] + nums[i]
	}
	memo := make([][]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(nums))
	}

	gain := dfs(1, n, nums, presum, memo)
	return gain >= presum[n]-gain
}

func dfs(start int, end int, nums []int, presum []int, memo [][]int) int {
	// base case
	if start == end {
		return nums[start]
	}

	// memoization
	if memo[start][end] != 0 {
		return memo[start][end]
	}

	memo[start][end] = max(presum[end]-presum[start-1]-dfs(start+1, end, nums, presum, memo),
		presum[end]-presum[start-1]-dfs(start, end-1, nums, presum, memo))

	return memo[start][end]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
