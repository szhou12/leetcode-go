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

	arr = append(arr, cost[cur])
	// slices.Sort() is the new feature in go 1.21. Faster then sort.Ints()
	slices.Sort(arr) // increasing order
	n := len(arr)

	// update res[cur]
	if n < 3 {
		res[cur] = 1
	} else {
		cand1 := int64(arr[n-3] * arr[n-2] * arr[n-1])
		cand2 := int64(arr[0] * arr[1] * arr[n-1])
		res[cur] = max(0, max(cand1, cand2))
	}

	// assemble subtree[cur]
	// arr at least length 1 so to have the smallest negative
	if n >= 1 {
		subtree[cur] = append(subtree[cur], arr[0])
	}
	// arr at least length 2 so to have the 2nd smallest negative without overlapping with the smallest negative
	if n >= 2 {
		subtree[cur] = append(subtree[cur], arr[1])
	}
	// arr at least length 5 so to have the 3rd largest positive without overlapping with 2 smallest negatives
	// i.e. 如果只有4个元素，第三大就是第二小，就是在重复添加
	// [ X X p p p ]
	if n >= 5 {
		subtree[cur] = append(subtree[cur], arr[n-3])
	}
	// arr at least length 4 so to have the 2nd largest positive without overlapping with 2 smallest negatives
	if n >= 4 {
		subtree[cur] = append(subtree[cur], arr[n-2])
	}
	// arr at least length 3 so to have the larest positive without overlapping with 2 smallest negatives
	if n >= 3 {
		subtree[cur] = append(subtree[cur], arr[n-1])
	}
}
