package leetcode

func collectTheCoins(coins []int, edges [][]int) int {
	// Step 1: pre-processing
	n := len(coins)
	degree := make([]int, n)
	next := make([]map[int]bool, n) // find all its neighbors for each node
	for i := 0; i < n; i++ {
		next[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		degree[a]++
		degree[b]++
		next[a][b] = true
		next[b][a] = true
	}

	// Step 2: Prune redundant leaf nodes that have no coins
	deleted := prune(&degree, &next, &coins, n)

	// Step 3: Mark depth
	depth := mark(&degree, &next, &deleted, n)

	// Step 4: Calculate result
	m := 0
	for i := 0; i < n; i++ {
		if depth[i] >= 3 {
			m++
		}
	}
	if m-1 >= 0 {
		return 2 * (m - 1)
	}
	return 0
}

func mark(degree *[]int, next *[]map[int]bool, deleted *[]int, n int) []int {
	depth := make([]int, n)
	for i := 0; i < n; i++ {
		depth[i] = -1
	}
	queue := make([]int, 0)

	// Start nodes
	for i := 0; i < n; i++ {
		if (*degree)[i] == 1 && (*deleted)[i] != 1 {
			depth[i] = 1
			queue = append(queue, i)
		}
	}

	// Loop
	for len(queue) > 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			// Cur
			cur := queue[0]
			queue = queue[1:]

			// Push neighbors
			for nei, _ := range (*next)[cur] {
				(*degree)[nei]--
				delete((*next)[nei], cur)
				depth[nei] = max(depth[nei], depth[cur]+1)
				if (*degree)[nei] == 1 && (*deleted)[nei] != 1 {
					queue = append(queue, nei)
				}
			}
		}
	}

	return depth
}

func prune(degree *[]int, next *[]map[int]bool, coins *[]int, n int) []int {
	deleted := make([]int, n)
	queue := make([]int, 0)

	// Start nodes
	for i := 0; i < n; i++ {
		if (*degree)[i] == 1 && (*coins)[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Loop
	for len(queue) > 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			// Cur
			cur := queue[0]
			queue = queue[1:]
			deleted[cur] = 1

			// Push neighbors
			for nei, _ := range (*next)[cur] {
				(*degree)[nei]--
				delete((*next)[nei], cur)
				if (*degree)[nei] == 1 && (*coins)[nei] == 0 {
					queue = append(queue, nei)
				}
			}
		}
	}

	return deleted
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
