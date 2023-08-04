package leetcode

func stoneGameII(piles []int) int {
	n := len(piles)
	piles = append([]int{0}, piles...)

	presum := make([]int, len(piles))
	for i := 1; i <= n; i++ {
		presum[i] = presum[i-1] + piles[i]
	}

	memo := make([][]int, len(piles))
	for i := 0; i < len(piles); i++ {
		memo[i] = make([]int, len(piles))
	}

	return dfs(1, 1, piles, presum, memo)

}

// dfs(start, M) := max # stones that cur player can get by first X piles (1 <= X <= 2*M) from piles[start:]
func dfs(start int, M int, piles []int, presum []int, memo [][]int) int {
	// base case
	if start >= len(piles) {
		return 0
	}

	// memoization
	if memo[start][M] != 0 {
		return memo[start][M]
	}

	res := 0
	for x := 1; x <= 2*M; x++ {
		// x = end-start+1
		end := start + x - 1
		
		if end >= len(piles) {
			break
		}

		// stones gained at cur round
		stones := presum[end] - presum[start-1]
		// stones gained from rest of rounds
		stones += presum[len(piles)-1] - presum[end]
		stones -= dfs(end+1, max(x, M), piles, presum, memo)
		// select max
		res = max(res, stones)

	}

	memo[start][M] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
