package leetcode

import "container/heap"

// For later convenience, maintain order: d < l < r < u
var dir = [][]int{
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, 0},
}

func findShortestWay(maze [][]int, ball []int, hole []int) string {
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
	// Start node = ball
	heap.Push(minHeap, Tuple{cost: 0, x: ball[0], y: ball[1], route: ""})
	// Loop
	for (*minHeap).Len() > 0 {
		// Current
		temp := heap.Pop(minHeap).(Tuple)
		d, x, y, route := temp.cost, temp.x, temp.y, temp.route
		// check if already visited
		if dist[x][y] != -1 {
			continue
		}
		// Update
		dist[x][y] = d

		// early return
		if x == hole[0] && y == hole[1] {
			return route
		}

		// Make the next move
		for i := 0; i < len(dir); i++ {
			dx, dy := x, y
			steps := 0

			// keep going until will out-of-bound OR will hit the wall
			for inBoard(maze, dx+dir[i][0], dy+dir[i][1]) && maze[dx+dir[i][0]][dy+dir[i][1]] != 1 {
				dx += dir[i][0]
				dy += dir[i][1]
				steps++

				// if arrive at hole, BREAK!!!
				if dx == hole[0] && dy == hole[1] {
					break
				}
			}

			// check if already visited
			if dist[dx][dy] != -1 {
				continue
			}

			if i == 0 {
				heap.Push(minHeap, Tuple{d + steps, dx, dy, route + "d"})
			} else if i == 1 {
				heap.Push(minHeap, Tuple{d + steps, dx, dy, route + "l"})
			} else if i == 2 {
				heap.Push(minHeap, Tuple{d + steps, dx, dy, route + "r"})
			} else {
				heap.Push(minHeap, Tuple{d + steps, dx, dy, route + "u"})
			}

		}

	}

	return "impossible"
}

func inBoard(maze [][]int, x int, y int) bool {
	m, n := len(maze), len(maze[0])
	return 0 <= x && x < m && 0 <= y && y < n
}

type Tuple struct {
	cost  int
	x     int
	y     int
	route string
}

type PQ []Tuple // [Tuple{path cost, x, y, route}]

func (pq PQ) Len() int {
	return len(pq)
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq PQ) Less(i, j int) bool {
	if pq[i].cost == pq[j].cost {
		return pq[i].route < pq[j].route
	}
	return pq[i].cost < pq[j].cost
}
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Tuple))
}
func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}
