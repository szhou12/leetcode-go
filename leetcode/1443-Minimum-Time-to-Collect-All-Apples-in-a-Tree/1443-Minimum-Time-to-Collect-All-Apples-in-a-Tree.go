package leetcode

func minTime(n int, edges [][]int, hasApple []bool) int {
	// Step 1: construct adj-list for this tree (undirected graph)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)
	}

	// Step 2: DFS
	var dfs func(cur, parent int) int
	dfs = func(cur, parent int) int {
		time := 0

		for _, nei := range next[cur] {
			if nei == parent {
				continue
			}
			neiTime := dfs(nei, cur)
			// ONLY worth going down to the subtree rooted at nei if
			// 1) nei itself has an apple OR 2) the subtree rooted at nei has at least one apple somewhere (i.e. neiTime > 0)
			if hasApple[nei] || neiTime > 0 {
				time += 2 + neiTime
			}
		}


		return time
	}

	return dfs(0, -1)
}