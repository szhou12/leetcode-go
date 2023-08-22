package leetcode

func magnificentSets(n int, edges [][]int) int {
	// Step 1: reconstruct adj-list representation
	next := make([][]int, n+1)
	for i := 0; i < len(next); i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)
	}

	cc := make(map[int]int) // {root: max depth}
	// Step 2: BFS on each connected component
	for start := 1; start <= n; start++ {
		queue := make([][]int, 0) // [[node, level]]
		maxDepth := 0
		root := start               // representative of currrent connected component
		visited := make([]int, n+1) // node[level]

		// start node
		queue = append(queue, []int{start, 1})
		visited[start] = 1

		// loop
		for len(queue) > 0 {
			// cur
			temp := queue[0]
			queue = queue[1:]
			cur, level := temp[0], temp[1]
			// update
			maxDepth = max(maxDepth, level)
			root = min(root, cur)

			// Make the next move
			for _, nei := range next[cur] {
				if visited[nei] == 0 {
					queue = append(queue, []int{nei, level + 1})
					visited[nei] = level + 1
				} else if visited[nei] == level { // same-level nodes CANNOT be connected
					return -1
				}
			}
		}

		cc[root] = max(cc[root], maxDepth)
	}

	// Accumulate all cc's max depth
	res := 0
	for _, depth := range cc {
		res += depth
	}
	return res
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
