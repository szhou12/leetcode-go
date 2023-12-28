package leetcode

import "math"

// Easy-To-Understand Solution but will cause TLE
func minimumCost_TLE(source string, target string, original []string, changed []string, cost []int) int64 {
	// Step 1: Assign integer node ID to each string in original & changed
	// deduplication
	strs := make(map[string]bool)
	for i, _ := range original {
		org := original[i]
		chg := changed[i]
		if _, ok := strs[org]; !ok {
			strs[org] = true
		}
		if _, ok := strs[chg]; !ok {
			strs[chg] = true
		}
	}
	// assign ID
	m := len(strs) // total number of nodes
	node := make(map[string]int)
	id := 0
	for str, _ := range strs {
		node[str] = id
		id++
	}

	// Step 2: Floyd
	// dist[i][j] := the minimum cost to travel from node i to node j
	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i != j {
				dist[i][j] = math.MaxInt / 3
			}
		}
	}
	for k, _ := range original {
		from := node[original[k]]
		to := node[changed[k]]
		weight := cost[k]
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				dist[i][j] = min(dist[i][j], dist[i][from]+weight+dist[to][j])
			}
		}
	}

	// Step 3: DP
	n := len(source)
	// dp[i] := the minimum cost to convert source[:i] to target[:i]
	dp := make([]int, n+1)
	source = "#" + source
	target = "#" + target
	// base case
	dp[0] = 0
	// recurrence
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt / 3
		if source[i] == target[i] { // IMPORTANT !!!
			dp[i] = dp[i-1]
		}
		for j := 1; j <= i; j++ {
			// Accessing hashmap leads to TLE
			from, ok1 := node[source[j:i+1]]
			to, ok2 := node[target[j:i+1]]
			if ok1 && ok2 {
				dp[i] = min(dp[i], dp[j-1]+dist[from][to])
			}
		}
	}

	if dp[n] >= math.MaxInt/3 {
		return -1
	}
	return int64(dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
