package leetcode

import "math"

func minimumOperations(grid [][]int) int {
	n := len(grid[0])
	// dp[i][p] := min operations to fill up first i cols with i-th col being set up with number p
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 10)
		for p := 0; p <= 9; p++ {
			dp[i][p] = math.MaxInt / 2
		}
	}

	// base case: dp[0][p] = min ops of first 0 col = cost to set 0-th col to p
	for p := 0; p <= 9; p++ {
		dp[0][p] = calcCost(grid, 0, p)
	}

	// recurrence: dp[i][p] = min(dp[i-1][q] + cost(p) for q=0..9) for p=0..9 and p!=q
	for i := 1; i < n; i++ {
		for p := 0; p <= 9; p++ { // 这一层循环可优化：只循环三个数: 最高频，次高频，次次高频. 足够保证一定能挑到一个数与左右两列的数都不同，并且操作数是最少的
			// calculate cost to set i-th col to p
			cost := calcCost(grid, i, p)
			for q := 0; q <= 9; q++ {
				if p == q {
					continue
				}
				dp[i][p] = min(dp[i][p], dp[i-1][q]+cost)
			}
		}
	}

	res := math.MaxInt
	for p := 0; p <= 9; p++ {
		res = min(res, dp[n-1][p])
	}
	return res
}

func calcCost(grid [][]int, col int, target int) int {
	m := len(grid)
	cost := 0
	for r := 0; r < m; r++ {
		if grid[r][col] != target {
			cost++
		}
	}
	return cost
}
