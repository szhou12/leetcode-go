package leetcode

import "container/heap"

var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func minimumTime(grid [][]int) int {
	// Check if able to make first move
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}

	m, n := len(grid), len(grid[0])

	// record Min Arrival Time at each cell
	arrival := make([][]int, m)
	for i := 0; i < m; i++ {
		arrival[i] = make([]int, n)
	}
	// default arrival time = -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arrival[i][j] = -1
		}
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// heap.Push(minHeap, []int{0, 0, 0}) // 初始状态替换为下面两个
	if grid[0][1] <= 1 {
		heap.Push(minHeap, []int{1, 0, 1})
	}
	if grid[1][0] <= 1 {
		heap.Push(minHeap, []int{1, 1, 0})
	}
	for (*minHeap).Len() > 0 {

		// Step 1: address current step
		cur := heap.Pop(minHeap).([]int)
		t, x, y := cur[0], cur[1], cur[2]
		if arrival[x][y] != -1 { // already visited
			continue
		}
		arrival[x][y] = t
		if x == m-1 && y == n-1 {
			break
		}

		// Step 2: address next step
		// (x, y) -> (dx, dy)
		for k := 0; k < 4; k++ {
			dx := x + dir[k][0]
			dy := y + dir[k][1]

			// check out of bound
			if !(0 <= dx && dx < m && 0 <= dy && dy < n) {
				continue
			}
			// check already visited
			if arrival[dx][dy] != -1 {
				continue
			}

			// Push: 分类讨论
			if t+1 >= grid[dx][dy] {
				heap.Push(minHeap, []int{t + 1, dx, dy})
			} else if (grid[dx][dy]-t)%2 == 1 {
				heap.Push(minHeap, []int{grid[dx][dy], dx, dy})
			} else {
				heap.Push(minHeap, []int{grid[dx][dy] + 1, dx, dy})
			}

		}
	}

	return arrival[m-1][n-1]

}

type PQ [][]int // PQ[i] = [arrivalTime, x, y]

func (pq PQ) Len() int { return len(pq) }

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
	(*pq) = (*pq)[:n-1]
	return temp
}
