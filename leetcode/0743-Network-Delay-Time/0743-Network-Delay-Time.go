package leetcode

import "container/heap"

func networkDelayTime(times [][]int, n int, k int) int {
	// Step 1: reconstruct adj-list representation
	adj := make([][]Pair, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = make([]Pair, 0)
	}
	for _, edge := range times {
		from, to, weight := edge[0], edge[1], edge[2]
		adj[from] = append(adj[from], Pair{node: to, weight: weight})
	}

	// Step 2: Dijkstra - find shortest path from node k to any node i
	dist := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = -1
	}

	minHeap := &PQ{}
	heap.Init(minHeap)
	// Start node = node k
	heap.Push(minHeap, []int{0, k})
	// Loop
	for (*minHeap).Len() > 0 {
		// Current node
		temp := heap.Pop(minHeap).([]int)
		d, cur := temp[0], temp[1]
		// check if alredy visited
		if dist[cur] != -1 {
			continue
		}
		dist[cur] = d

		// Make the next move
		for _, nei := range adj[cur] {
			next, weight := nei.node, nei.weight
			// check if already visited
			if dist[next] != -1 {
				continue
			}
			heap.Push(minHeap, []int{d + weight, next})
		}
	}

	// Step 3: calculate min time for signals to reach all nodes
	res := -1
	for i := 1; i <= n; i++ {
		if dist[i] == -1 {
			return -1
		}
		res = max(res, dist[i])
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Pair struct {
	node   int
	weight int
}

type PQ [][]int // [[path cost, node]]

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
	*pq = append(*pq, x.([]int))
}
func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp

}
