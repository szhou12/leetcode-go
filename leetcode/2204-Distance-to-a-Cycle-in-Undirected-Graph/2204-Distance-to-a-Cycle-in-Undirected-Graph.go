package leetcode

// My Solution: DP思想
// preset[i] = min distance of all nodes arriving node i
func distanceToCycle(n int, edges [][]int) []int {
	// Step 1: Reconstruct adj-list repr + Calculate degree
	degree := make([]int, n)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)
		degree[a]++
		degree[b]++
	}

	// Step 2: Topological Sort
	preset := make([]map[int]int, n) // preset[i] = min distance all nodes arriving at node i
	for i := 0; i < n; i++ {
		preset[i] = make(map[int]int)
	}
	queue := make([]int, 0)
	// Start nodes: degree == a
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
			preset[i][i] = 0
		}
	}
	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]

		// Make the next move
		for _, nei := range next[cur] {
			// add nei itself to preset[nei]
			if _, ok := preset[nei][nei]; !ok {
				preset[nei][nei] = 0
			}
			// add preset[cur] to preset[nei]
			for node, dist := range preset[cur] {
				if _, ok := preset[nei][node]; !ok { // 没见到过的node = 最先到 = 最短距离
					preset[nei][node] = dist + 1
				}
			}

			degree[nei]--
			if degree[nei] == 1 {
				queue = append(queue, nei)
			}

		}
	}

	// Step 3: form res from preset[i] where i = cycle nodes
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if degree[i] > 1 { // cycle nodes
			res[i] = 0
			for node, dist := range preset[i] {
				res[node] = dist
			}
		}
	}

	return res

}

// 参考答案: Topological Sort + BFS
func distanceToCycle_reference(n int, edges [][]int) []int {
	// Step 1: Reconstruct adj-list repr + Calculate degree
	degree := make([]int, n)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)
		degree[a]++
		degree[b]++
	}

	// Step 2: Topological Sort to find the cycle nodes
	queue := make([]int, 0)
	// Start nodes: degree == 1
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}
	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]

		// Make the next move
		for _, nei := range next[cur] {
			degree[nei]--
			if degree[nei] == 1 {
				queue = append(queue, nei)
			}
		}
	}

	// Step 3: BFS count min distance, starting from cycle nodes
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if degree[i] > 1 {
			res[i] = 0
		} else {
			res[i] = -1
		}
	}
	queue = make([]int, 0)
	// Start nodes: cycle nodes
	for i := 0; i < n; i++ {
		if res[i] == 0 {
			queue = append(queue, i)
		}
	}
	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]

		// Make the next move
		for _, nei := range next[cur] {
			// Check if nei already visited (avoid other cycle nodes) and counted
			if res[nei] != -1 {
				continue
			}
			res[nei] = res[cur] + 1
			queue = append(queue, nei)
		}
	}

	return res
}
