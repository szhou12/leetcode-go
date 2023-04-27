package leetcode

import "sort"

func eventualSafeNodes(graph [][]int) []int {
	// Step 1: Reconstruct adj-list repr + Calculate out-degree
	n := len(graph)
	degree := make([]int, n)
	prev := make([][]int, n)
	for i := 0; i < n; i++ {
		prev[i] = make([]int, 0)
	}
	for from, tos := range graph {
		for _, to := range tos {
			prev[to] = append(prev[to], from)
			degree[from]++
		}
	}

	// Step 2: Topological Sort
	res := make([]int, 0)
	queue := make([]int, 0)
	// Start nodes: out-degree == 0
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
		// cur is a 'deadend', add to res
		res = append(res, cur)

		// Make the next move
		for _, nei := range prev[cur] {
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Step 3: Sort in ascending order
	sort.Ints(res)

	return res
}
