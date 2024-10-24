package otheralgorithms

import (
	"math"
	"sort"
)

/*
Cost array of length n where cost[i] = cost of buying ith item.
For each round, there are 3 options to buy:
1. buy leftmost item i at cost[i]
2. but rightmost item j at cost[j]
3. buy both leftmost and rightmost items at jointCost = X. One can only use this option at most k times.
Find the minimum cost to buy all items.

Example: cost = [1,2,3], joinCost = 2, k = 1
Round 1: buy leftmost item at cost[0] = 1
Round 2: buy rest of two items at jointCost = 2
Total cost = 1 + 2 = 3
*/

// dp[l][r][t] = min cost to buy items from cost[l:r] (inclusive) with t (0 <= t <= k) times of jointCost
// Base cases:
// 1. dp[l][r][t] = 0 if l > r
// 2. dp[l][r][t] = cost[l] if l == r
// Recurrence:
//
//	cost of option 1 = cost[l] + dp[l+1][r][t]
//	cost of option 2 = cost[r] + dp[l][r-1][t]
//	cost of option 3 = jointCost + dp[l+1][r-1][t+1] if t < k && l < r
//
// dp[l][r][t] = min(cost of option 1, cost of option 2, cost of option 3)
// Time complexity = O(n^2 * k)
func minCostPurchase(cost []int, jointCost int, k int) int {
	n := len(cost)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k+1)
			for t := 0; t <= k; t++ {
				dp[i][j][t] = math.MaxInt
			}
		}
	}

	var solveDP func(l, r, t int) int
	solveDP = func(l, r, t int) int {
		// base cases
		if l > r {
			return 0
		}
		if l == r {
			return cost[l]
		}

		// memo
		if dp[l][r][t] != math.MaxInt {
			return dp[l][r][t]
		}

		// recurrence
		cost1 := cost[l] + solveDP(l+1, r, t)
		dp[l][r][t] = min(dp[l][r][t], cost1)

		cost2 := cost[r] + solveDP(l, r-1, t)
		dp[l][r][t] = min(dp[l][r][t], cost2)

		if l < r && t < k {
			cost3 := jointCost + solveDP(l+1, r-1, t+1)
			dp[l][r][t] = min(dp[l][r][t], cost3)
		}

		return dp[l][r][t]
	}

	return solveDP(0, n-1, 0)
}

/*
Optimal Solution: Greedy
Idea: 我总有办法通过单独拿走左右两边不是我要用jointCost拿的书，来构造左右两边的书可以同时被拿走的情况
1. Sort the cost array in ascending order
2. Iterate from largest cost to smallest cost: 
	Iteratively, if the sum of the two largest remaining items is greater than jointCost and kRemaining > 0, use jointCost to take them both
*/
func minCostPurchase_optimal(cost []int, jointCost int, k int) int {
	n := len(cost)
	sort.Ints(cost) // Sort in ascending order

	left := 0
	right := n - 1
	kRemaining := k
	totalCost := 0

	for left <= right {
		if left == right {
			// Only one item remains
			totalCost += cost[left]
			left++
		} else {
			// Sum of the two largest remaining items
			sumCost := cost[right] + cost[right-1]
			if sumCost > jointCost && kRemaining > 0 {
				// Use joint cost option
				totalCost += jointCost
				right -= 2
				kRemaining--
			} else {
				// Buy the rightmost item separately
				totalCost += cost[right]
				right--
			}
		}
	}

	return totalCost
}
