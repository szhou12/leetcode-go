package leetcode

import "container/heap"

var dir = [][]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])

	minHeap := &PQ{}
	heap.Init(minHeap)

	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	waterLevel := -1
	volume := 0

	// Start node = outward bound (periphery)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(minHeap, []int{heightMap[i][j], i, j})
			}
		}
	}

	// Loop
	for (*minHeap).Len() > 0 {
		// Current
		temp := heap.Pop(minHeap).([]int)
		height, x, y := temp[0], temp[1], temp[2]
		// check if already visited
		if visited[x][y] == 1 {
			continue
		}
		// Update
		if height > waterLevel { // the smallest height in PQ > current water level, time for the water level go up
			waterLevel = height
		}
		volume += waterLevel - height
		visited[x][y] = 1

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
			heap.Push(minHeap, []int{heightMap[dx][dy], dx, dy})
		}
	}

	return volume
}

type PQ [][]int // [[height[x][y], x, y]]

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
	*pq = (*pq)[:n-1]
	return temp
}
