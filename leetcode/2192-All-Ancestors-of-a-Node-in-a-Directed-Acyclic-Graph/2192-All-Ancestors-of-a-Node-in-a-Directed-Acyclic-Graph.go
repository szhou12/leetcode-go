package leetcode

import "sort"

func getAncestors(n int, edges [][]int) [][]int {
	// Step 1: Reconstruct adj-list repr + Calculate degree
	degree := make([]int, n)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range edges { // O(E)
		from, to := edge[0], edge[1]
		degree[to]++
		next[from] = append(next[from], to)
	}

	// Step 2: Topological Sort O(V+E)
	ancestors := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		ancestors[i] = make(map[int]bool)
	}
	queue := make([]int, 0)
	// Start ndoes: degree == 0
	for i := 0; i < n; i++ {
		if degree[i] == 0 {
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
			// for nei, 记录cur一系的血脉
			for grand, _ := range ancestors[cur] {
				ancestors[nei][grand] = true
			}
			ancestors[nei][cur] = true

			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Step 3: Convert set to slice and Sort
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, 0)
	}
	for i := 0; i < n; i++ { // O(n)
		for ans, _ := range ancestors[i] { // O(n)
			res[i] = append(res[i], ans)
		}
		sort.Ints(res[i]) // O(nlogn)
	}

	return res
}
