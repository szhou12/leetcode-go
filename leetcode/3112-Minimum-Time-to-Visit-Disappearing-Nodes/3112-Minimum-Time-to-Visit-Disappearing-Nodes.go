package leetcode

import "container/heap"

func minimumTime(n int, edges [][]int, disappear []int) []int {
	// Step 1: reconstruct adj-list representation
	next := make([][]Pair, n)
	for i := 0; i < n; i++ {
		next[i] = make([]Pair, 0)
	}
	for _, edge := range edges {
		a, b, wei := edge[0], edge[1], edge[2]
		next[a] = append(next[a], Pair{node: b, weight: wei})
		next[b] = append(next[b], Pair{node: a, weight: wei})
	}

	// Step 2: Dijkstra
	ds := make([]int, n)
	for i := 0; i < n; i++ {
		ds[i] = -1
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// start node
	heap.Push(minHeap, []int{0, 0})

	// loop
	for minHeap.Len() > 0 {
		// pop current
		temp := heap.Pop(minHeap).([]int)
		d, cur := temp[0], temp[1]
		// check if already visited
		if ds[cur] != -1 {
			continue
		}
		// check if already disappeared. CAN'T EXPAND ANYMORE
		if disappear[cur] <= d {
			ds[cur] = -1
			continue
		}
		// update valid shortest distance from source to currrent
		ds[cur] = d

		// make next move
		for _, nei := range next[cur] {
			node, weight := nei.node, nei.weight
			// check if already visited
			if ds[node] != -1 {
				continue
			}
			// push
			heap.Push(minHeap, []int{d + weight, node})
		}
	}

	return ds

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
