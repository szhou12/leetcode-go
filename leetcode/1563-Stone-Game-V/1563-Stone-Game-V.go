package leetcode

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	stoneValue = append([]int{0}, stoneValue...)
	presum := make([]int, len(stoneValue))
	for i := 1; i <= n; i++ {
		presum[i] = presum[i-1] + stoneValue[i]
	}
	memo := make([][]int, len(stoneValue))
	for i := 0; i < len(stoneValue); i++ {
		memo[i] = make([]int, len(stoneValue))
	}

	return dfs(1, n, presum, memo, stoneValue)
}

// dfs(i, j) := max stones that cur player (Alice) can get from stoneValue[i:j+1]
func dfs(left int, right int, presum []int, memo [][]int, stoneValue []int) int {
	// base case:
	// when there are only 2 stones left, Alice can only get the smaller one
	if right == left+1 {
		return min(stoneValue[left], stoneValue[right])
	}
	if left >= right {
		return 0
	}

	// memoization
	if memo[left][right] != 0 {
		return memo[left][right]
	}

	// recursion
	for k := left; k < right; k++ {
		// left: [left, k] right: [k+1, right]
		leftSum := presum[k] - presum[left-1]
		rightSum := presum[right] - presum[k]

		if leftSum < rightSum {
			memo[left][right] = max(memo[left][right], leftSum+dfs(left, k, presum, memo, stoneValue))
		} else if leftSum > rightSum {
			memo[left][right] = max(memo[left][right], rightSum+dfs(k+1, right, presum, memo, stoneValue))
		} else { // leftSum == rightSum
			memo[left][right] = max(memo[left][right],
				leftSum+max(dfs(left, k, presum, memo, stoneValue), dfs(k+1, right, presum, memo, stoneValue)))
		}
	}

	return memo[left][right]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
