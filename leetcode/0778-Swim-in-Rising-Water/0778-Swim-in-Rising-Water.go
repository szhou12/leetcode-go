package leetcode

import "container/heap"

var dir = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func swimInWater(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// Start node = grid[0][0]
	heap.Push(minHeap, []int{grid[0][0], 0, 0})
	time := grid[0][0]

	// Loop
	for (*minHeap).Len() > 0 {
		// Current node
		temp := heap.Pop(minHeap).([]int)
		horizon, x, y := temp[0], temp[1], temp[2]
		// check if already visited
		if visited[x][y] == 1 {
			continue
		}
		visited[x][y] = 1
		time = max(time, horizon)

		// return if at destination
		if x == m-1 && y == n-1 {
			return time
		}

		// Make the next move
		for i := 0; i < len(dir); i++ {
			dx := x + dir[i][0]
			dy := y + dir[i][1]

			// check if out-of-bound
			if !(0 <= dx && dx < m && 0 <= dy && dy < n) {
				continue
			}
			// check if already visited
			if visited[dx][dy] == 1 {
				continue
			}

			heap.Push(minHeap, []int{grid[dx][dy], dx, dy})
		}
	}

	return -1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type PQ [][]int // [[grid value, x, y]]

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
