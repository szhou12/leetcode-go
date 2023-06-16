package leetcode

import (
	"container/heap"
	"math"
)

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	// Step 1: Reconstruct adj-list repr + Modify -1 -> 1
	next := make([]map[int]int, n) // [{to, weight}]
	for i := 0; i < n; i++ {
		next[i] = make(map[int]int)
	}
	// mark edges with negative weights
	negEdges := make([]map[int]int, n) // [{to, weight}]
	for i := 0; i < n; i++ {
		negEdges[i] = make(map[int]int)
	}
	for _, edge := range edges {
		u, v, weight := edge[0], edge[1], edge[2]
		if weight == -1 {
			weight = 1
			negEdges[u][v] = 1
			negEdges[v][u] = 1
		}
		next[u][v] = weight
		next[v][u] = weight
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// Step 2: Dijkstra: dest -> each node
	dist2 := make([]int, n)
	for i := 0; i < n; i++ {
		dist2[i] = math.MaxInt / 3
	}
	// Start node = destination
	heap.Push(minHeap, []int{0, destination})
	// Loop
	for (*minHeap).Len() > 0 {
		// Cur
		temp := heap.Pop(minHeap).([]int)
		d, cur := temp[0], temp[1]
		// check if already visited
		if dist2[cur] != math.MaxInt/3 {
			continue
		}
		dist2[cur] = d

		// Make the next move
		for nei, weight := range next[cur] {
			// check if already visited
			if dist2[nei] != math.MaxInt/3 {
				continue
			}
			heap.Push(minHeap, []int{d + weight, nei})
		}
	}

	// Step 3: Dijkstra: source -> each node
	dist1 := make([]int, n)
	for i := 0; i < n; i++ {
		dist1[i] = math.MaxInt / 3
	}
	// Start node = source
	heap.Push(minHeap, []int{0, source})
	// Loop
	for (*minHeap).Len() > 0 {
		// Cur
		temp := heap.Pop(minHeap).([]int)
		d, cur := temp[0], temp[1]
		// check if already visited
		if dist1[cur] != math.MaxInt/3 {
			continue
		}
		dist1[cur] = d
		// check if arrive at destination & path cost == target
		if cur == destination && d != target {
			return [][]int{}
		}

		// Make the next move
		for nei, weight := range next[cur] {
			// check if already visited
			if dist1[nei] != math.MaxInt/3 {
				continue
			}
			// increase edge weight if < target
			if negEdges[cur][nei] == 1 && dist1[cur] + weight + dist2[nei] < target {
				weight = target - dist1[cur] - dist2[nei]
				// update adj-list
				next[cur][nei] = weight
				next[nei][cur] = weight
			}
			heap.Push(minHeap, []int{d + weight, nei})
		}
	}

	// Step 4: Re-assign edge weight
	for i := 0; i < len(edges); i++ {
		u, v := edges[i][0], edges[i][1]
		edges[i][2] = next[u][v]
	}

	return edges

}

type PQ [][]int // [[shortest distance, node]]

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
