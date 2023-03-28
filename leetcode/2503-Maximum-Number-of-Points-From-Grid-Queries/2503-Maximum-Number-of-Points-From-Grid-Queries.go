package leetcode

import (
	"container/heap"
	"sort"
)

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])

	// Step 1: sort queries in ascending order so that we start with smaller-value query
	qs := make([][]int, len(queries)) // qs = [[query, index]]
	for i, query := range queries {
		qs[i] = []int{query, i}
	}
	sort.Slice(qs, func(i, j int) bool {
		return qs[i][0] < qs[j][0]
	})

	// Step 2: For each query in order, find # cells < query by using minheap
	// Preparation
	minHeap := &PQ{}
	heap.Init(minHeap)

	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}

	count := 0
	res := make([]int, len(queries))

	// Start node = upper-left cell
	heap.Push(minHeap, []int{grid[0][0], 0, 0})

	// Loop
	for _, q := range qs {
		query, i := q[0], q[1]

		for (*minHeap).Len() > 0 && (*minHeap)[0][0] < query {
			// Current legal candidate
			temp := heap.Pop(minHeap).([]int)
			x, y := temp[1], temp[2]

			// check if already visited
			if visited[x][y] == 1 {
				continue
			}

			visited[x][y] = 1
			count++

			// Make the next move
			for _, dir := range dirs {
				dx := x + dir[0]
				dy := y + dir[1]
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

		res[i] = count
	}

	return res

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
