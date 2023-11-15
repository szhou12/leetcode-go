package leetcode

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	n := len(values)
	// Step 1: Reconstruct adj-list representation of the tree
	next := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)
	}

	// Step 2: DFS
	subtreeSum := make([]int, n)
	dfsSubtree(0, -1, next, values, subtreeSum)

	maxSum := dfs(0, -1, next, values, subtreeSum)

	return int64(maxSum)

}

// dfsSubtree(cur) = total score that can be collected from the cur-rooted subtree
func dfsSubtree(cur int, parent int, next [][]int, values []int, subtreeSum []int) int {
	// base case: no code needed. auto-stop when at leaf

	// recursion
	sum := values[cur]
	for _, nei := range next[cur] {
		if nei == parent {
			continue
		}
		sum += dfsSubtree(nei, cur, next, values, subtreeSum)
	}

	subtreeSum[cur] = sum

	return sum
}

// dfs(cur) = max score that can be collected from the cur-rooted subtree while keeping the subtree 'healthy'
// max score = max(case 1, case 2)
//  case 1: Do NOT take the root (cur), all children nodes can be taken
//  case 2: Do take the root (cur), each path to leaf node must have one node not taken (recursive call)
func dfs(cur int, parent int, next [][]int, values []int, subtreeSum []int) int {
	// base case: if only leaf node left, MUST NOT take it, otherwise, not 'healthy' for leaf-rooted subtree
	// Note: must check cur != 0 in case of the structure: [root] -...- [leaf]
	if len(next[cur]) == 1 && cur != 0 {
		return 0
	}

	// recursion: take the root (cur)
	sum := values[cur]
	for _, nei := range next[cur] {
		if nei == parent {
			continue
		}
		sum += dfs(nei, cur, next, values, subtreeSum)
	}

	return max(sum, subtreeSum[cur] - values[cur])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}