package leetcode

import "math"

func numberOfSets(n int, maxDistance int, roads [][]int) int {
	res := 0 // number of possible sets of closing branches

	// brute force all possible combinations of closing branches
	for state := 0; state < (1 << n); state++ { // O(2^n)
		// Step 1: init dp[][] = shortest distance between node i and j
		dp := make([][]int, n)
		for i := 0; i < n; i++ {
			dp[i] = make([]int, n)
			for j := 0; j < n; j++ {
				dp[i][j] = math.MaxInt / 3
			}
		}
		// if node is not closed yet, then the initial shortest distance to itself is 0
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				dp[i][i] = 0
			}
		}

		// Step 2: Floyd–Warshall Algorithm to calculate shortest distance between any two node i and j
		for _, road := range roads {
			a, b, weight := road[0], road[1], road[2]
			// skip if node a is closed
			if (state>>a)&1 == 0 {
				continue
			}
			// skip if node b is closed
			if (state>>b)&1 == 0 {
				continue
			}
			for i := 0; i < n; i++ {
				// skip if node i is closed
				if (state>>i)&1 == 0 {
					continue
				}
				// skip if node j is closed
				for j := 0; j < n; j++ {
					if (state>>j)&1 == 0 {
						continue
					}

					// update dp[i][j]
					dp[i][j] = min(dp[i][j], dp[i][a]+weight+dp[b][j])
					dp[i][j] = min(dp[i][j], dp[i][b]+weight+dp[a][j])
				}

			}
		}

		// Step 3: check if all shorest distance between any two node i and j <= maxDistance
		allReachable := true
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if (state>>j)&1 == 0 {
					continue
				}
				if dp[i][j] > maxDistance {
					allReachable = false
					break
				}
			}
			if !allReachable {
				break
			}

		}
		if allReachable {
			res++
		}
	}

	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}