package leetcode

import "container/heap"

func minTimeToReach(moveTime [][]int) int {
	dirs := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	n, m := len(moveTime), len(moveTime[0])
	visited := make([][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, m)
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// Start node: (0, 0)
	heap.Push(minHeap, []int{0, 0, 0})
	visited[0][0] = 1

	// Loop
	for minHeap.Len() > 0 {
		// pop the current
		temp := heap.Pop(minHeap).([]int)
		curTime, x, y := temp[0], temp[1], temp[2]

		if x == n-1 && y == m-1 {
			return curTime
		}

		visited[x][y] = 1

		// make the next move
		for _, dir := range dirs {
			dx, dy := x+dir[0], y+dir[1]
			// skip if out of bound
			if !(0 <= dx && dx < n && 0 <= dy && dy < m) {
				continue
			}
			// skip if visited
			if visited[dx][dy] == 1 {
				continue
			}

			nextTime := max(curTime, moveTime[dx][dy]) + 1
			heap.Push(minHeap, []int{nextTime, dx, dy})

			visited[dx][dy] = 1
		}
	}

	return -1 // no actual meaning, just to make the compiler happy

}

type PQ [][]int // [[time, x, y], ...] where time = shortest time from (0, 0) to (x, y)

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
