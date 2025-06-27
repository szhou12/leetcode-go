package leetcode

func countComponents(n int, edges [][]int) int {
	// step 1: convert edges to adj-list representation
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)
	}

	// DFS to find connected components
	visited := make([]bool, n)
	components := 0
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		// each dfs call will mark all nodes connected to i as visited
		dfs(i, next, visited)
		components++
	}

	return components
}

func dfs(node int, next [][]int, visited []bool) {
	// base case: no base case needed as func will automatically ends if no next[node]

	// mark node as visited
	visited[node] = true

	// make the next move
	for _, nei := range next[node] {
		// check if already visited
		if visited[nei] {
			continue
		}
		dfs(nei, next, visited)
	}
}