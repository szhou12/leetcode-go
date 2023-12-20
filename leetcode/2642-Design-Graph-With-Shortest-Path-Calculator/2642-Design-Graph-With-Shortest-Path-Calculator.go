package leetcode

import "math"

type Graph struct {
	dp [][]int
}

func Constructor(n int, edges [][]int) Graph {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				dp[i][j] = math.MaxInt / 3
			}
		}
	}

	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// directed grapth, thus only one direction
				dp[i][j] = min(dp[i][j], dp[i][from]+weight+dp[to][j])
			}
		}
	}
	graph := Graph{dp: dp}
	return graph
}

func (this *Graph) AddEdge(edge []int) {
	from, to, weight := edge[0], edge[1], edge[2]
	n := len(this.dp)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			this.dp[i][j] = min(this.dp[i][j], this.dp[i][from]+weight+this.dp[to][j])
		}
	}
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	if this.dp[node1][node2] < math.MaxInt/3 {
		return this.dp[node1][node2]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
