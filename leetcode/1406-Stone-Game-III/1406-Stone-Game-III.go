package leetcode

import "math"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)

	stoneValue = append([]int{0}, stoneValue...)

	presum := make([]int, len(stoneValue))
	for i := 1; i <= n; i++ {
		presum[i] = presum[i-1] + stoneValue[i]
	}

	memo := make([]int, len(stoneValue))
	for i := 0; i < len(stoneValue); i++ {
		memo[i] = math.MinInt
	}

	gain := dfs(1, stoneValue, presum, memo)
	total := presum[n]

	if gain == total-gain {
		return "Tie"
	} else if gain > total-gain {
		return "Alice"
	} else {
		return "Bob"
	}

}

// dfs(start) := max stones that cur player gains if select X=1,2,3 stones from stoneValue[start:]
func dfs(start int, stoneValue []int, presum []int, memo []int) int {
	// base case
	if start >= len(stoneValue) {
		return 0
	}

	// memoization
	if memo[start] > math.MinInt {
		return memo[start]
	}

	res := math.MinInt
	for x := 1; x <= 3; x++ {
		// x = end - start + 1
		end := start + x - 1

		if end >= len(stoneValue) {
			break
		}

		gain := presum[end] - presum[start-1]
		gain += presum[len(stoneValue)-1] - presum[end]
		gain -= dfs(end+1, stoneValue, presum, memo)

		res = max(res, gain)
	}

	memo[start] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
