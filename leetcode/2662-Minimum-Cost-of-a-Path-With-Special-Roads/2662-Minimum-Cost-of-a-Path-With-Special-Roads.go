package leetcode

import (
	"container/heap"
	"math"
)

func minimumCost(start []int, target []int, specialRoads [][]int) int {
	minHeap := &PQ{}
	heap.Init(minHeap)

	// Start nodes: start + end node of each special road
	heap.Push(minHeap, []int{0, encode(start[0], start[1])})
	
	// 为什么这么做是正确的? 因为题目给出条件: target node的坐标值 >= 任何一条special road的B点的坐标值。所以可以利用到性质
	for _, road := range specialRoads { // 物理意义: 不走special road, 直接计算从起点到B点的曼哈顿距离的情况
		x2, y2 := road[2], road[3]
		heap.Push(minHeap, []int{manhattanDist(start[0], start[1], x2, y2), encode(x2, y2)})
	}

	// { node i: shortest path from source to node i }
	dist := make(map[int]int)
	// shortest path from source to target
	res := math.MaxInt
	// Loop
	for (*minHeap).Len() > 0 {
		// Cur
		temp := heap.Pop(minHeap).([]int)
		d, cur := temp[0], temp[1]
		// check if already visited
		if _, ok := dist[cur]; ok {
			continue
		}
		dist[cur] = d
		curX, curY := decode(cur)
		res = min(res, d+manhattanDist(curX, curY, target[0], target[1]))

		// Make the next move
		for _, road := range specialRoads {
			x1, y1, x2, y2, cost := road[0], road[1], road[2], road[3], road[4]
			nei := encode(x2, y2)
			// check if already visited
			if _, ok := dist[nei]; ok {
				continue
			}
			heap.Push(minHeap, []int{d + manhattanDist(curX, curY, x1, y1) + cost, nei})
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func manhattanDist(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func encode(x, y int) int {
	return (x << 32) + y
}

func decode(n int) (int, int) {
	return n >> 32, n % (1 << 32)
}

type PQ [][]int // [[shortest path from source to node, node]]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
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
