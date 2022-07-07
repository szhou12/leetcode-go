package leetcode

import "sort"

type Pair struct {
	score int
	node  int
}

func maximumScore(scores []int, edges [][]int) int {
	n := len(scores) // total # of nodes
	neighbors := make([][]Pair, n)

	for _, edge := range edges { // O(E)
		a := edge[0] // node a
		b := edge[1] // node b

		neighbors[a] = append(neighbors[a], Pair{scores[b], b})
		neighbors[b] = append(neighbors[b], Pair{scores[a], a})
	}

	// Greedy part: only keep top 3 neighbor nodes who have the highest scores
	for node := 0; node < n; node++ { // O(V)
		sort.Slice(neighbors[node], func(i, j int) bool { // O(VlogV) ?
			return neighbors[node][i].score > neighbors[node][j].score
		})
		if len(neighbors[node]) > 3 {
			neighbors[node] = neighbors[node][:3]
		}
	}

	res := -1                    // -1 if no such sequence exists
	for _, edge := range edges { // O(E)
		a := edge[0] // node a
		b := edge[1] // node b

		for _, i := range neighbors[a] { // O(3)
			for _, j := range neighbors[b] { // O(3)
				if i.node != b && i.node != j.node && j.node != a {
					res = max(res, i.score+scores[a]+scores[b]+j.score)
				}
			}
		}
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
