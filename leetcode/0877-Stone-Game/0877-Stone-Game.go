package leetcode

// dfs(start, end) := max # stones that cur player can get by selecting from piles[start:end+1]
func stoneGame(piles []int) bool {
	n := len(piles)

	piles = append([]int{0}, piles...)
	presum := make([]int, len(piles))
	for i := 1; i <= n; i++ {
		presum[i] = presum[i-1] + piles[i]
	}
	memo := make([][]int, len(piles)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(piles)+1)
	}

	gain := dfs(1, n, piles, presum, memo)
	return gain > presum[n]-gain
}

func dfs(start int, end int, piles []int, presum []int, memo [][]int) int {
	// base case
	if start == end {
		return piles[start]
	}

	// memoization
	if memo[start][end] != 0 {
		return memo[start][end]
	}

	memo[start][end] = max((presum[end]-presum[start-1])-dfs(start+1, end, piles, presum, memo),
		(presum[end]-presum[start-1])-dfs(start, end-1, piles, presum, memo))

	return memo[start][end]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
