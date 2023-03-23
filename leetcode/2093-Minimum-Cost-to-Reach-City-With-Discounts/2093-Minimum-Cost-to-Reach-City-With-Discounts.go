package leetcode

import "container/heap"

func minimumCost(n int, highways [][]int, discounts int) int {
	// Step 1: reconstruct adj list representation
	adj := make([][]Pair, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]Pair, 0)
	}
	for _, edge := range highways {
		u, v, price := edge[0], edge[1], edge[2]
		adj[u] = append(adj[u], Pair{node: v, weight: price})
		adj[v] = append(adj[v], Pair{node: u, weight: price})
	}

	// Step 2: Dijkstra - find shortest path
	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, discounts+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < discounts+1; j++ {
			cost[i][j] = -1
		}
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// Start node = node 0
	heap.Push(minHeap, []int{0, 0, 0})
	// Loop
	for (*minHeap).Len() > 0 {
		// Current
		temp := heap.Pop(minHeap).([]int)
		d, cur, discnts := temp[0], temp[1], temp[2]

		// if cur == dst, return
		if cur == n-1 {
			return d
		}

		// check if already visited
		if cost[cur][discnts] != -1 {
			continue
		}
		cost[cur][discnts] = d

		// Make the next move
		for _, nei := range adj[cur] {
			next, weight := nei.node, nei.weight

			// Don't use discount + Not already visited
			if cost[next][discnts] == -1 {
				heap.Push(minHeap, []int{d + weight, next, discnts})
			}

			// Use 1 discount + Not already visited
			if discnts+1 <= discounts && cost[next][discnts+1] == -1 {
				heap.Push(minHeap, []int{d + weight/2, next, discnts + 1})
			}
		}
	}

	return -1
}

type Pair struct {
	node   int
	weight int
}

type PQ [][]int // [[path cost, node, # discounts used]]
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
