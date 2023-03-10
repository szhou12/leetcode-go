package leetcode

var M = int(1e9 + 7)

func countPaths(n int, roads [][]int) int {
	// Step 1: reconstruct adj-list representation graph

	// Step 2: Dijkstra - find shortest path from node 0 (start node) to any node i

	// Step 3: DFS + Memo - count # of ways going from node 0 to node n-1 by shortest path
}
