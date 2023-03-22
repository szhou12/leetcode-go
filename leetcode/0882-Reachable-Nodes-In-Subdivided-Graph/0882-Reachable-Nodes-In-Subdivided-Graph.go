package leetcode

import "container/heap"

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	// Step 1: reconstruct adj list representation
	adj := make([][]Pair, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]Pair, 0)
	}
	for _, edge := range edges {
		u, v, cnt := edge[0], edge[1], edge[2]
		// NOTE: edge weight = cnt+1
		adj[u] = append(adj[u], Pair{node: v, weight: cnt + 1})
		adj[v] = append(adj[v], Pair{node: u, weight: cnt + 1})
	}

	// Step 2: Dijkstra - find shortest path from node 0 to node i (0 -> i)
	dist := Dijkstra(n, maxMoves, adj)

	// Step 3: Count reachable nodes
	res := countNodes(n, maxMoves, edges, dist)
	return res
}

func countNodes(n int, maxMoves int, edges [][]int, dist []int) int {
	count := 0

	// Step 1: count small nodes between each edge
	for _, edge := range edges {
		u, v, cnt := edge[0], edge[1], edge[2]
		sum := 0
		if dist[u] != -1 {
			sum += maxMoves - dist[u]
		}
		if dist[v] != -1 {
			sum += maxMoves - dist[v]
		}
		count += min(cnt, sum)
	}

	// Step 2: count large nodes
	for i := 0; i < n; i++ {
		if dist[i] != -1 {
			count++
		}
	}

	return count
}

func Dijkstra(n int, maxMoves int, adj [][]Pair) []int {
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	minHeap := &PQ{}
	heap.Init(minHeap)
	// Start node = node 0
	heap.Push(minHeap, []int{0, 0})
	// Loop
	for (*minHeap).Len() > 0 {
		// current
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
			// check if reachable (VERY IMPORTANT!!!), only push reachable
			if d+weight <= maxMoves {
				heap.Push(minHeap, []int{d + weight, next})
			}
		}
	}
	return dist
}

type Pair struct {
	node   int
	weight int
}

type PQ [][]int // [[path distance, node]]

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
