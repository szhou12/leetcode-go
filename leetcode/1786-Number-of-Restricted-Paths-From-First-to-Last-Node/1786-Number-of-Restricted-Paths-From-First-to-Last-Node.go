package leetcode

import "container/heap"

var M = int(1e9 + 7)

func countRestrictedPaths(n int, edges [][]int) int {
	// Step 1: reconstruct adj-list representation graph
	adj := make([][]Pair, n+1)
	for i := 0; i < len(adj); i++ {
		adj[i] = make([]Pair, 0)
	}
	for _, edge := range edges {
		u, v, weight := edge[0], edge[1], edge[2]
		adj[u] = append(adj[u], Pair{node: v, weight: weight})
		adj[v] = append(adj[v], Pair{node: u, weight: weight})
	}

	// Step 2: Dijkstra - find shortest path between any node i and node n (last node)
	dist := Dijkstra(n, adj)

	// Step 3: DFS - count # of restricted ways
	memo := make([]int, n+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	return dfs(1, n, dist, adj, &memo)
}

func dfs(cur int, n int, dist []int, adj [][]Pair, memo *[]int) int {
	// base case
	if cur == n {
		return 1
	}

	if (*memo)[cur] != -1 {
		return (*memo)[cur]
	}

	total := 0

	for _, nei := range adj[cur] {
		next := nei.node
		if dist[cur] > dist[next] {
			total += dfs(next, n, dist, adj, memo)
			total %= M
		}
	}
	(*memo)[cur] = total
	return total
}

func Dijkstra(n int, adj [][]Pair) []int {
	dist := make([]int, n+1)
	for i := 0; i < len(dist); i++ {
		dist[i] = -1
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// start node = node n
	heap.Push(minHeap, []int{0, n})
	// loop
	for (*minHeap).Len() > 0 {
		// Pop
		temp := heap.Pop(minHeap).([]int)
		d, cur := temp[0], temp[1]
		// check if already visited
		if dist[cur] != -1 {
			continue
		}
		dist[cur] = d

		// make the next move
		for _, nei := range adj[cur] {
			next, weight := nei.node, nei.weight
			// check if already visited
			if dist[next] != -1 {
				continue
			}
			// Push
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
