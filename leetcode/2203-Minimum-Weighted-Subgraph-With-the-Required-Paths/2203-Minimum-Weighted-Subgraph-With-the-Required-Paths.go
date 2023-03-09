package leetcode

import (
	"container/heap"
	"math"
)

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	// Step 1: reconstruct graph - linked-list representation graph[fromNode] = [[toNode, weight]]
	next := make([][]Pair, n)
	for i := 0; i < n; i++ {
		next[i] = make([]Pair, 0)
	}
	prev := make([][]Pair, n)
	for i := 0; i < n; i++ {
		prev[i] = make([]Pair, 0)
	}
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		next[from] = append(next[from], Pair{node: to, weight: weight})
		prev[to] = append(prev[to], Pair{node: from, weight: weight})
	}

	// Step 2: find shortest path from source node to any node i
	dist2Src1 := Dijkstra(n, next, src1)
	dist2Src2 := Dijkstra(n, next, src2)
	dist2Dest := Dijkstra(n, prev, dest)

	// Step 3: for any node i, find min weight subgraph
	res := math.MaxInt / 3 // divides by 3 in case addition overflow
	for i := 0; i < n; i++ {
		res = min(res, dist2Src1[i]+dist2Src2[i]+dist2Dest[i])
	}
	if res < math.MaxInt/3 {
		return int64(res)
	}
	return -1

}

func Dijkstra(n int, adj [][]Pair, start int) []int {
	dist := make([]int, n)
	for i := 0; i < n; i++ { // 默认值=无法到达的情况=无限大
		dist[i] = math.MaxInt / 3
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// start node/初始值
	heap.Push(minHeap, []int{0, start})
	// loop
	for (*minHeap).Len() > 0 {
		// pop current
		temp := heap.Pop(minHeap).([]int)
		d, cur := temp[0], temp[1]
		// check if node cur already visited
		if dist[cur] < math.MaxInt/3 {
			continue
		}
		// else: update shortest path from start to cur
		dist[cur] = d

		// make next move
		for _, nei := range adj[cur] {
			next, weight := nei.node, nei.weight
			// check if node next already visited
			if dist[next] < math.MaxInt/3 {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type PQ [][]int // [[distance, node]]: shortest distance from source/初始值 to the node

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
