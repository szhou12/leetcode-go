package leetcode

import "container/heap"

func minTimeToReach(moveTime [][]int) int {
	dirs := [][]int{
		{0, -1},
		{0, 1},
		{1, 0},
		{-1, 0},
	}

	n, m := len(moveTime), len(moveTime[0])

	visited := make([][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, m)
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// Start node
	heap.Push(minHeap, []int{0, 1, 0, 0})
	visited[0][0] = 1

	// Loop
	for minHeap.Len() > 0 {
		// pop the current
		temp := heap.Pop(minHeap).([]int)
		curTime, inc, x, y := temp[0], temp[1], temp[2], temp[3]

		// early return
		if x == n-1 && y == m-1 {
			return curTime
		}

		// make the next move
		for _, dir := range dirs {
			dx, dy := x+dir[0], y+dir[1]

			// check if out of bound
			if !(0 <= dx && dx < n && 0 <= dy && dy < m) {
				continue
			}

			// check if visited
			if visited[dx][dy] == 1 {
				continue
			}

			nextTime := max(curTime, moveTime[dx][dy]) + inc
			var nextInc int
			if inc == 1 {
				nextInc = 2
			} else {
				nextInc = 1
			}

			heap.Push(minHeap, []int{nextTime, nextInc, dx, dy})

			visited[dx][dy] = 1
		}

	}

	return -1

}

type PQ [][]int // [[time, increment, x, y], ...] where shortest-time = shortest time from (0, 0) to (x, y), increment = taking 1 or 2 seconds to make the next move

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
