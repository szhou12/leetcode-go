package networkflow

import "fmt"

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



