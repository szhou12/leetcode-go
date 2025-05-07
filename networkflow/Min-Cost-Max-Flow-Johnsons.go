package networkflow

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	from         int
	to           int
	reverseEdge  *Edge
	capacity     int
	cost         int // will be modified by Johnson's potential in order to use Dijkstra, thus no longer the original cost to calculate min cost
	originalCost int
	flow         int
}

func NewEdge(from int, to int, reverseEdge *Edge, capacity int, cost int) *Edge {
	return &Edge{
		from:         from,
		to:           to,
		reverseEdge:  reverseEdge,
		capacity:     capacity,
		cost:         cost,
		originalCost: cost,
		flow:         0,
	}
}

func (e *Edge) IsReverse() bool {
	return e.capacity == 0
}

func (e *Edge) RemainingCapacity() int {
	return e.capacity - e.flow
}

func (e *Edge) Augment(bottleneck int) {
	e.flow += bottleneck
	e.reverseEdge.flow -= bottleneck
}

func (e Edge) String() string {
	return fmt.Sprintf("Edge{ %d -> %d | capacity: %d | cost: %d flow: %d | isReverse: %t}",
		e.from, e.to, e.capacity, e.flow, e.cost, e.IsReverse())
}

// edges = [[from, to, capacity, cost], ...]
func MinCostMaxFlowJohnson(n int, edges [][]int, s int, t int) (int, int) {
	maxFlow, minCost := 0, 0

	// init an empty graph with n nodes including the source s and sink t
	graph := make([][]*Edge, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]*Edge, 0)
	}

	// Construct residual graph
	// add original edges
	// add reverse edges
	for _, edge := range edges {
		from, to, capacity, cost := edge[0], edge[1], edge[2], edge[3]
		forwardEdge := NewEdge(from, to, nil, capacity, cost)
		backwardEdge := NewEdge(to, from, nil, 0, -cost)
		forwardEdge.reverseEdge = backwardEdge
		backwardEdge.reverseEdge = forwardEdge
		graph[from] = append(graph[from], forwardEdge)
		graph[to] = append(graph[to], backwardEdge)
	}

	// dist[i] := shortest path from source s to node i
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		// dist[i] = math.MaxInt // use MaxInt will cause overflow in Bellman-Ford
		dist[i] = int(10e7)
	}
	dist[s] = 0

	// Bellman-Ford: find shortest path from source to any other node
	// O(VE)
	for i := 0; i < n-1; i++ {
		// relax each eage
		for _, edges := range graph {
			for _, e := range edges {
				if e.RemainingCapacity() > 0 && dist[e.from]+e.cost < dist[e.to] {
					dist[e.to] = dist[e.from] + e.cost
				}
			}
		}
	}


	// Johnson's reweighting: adjust edge costs to be non-negative so that Dijkstra can be applied later
	for from := 0; from < n; from++ {
		for _, edge := range graph[from] {
			if edge.RemainingCapacity() > 0 {
				edge.cost += dist[from] - dist[edge.to]
			} else {
				edge.cost = 0
			}
		}
	}

	// Ford-Fulkerson
	for {
		path, ok := getAugmentingPath(graph, s, t)
		if !ok {
			break
		}

		// find bottleneck
		bottleneck := math.MaxInt
		for _, edge := range path {
			bottleneck = min(bottleneck, edge.RemainingCapacity())
		}


		// augment the path
		for _, edge := range path {
			edge.Augment(bottleneck)
			minCost += bottleneck * edge.originalCost
		}

		maxFlow += bottleneck
	}

	return maxFlow, minCost
}

func getAugmentingPath(graph [][]*Edge, s int, t int) ([]*Edge, bool) {
	n := len(graph)

	prev := make([]*Edge, n)

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = int(10e7)
	}

	minHeap := &PQ{}
	heap.Init(minHeap)

	// Start node
	heap.Push(minHeap, []int{0, s}) // [distance, node]

	// Loop
	for minHeap.Len() > 0 {
		// pop current
		temp := heap.Pop(minHeap).([]int)
		distance, cur := temp[0], temp[1]

		// check if cur already visited
		if dist[cur] != int(10e7) {
			continue
		}

		// update shorest distance on cur
		dist[cur] = distance

		// make the next move
		for _, neiEdge := range graph[cur] {
			to, cost := neiEdge.to, neiEdge.cost

			// check if neighbor node is already visited
			if dist[to] != int(10e7) {
				continue
			}

			if neiEdge.RemainingCapacity() > 0 {
				prev[to] = neiEdge
				heap.Push(minHeap, []int{distance + cost, to})
			}
		}

	}

	path := make([]*Edge, 0)

	// check if t is reached
	if dist[t] == int(10e7) {
		return path, false
	}

	// Johnson's reweighting again: adjust edge costs
	// path augmentation will change the residual graph structure
	// The previous distances (used for reweighting) are no longer valid because:
	// 1. Some edges might be saturated (remaining capacity = 0)
	// 2. New reverse edges might have become available (remaining capacity > 0), and since reverse edges have -cost, we need to reweight them in order to use Dijkstra later
	// 3. The shortest paths in the residual graph have changed
	for from := 0; from < n; from++ {
		for _, edge := range graph[from] {
			if edge.RemainingCapacity() > 0 {
				edge.cost += dist[from] - dist[edge.to]
			} else {
				edge.cost = 0
			}
		}
	}

	// backtrack to get the path
	node := t
	for node != s {
		path = append(path, prev[node])
		node = prev[node].from
	}

	// reverse the path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, true
}

type PQ [][]int // [[shortest path, node i], ...]

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
