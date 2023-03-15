package leetcode

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Step 1: reconstruct adj-list representation
	adj := make([][]Pair, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]Pair, 0)
	}
	for i := 0; i < len(edges); i++ {
		from, to := edges[i][0], edges[i][1]
		weight := succProb[i]
		adj[from] = append(adj[from], Pair{node: to, weight: weight})
		adj[to] = append(adj[to], Pair{node: from, weight: weight})
	}

	// Step 2: find the 'shortest' (max prob) path from node start to node i
	dist := Dijkstra(n, adj, start)

	return dist[end]
}

type Pair struct {
	node   int
	weight float64
}

type PQ [][]float64 // [[max prob path, node]]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] > pq[j][0]
}

func (pq *PQ) Push(x interface{}) {
	(*pq) = append((*pq), x.([]float64))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	(*pq) = (*pq)[:n-1]
	return temp
}
