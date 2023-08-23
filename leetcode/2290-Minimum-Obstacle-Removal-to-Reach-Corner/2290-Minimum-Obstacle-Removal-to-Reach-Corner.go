package leetcode

import "container/heap"

var dir = [][]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

// My Solution: Dijkstra
func minimumObstacles_dij(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	minHeap := &PQ{}
	heap.Init(minHeap)

	visited := make([][]int, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, n)
	}

	// Start node: (0, 0)
	heap.Push(minHeap, []int{0, 0, 0, 0})
	visited[0][0] = 1

	// Loop
	for (*minHeap).Len() > 0 {
		// cur
		cur := heap.Pop(minHeap).([]int)
		x, y, cost, timestamp := cur[0], cur[1], cur[2], cur[3]
		// check if at destination
		if x == m-1 && y == n-1 {
			return cost
		}

		// Make the next move
		for k := 0; k < 4; k++ {
			dx := x + dir[k][0]
			dy := y + dir[k][1]
			// check if out of bound
			if !(0 <= dx && dx < m && 0 <= dy && dy < n) {
				continue
			}
			// check if already visited
			if visited[dx][dy] == 1 {
				continue
			}

			// if encounter obstacle, add cost
			if grid[dx][dy] == 1 {
				heap.Push(minHeap, []int{dx, dy, cost + 1, timestamp + 1})
			} else if grid[dx][dy] == 0 {
				heap.Push(minHeap, []int{dx, dy, cost, timestamp + 1})
			}
			visited[dx][dy] = 1
		}
	}

	// this return has no meaning, just to pass the compiler
	return 0
}

type PQ [][]int // [[i, j, min cost arrival, timestamp]]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	if pq[i][2] == pq[j][2] {
		return pq[i][3] < pq[j][3]
	}
	return pq[i][2] < pq[j][2]
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
