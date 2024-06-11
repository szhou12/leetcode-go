package leetcode

import (
	"container/heap"
	"math"
)

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	// Step 1: Reconstruct adj-list repr + Modify -1 -> 1
	next := make([]map[int]int, n) // next[cur] = {nei_1: weight_1, nei_2: weight_2}
	for i := 0; i < n; i++ {
		next[i] = make(map[int]int)
	}
	// mark modifiable edges
	modify := make([]map[int]bool, n) // modify[cur] = {nei_1: true, nei_2: true}
	for i := 0; i < n; i++ {
		modify[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		u, v, weight := edge[0], edge[1], edge[2]
		if weight == -1 {
			weight = 1
			modify[u][v] = true
			modify[v][u] = true
		}
		next[u][v] = weight
		next[v][u] = weight
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// Step 2: Dijkstra: destination -> each node
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

		// Upddate
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

		// Update
		dist1[cur] = d

		// check if when arrive at destination, path cost == target
		// 如果到达destination时，最短路径的cost还不等于target，说明这条最短路径不途径任何一条modifiable edge
		// 因为在每次进PQ时，我们都已经提前修改了modifiable edge的边权，保证只要通路上途径它，它一定不是造成path cost != target的凶手之一。
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
			if modify[cur][nei] && dist1[cur] + weight + dist2[nei] < target {
				weight = target - dist1[cur] - dist2[nei]
				// update adj-list
				next[cur][nei] = weight
				next[nei][cur] = weight
			}
			heap.Push(minHeap, []int{d + weight, nei}) // 关键点: 下一次pop时，不一定这个调整过边权的nei是堆里的最小值，也有可能到达其他nei(cur -> nei是不可修改边)的最短距离更小，从而在下一次pop时被选中弹出。所以，要贪心地反复调整之后弹出的节点，它的nei出现modifiable edge的权重。从而“堵死”从起点到终点的最短路径上，经过modifiable edge，竟然path cost还不等于target的情况。也就是说，这条完整的最短路径，你要不然不途径任何一条modifiable edge，行，你可以有path cost != target；否则只要途径哪怕一条modifiable edge, 这个贪心思想都确保了path cost == target
		}
	}

	// Step 4: Re-assign edge weight
	for i := 0; i < len(edges); i++ {
		u, v := edges[i][0], edges[i][1]
		edges[i][2] = next[u][v]
	}

	return edges

}

type PQ [][]int // [[shortest distance from start -> node, node]]

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
