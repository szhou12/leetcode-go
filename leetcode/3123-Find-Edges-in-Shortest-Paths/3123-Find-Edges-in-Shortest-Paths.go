package leetcode

import (
	"container/heap"
	"math"
)

func findAnswer(n int, edges [][]int) []bool {
	// Step 1: Reconstruct adj-list representation of the graph
	next := make([][]Pair, n)
	for i := 0; i < n; i++ {
		next[i] = make([]Pair, 0)
	}
	for _, edge := range edges {
		a, b, weight := edge[0], edge[1], edge[2]
		next[a] = append(next[a], Pair{node: b, weight: weight})
		next[b] = append(next[b], Pair{node: a, weight: weight})
	}

	// Step 2: Dijkstra
	ds := Dijkstra(n, next, 0)   // shortest path from node 0 to any nodes
	de := Dijkstra(n, next, n-1) // shortest path from node n-1 to any nodes

	// Step 3: d(0, a) + w(a, b) + d(b, n-1) == d(0, n-1)
	res := make([]bool, len(edges))
	for i, edge := range edges {
		a, b, wei := edge[0], edge[1], edge[2]
		if ds[a]+wei+de[b] == ds[n-1] || ds[b]+wei+de[a] == ds[n-1] {
			res[i] = true
		}
	}
	return res
}

func Dijkstra(n int, next [][]Pair, start int) []int {
	dist := make([]int, n) // shortest distance from node start to any nodes
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt / 3
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// start node
	heap.Push(minHeap, []int{0, start})
	// loop
	for minHeap.Len() > 0 {
		// pop current
		temp := heap.Pop(minHeap).([]int)
		d, cur := temp[0], temp[1]
		// check if already visited
		if dist[cur] != math.MaxInt/3 {
			continue
		}
		// update shortest distance from source to current node
		dist[cur] = d

		// make next move
		for _, nei := range next[cur] {
			node, weight := nei.node, nei.weight
			// check if already visited
			if dist[node] != math.MaxInt/3 {
				continue
			}
			heap.Push(minHeap, []int{d + weight, node})
		}
	}

	return dist
}

type Pair struct {
	node   int
	weight int
}

type PQ [][]int // [[shortest path, node i]] := shortest path from source node to node i

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
