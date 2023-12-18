package leetcode

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Step 1: init dp[][]
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				dp[i][j] = 0
			} else {
				dp[i][j] = math.MaxInt / 3 // divided by 3 in case of overflow when dp[i][a]+weight+dp[b][j] (adding 3 parts)
			}
		}
	}

	// Step 2: Floyd
	for _, edge := range edges {
		a, b, weight := edge[0], edge[1], edge[2]
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dp[i][j] = min(dp[i][j], dp[i][a]+weight+dp[b][j])
				dp[i][j] = min(dp[i][j], dp[i][b]+weight+dp[a][j])
			}
		}
	}

	// Step 3: find the city with the smallest number of reachale cities
	reachables := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if dp[i][j] <= distanceThreshold {
				reachables[i]++
			}
		}
	}
	minCities := math.MaxInt
	res := -1
	for i, cities := range reachables {
		if cities <= minCities {
			res = i
			minCities = cities
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
