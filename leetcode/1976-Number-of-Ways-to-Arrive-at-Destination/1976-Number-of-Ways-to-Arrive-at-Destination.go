package leetcode

import "container/heap"

var M = int(1e9 + 7)

func countPaths(n int, roads [][]int) int {
	// Step 1: reconstruct adj-list representation graph
	adj := make([][]Pair, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]Pair, 0)
	}
	for _, edge := range roads {
		from, to, weight := edge[0], edge[1], edge[2]
		adj[from] = append(adj[from], Pair{node: to, weight: weight})
		adj[to] = append(adj[to], Pair{node: from, weight: weight})
	}

	// Step 2: Dijkstra - find shortest path from node 0 (start node) to any node i
	dist := Dijkstra(n, adj)

	// Step 3: DFS + Memo - count # of ways going from node 0 to node n-1 by shortest path
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	return dfs(n-1, dist[n-1], dist, adj, &memo)
}

/*
 Count # of ways going from node 0 to destination (input node) by shortest path (input distance)
*/
func dfs(node int, distance int, dist []int, adj [][]Pair, memo *[]int) int {
	// if unable to go from node 0 to i by distance, return 0
	if dist[node] != distance {
		return 0
	}
	// base case: if node=start node, return 1 (only 1 way to go from start node to start node)
	if node == 0 {
		return 1
	}
	// Memo: if # paths from node 0 to i already counted, directly return memo[i]
	if (*memo)[node] != -1 {
		return (*memo)[node]
	}

	total := 0
	for _, nei := range adj[node] {
		next, weight := nei.node, nei.weight
		total += dfs(next, distance-weight, dist, adj, memo)
		total %= M
	}
	// record total counts at current level
	(*memo)[node] = total
	return total
}

func Dijkstra(n int, adj [][]Pair) []int {
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// start node
	heap.Push(minHeap, []int{0, 0})
	// loop
	for (*minHeap).Len() > 0 {
		// current node
		temp := heap.Pop(minHeap).([]int)
		d, cur := temp[0], temp[1]
		// check if already visited
		if dist[cur] != -1 {
			continue
		}
		dist[cur] = d

		// make next move
		for _, nei := range adj[cur] {
			next, weight := nei.node, nei.weight
			// check if alreay visited
			if dist[next] != -1 {
				continue
			}
			heap.Push(minHeap, []int{d + weight, next})
		}
	}
	return dist
}

type Pair struct {
	node   int
	weight int
}

type PQ [][]int // [[shortest path, node]]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq *PQ) Push(x interface{}) {
	(*pq) = append((*pq), x.([]int))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	(*pq) = (*pq)[:n-1]
	return temp
}
