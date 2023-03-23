package leetcode

import "container/heap"

// Solution idea 1 - Dijkstra
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// Step 1: reconstruct adj list representation
	adj := make([][]Pair, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]Pair, 0)
	}
	for _, edge := range flights {
		from, to, price := edge[0], edge[1], edge[2]
		adj[from] = append(adj[from], Pair{node: to, weight: price})
	}

	// Step 2: Dijkstra - find shortest path (min cost + <= k+1 stops) from node start to any node i
	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, k+2) // [0, ..., k+1] so length=k+2
	}
	for i := 0; i < n; i++ {
		for j := 0; j < k+2; j++ {
			cost[i][j] = -1
		}
	}
	minHeap := &PQ{}
	heap.Init(minHeap)

	// Start node = [src node, 0]
	heap.Push(minHeap, []int{0, src, 0})
	// Loop
	for (*minHeap).Len() > 0 {
		// Current node
		temp := heap.Pop(minHeap).([]int)
		d, cur, stops := temp[0], temp[1], temp[2]

		// Return: if cur == dst, it means we arrive at the destination with min cost within k+1 stops => This MUST be answer
		if cur == dst {
			return d
		}

		// check if already visited
		if cost[cur][stops] != -1 {
			continue
		}
		// check if already used up k+1 stops
		if stops == k+1 {
			continue
		}
		cost[cur][stops] = d

		// Make the next move
		for _, nei := range adj[cur] {
			next, weight := nei.node, nei.weight
			// check if already visited
			if cost[next][stops+1] != -1 {
				continue
			}
			heap.Push(minHeap, []int{d + weight, next, stops + 1})
		}
	}

	// We have return inside the loop, it means we MUST have reached the dst within the loop
	// if we are out of the loop and have not returned, it means dst is NOT reachable
	return -1

}

type Pair struct {
	node   int
	weight int
}

type PQ [][]int // [path cost, node, # stops used]

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
