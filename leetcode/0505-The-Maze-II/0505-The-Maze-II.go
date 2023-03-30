package leetcode

import "container/heap"

var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func shortestDistance(maze [][]int, start []int, destination []int) int {
	m, n := len(maze), len(maze[0])

	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dist[i][j] = -1
		}
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// Start node = start
	heap.Push(minHeap, []int{0, start[0], start[1]})

	// Loop
	for (*minHeap).Len() > 0 {
		// Current node
		temp := heap.Pop(minHeap).([]int)
		d, x, y := temp[0], temp[1], temp[2]

		// return if arrive at destination
		if x == destination[0] && y == destination[1] {
			return d
		}

		// check if already visited
		if dist[x][y] != -1 {
			continue
		}
		// Update
		dist[x][y] = d

		// Make the next move
		// dx0, dy0 := x+dirX, y+dirY
		// if (dirX != 0 && dirY != 0) && inBoard(maze, dx0, dy0) && dist[dx0][dy0] == -1 && maze[dx0][dy0] != 1 {
		// 	heap.Push(minHeap, []int{d + 1, dx0, dy0, dirX, dirY})
		// } else {
		// 	for i := 0; i < len(dir); i++ {
		// 		dx := x + dir[i][0]
		// 		dy := y + dir[i][1]

		// 		// check if out-of-bound
		// 		if !inBoard(maze, dx, dy) {
		// 			continue
		// 		}
		// 		// check if already visited
		// 		if dist[dx][dy] != -1 {
		// 			continue
		// 		}
		// 		// check if wall
		// 		if maze[dx][dy] == 1 {
		// 			continue
		// 		}
		// 		heap.Push(minHeap, []int{d + 1, dx, dy, dir[i][0], dir[i][1]})
		// 	}
		// }
		for i := 0; i < len(dir); i++ {
			dx, dy := x, y
			steps := 0

			// keep going until hit the wall OR out-of-bound
			for inBoard(maze, dx+dir[i][0], dy+dir[i][1]) && maze[dx+dir[i][0]][dy+dir[i][1]] != 1 {
				dx += dir[i][0]
				dy += dir[i][1]
				steps++
			}

			// check if already visited
			if dist[dx][dy] != -1 {
				continue
			}

			heap.Push(minHeap, []int{d + steps, dx, dy})
		}
	}

	return -1

}

func inBoard(maze [][]int, x int, y int) bool {
	m, n := len(maze), len(maze[0])
	return 0 <= x && x < m && 0 <= y && y < n
}

type PQ [][]int // [[distance cost, x, y]]

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
