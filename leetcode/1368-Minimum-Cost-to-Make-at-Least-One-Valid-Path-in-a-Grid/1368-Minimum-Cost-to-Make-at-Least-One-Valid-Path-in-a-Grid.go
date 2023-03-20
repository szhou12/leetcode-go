package leetcode

import "container/heap"

// Note: the order of element matches with the value in grid[][]
// 用 HashMap 可能代码更清晰
var dir = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	cost := make([][]int, m)
	for i := 0; i < m; i++ {
		cost[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cost[i][j] = -1
		}
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// start node: [0, 0]
	heap.Push(minHeap, []int{0, 0, 0})
	// loop
	for (*minHeap).Len() > 0 {
		// current
		temp := heap.Pop(minHeap).([]int)
		d, x, y := temp[0], temp[1], temp[2]
		// check if already visited
		if cost[x][y] != -1 {
			continue
		}
		cost[x][y] = d
		// check if at the end node
		if x == m-1 && y == n-1 {
			break
		}

		// make the next step
		for i := 0; i < len(dir); i++ {
			dx, dy := x+dir[i][0], y+dir[i][1]
			// check if out of bound
			if !(0 <= dx && dx < m && 0 <= dy && dy < n) {
				continue
			}
			// check if already visited
			if cost[dx][dy] != -1 {
				continue
			}
			// Note: 用dir的index来匹配grid[][]的value
			if grid[x][y] == i+1 {
				heap.Push(minHeap, []int{d, dx, dy})
			} else {
				heap.Push(minHeap, []int{d + 1, dx, dy})
			}
		}
	}

	return cost[m-1][n-1]
}

type PQ [][]int // [arrivalCost, x, y]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] < pq[j][1]
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
