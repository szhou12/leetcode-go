package leetcode

import "math"

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	// Step 1: construct adj-list representation of the graph
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)
	}

	// Step 3: Memoization to optimize DFS
	// memo[x][t] := max points that can be collected from subtree rooted at x, with t cut-in-half operations
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, 14) // t=0...13
		for j := 0; j < 14; j++ {
			memo[i][j] = math.MinInt
		}
	}

	// Step 2: DFS
	return dfs(0, -1, next, 0, coins, k, memo)

}

// dfs() := max points that can be collected from subtree rooted at cur
// times := # of cut-in-half operations accumulated from cur's ancestors (NOT including cur's own operation!)
func dfs(cur int, parent int, next [][]int, times int, coins []int, k int, memo [][]int) int {
	if times >= 13 {
		times = 13
	}

	if memo[cur][times] != math.MinInt {
		return memo[cur][times]
	}


	// Option 1: Not cut-in-half cur
	gain1 := reduce(coins[cur], times) - k
	for _, child := range next[cur] {
		if child == parent {
			continue
		}
		gain1 += dfs(child, cur, next, times, coins, k, memo)
	}

	// Option 2: Cut-in-half cur
	gain2 := reduce(coins[cur], times) / 2
	for _, child := range next[cur] {
		if child == parent {
			continue
		}
		gain2 += dfs(child, cur, next, times+1, coins, k, memo)
	}

	memo[cur][times] = max(gain1, gain2)

	return memo[cur][times]
}

// cut the value in half for t times
func reduce(value int, t int) int {
	for i := 0; i < t; i++ {
		value /= 2
	}
	return value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
