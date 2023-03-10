package leetcode

var M = int(1e9 + 7)

func countRestrictedPaths(n int, edges [][]int) int {
	// Step 1: reconstruct adj-list representation graph
	adj := make([][]Pair, n+1)
	for i := 0; i < len(adj); i++ {
		adj[i] = make([]Pair, 0)
	}
	for _, edge := range edges {
		u, v, weight := edge[0], edge[1], edge[2]
		adj[u] = append(adj[u], Pair{node: v, weight: weight})
		adj[v] = append(adj[v], Pair{node: u, weight: weight})
	}

	// Step 2: Dijkstra - find shortest path between any node i and node n (last node)
	dist := Dijkstra(n, adj)

	// Step 3: DFS - count # of restricted ways
	memo := make([]int, n+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	return dfs(1, dist, adj, &memo)
}

func dfs(cur int, dist []int, adj [][]Pair, memo *[]int) int {

}

func Dijkstra(n int, adj [][]Pair) []int {

}

type Pair struct {
	node   int
	weight int
}
