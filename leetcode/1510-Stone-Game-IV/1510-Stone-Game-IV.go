package leetcode

func winnerSquareGame(n int) bool {
	memo := make([]int, n+1)

	return dfs(n, memo)
}

// dfs(x) := if cur player can win when there are x stones left
// base case: x == 0 -> false
// recursion:
//  for i := 1; i*i <= x; i++ { if !dfs(x-i*i) -> return true } return false 
func dfs(x int, memo []int) bool {
	// base case
	if x == 0 {
		return false
	}

	// memo
	if memo[x] == 2 {
		return true
	}
	if memo[x] == 1 {
		return false
	}

	// recursion
	for i := 1; i*i <= x; i++ {
		if !dfs(x-i*i, memo) {
			memo[x] = 2
			return true
		}
	}

	memo[x] = 1
	return false
}