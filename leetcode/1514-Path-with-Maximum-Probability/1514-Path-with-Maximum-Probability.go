package leetcode

import "container/heap"

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Step 1: reconstruct adj-list representation
	adj := make([][]Pair, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]Pair, 0)
	}
	for i := 0; i < len(edges); i++ {
		from, to := edges[i][0], edges[i][1]
		weight := succProb[i]
		adj[from] = append(adj[from], Pair{node: to, weight: weight})
		adj[to] = append(adj[to], Pair{node: from, weight: weight})
	}

	// Step 2: find the 'shortest' (max prob) path from node start to node i
	dist := Dijkstra(n, adj, start)

	return dist[end]
}

func Dijkstra(n int, adj [][]Pair, start int) []float64 {
	dist := make([]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = 0
	}

	minHeap := &PQ{}
	heap.Init(minHeap)
	// start node
	heap.Push(minHeap, []float64{1, float64(start)})
	// loop
	for (*minHeap).Len() > 0 {
		// current
		temp := heap.Pop(minHeap).([]float64)
		d, cur := temp[0], int(temp[1])
		// check if already visited
		if dist[cur] != 0 {
			continue
		}
		dist[cur] = d

		// make next move
		for _, nei := range adj[cur] {
			next, weight := nei.node, nei.weight
			// check if already visited
			if dist[next] != 0 {
				continue
			}
			heap.Push(minHeap, []float64{d * weight, float64(next)})
		}
	}
	return dist
}

type Pair struct {
	node   int
	weight float64
}

type PQ [][]float64 // [[max prob path, node]]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] > pq[j][0]
}

func (pq *PQ) Push(x interface{}) {
	(*pq) = append((*pq), x.([]float64))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	(*pq) = (*pq)[:n-1]
	return temp
}
