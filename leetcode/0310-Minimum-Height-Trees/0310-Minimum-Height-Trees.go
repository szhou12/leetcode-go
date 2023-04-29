package leetcode

func findMinHeightTrees(n int, edges [][]int) []int {
	// Edge Case: onlye one node
	if n == 1 {
		return []int{0}
	}

	// Step 1: Reconstruct adj-list + Calculate degree
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
	popCount := 0 // # of nodes popped from queue
	queue := make([]int, 0)
	// Start nodes: degree == 1
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}
	// Loop
	for len(queue) > 0 {
		// Chech if only 1 or 2 center nodes left
		if n-popCount == 1 || n-popCount == 2 {
			break
		}

		// Per level
		size := len(queue)
		for size > 0 {
			// Cur
			cur := queue[0]
			queue = queue[1:]
			size--
			popCount++

			// Make the next move
			for _, nei := range next[cur] {
				degree[nei]--
				if degree[nei] == 1 {
					queue = append(queue, nei)
				}
			}
		}
	}

	// Step 3: leftover 1 or 2 center nodes are what we want
	res := make([]int, 0)
	res = append(res, queue...)
	return res
}
