package leetcode

import "math"

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	// Step 1: Reconstruct adj-list repr
	next := make([][]int, n)
	for i, from := range bombs {
		for j, to := range bombs {
			if i == j {
				continue
			}
			if reachable(from, to) {
				next[i] = append(next[i], j)
			}
		}
	}

	// Step 2: BFS on each node as start
	res := math.MinInt
	for i := 0; i < n; i++ {
		res = max(res, bfs(next, i))
	}
	return res

}

func bfs(next [][]int, start int) int {
	n := len(next)
	queue := make([]int, 0)
	visited := make([]int, n)

	count := 0

	// start node
	queue = append(queue, start)
	visited[start] = 1

	// loop
	for len(queue) > 0 {
		// cur
		cur := queue[0]
		queue = queue[1:]
		// update count
		count++

		// make the next move
		for _, nei := range next[cur] {
			// check if already visited
			if visited[nei] == 1 {
				continue
			}
			queue = append(queue, nei)
			visited[nei] = 1
		}
	}

	return count
}

func reachable(from, to []int) bool {
	x0, y0, r0 := from[0], from[1], from[2]
	x1, y1 := to[0], to[1]

	return r0*r0 >= (x0-x1)*(x0-x1)+(y0-y1)*(y0-y1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
