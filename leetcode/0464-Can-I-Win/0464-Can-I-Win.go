package leetcode

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	// corner case: if total sum < desiredTotal, no one can win
	// n*(n+1)/ 2
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}

	// 2^n many states
	// +1 bc we use 1-based index
	// memo[state] == 2: win
	// memo[state] == 1: lose
	memo := make([]int, 1<<(maxChoosableInteger+1))

	return dfs(0, 0, maxChoosableInteger, desiredTotal, memo)

}

// dfs(state, sum) := Whether current player can win given cur state of numbers are selected and cur selected total 'sum'
func dfs(state int, sum int, maxChoosableInteger int, desiredTotal int, memo []int) bool {
	// memoization
	if memo[state] == 2 {
		return true
	}
	if memo[state] == 1 {
		return false
	}

	// base case: all numbers are selected
	// no need to write base case bc recursion loop automatically ends when all selected

	// recursion
	for i := 1; i <= maxChoosableInteger; i++ {
		// check if i is selected already
		if (state>>i)&1 == 1 {
			continue
		}

		// check if selecting i can win
		if sum+i >= desiredTotal {
			memo[state] = 2
			return true
		}

		// if opponent can't win with rest of choices, cur player wins
		if !dfs(state+(1<<i), sum+i, maxChoosableInteger, desiredTotal, memo) {
			memo[state] = 2
			return true
		}
	}

	memo[state] = 1
	return false
}
