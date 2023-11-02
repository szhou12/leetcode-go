package leetcode

func countVisitedNodes(edges []int) []int {
    // Part 1: Topological Sort
	// Step 1: Reconstruct adj-list repr next[] + calculate in-degree
	// edges[] is already next[]
	n := len(edges)
	indegree := make([]int, n)
	for i := 0; i < n; i++ {
		// a -> b: a = i, b = edges[i]
		b := edges[i]
		indegree[b]++
	}

	// Step 2: Topo Sort to remove nodes not on cycle
	// Start nodes: indegree = 0
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]
		// update
		indegree[edges[cur]]--

		// Make the next move
		if indegree[edges[cur]] == 0 {
			queue = append(queue, edges[cur])
		}
	}

	// res[i] = path length from i to an already visited node
	res := make([]int, n)

	// Part 2: Calculate Cycle Length
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			continue
		}
		if res[i] != 0 {
			continue
		}

		// calculate cycle length
		cycleLen := 1
		j := i
		for edges[j] != i {
			j = edges[j]
			cycleLen++
		}

		// assign cycle length
		j = i
		res[j] = cycleLen
		for edges[j] != i {
			j = edges[j]
			res[j] = cycleLen
		}
	}


	// Part 3: DFS to calculate path length for nodes not on cycle
	for i := 0; i < n; i++ {
		if indegree[i] != 0 {
			continue
		}
		dfs(i, edges, res)
	}

	return res
}

func dfs(cur int, edges []int, res []int) int {
	// base case
	if res[cur] != 0 {
		return res[cur]
	}

	res[cur] = dfs(edges[cur], edges, res) + 1
	return res[cur]
}