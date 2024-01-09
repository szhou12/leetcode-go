package leetcode

import "slices"

func placedCoins(edges [][]int, cost []int) []int64 {
	n := len(cost)
	// Step 1: convert edges to adj-list representation
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)
	}

	// subtree[i] := 3 largest positives + 2 smallest negatives in subtree rooted at i
	subtree := make([][]int, n)
	for i := 0; i < n; i++ {
		subtree[i] = make([]int, 0)
	}

	res := make([]int64, n)

	// Step 2: DFS
	dfs(0, -1, next, cost, subtree, res)
	return res

}

func dfs(cur int, parent int, next [][]int, cost []int, subtree [][]int, res []int64) {
	// arr := 3 largest positives + 2 smallest negatives in subtree rooted at cur
	arr := make([]int, 0)

	for _, child := range next[cur] {
		if child == parent {
			continue
		}
		dfs(child, cur, next, cost, subtree, res)
		// collect 3 largest positives + 2 smallest from each child
		for _, v := range subtree[child] {
			arr = append(arr, v)
		}
	}

	// slices.Sort() is the new feature in go 1.21. Faster then sort.Ints()
	slices.Sort(arr)
	
	
}
