package leetcode

func countPalindromePaths(parent []int, s string) int64 {
	n := len(parent)
	// Step 1: Reconstruct adj-list repr
	next := make([][]Pair, n)
	for i := 1; i <n; i++ {
		child := Pair{node: i, char: s[i]}
		next[parent[i]] = append(next[parent[i]], child)
	}

	// Step 2: DFS
	count := make(map[int]int64) // {state {path string}: count}
	state := 0 // bit masked path string for root node 0: (26x)0000...0000
	var res int64 // total valid paths
	dfs(0, -1, state, next, count, &res)

	return res
}

func dfs(cur int, parent int, state int, next [][]Pair, count map[int]int64, res *int64) {
	// Base case: no need. auto-stop at leaf node

	// (u, root) == (v, root)
	if cnts, ok := count[state]; ok {
		*res += cnts
	}
	// (u, root) differs from (v, root) by ONLY 1 char
	for i := 0; i < 26; i++ {
		s := state ^ (1 << i)
		if cnts, ok := count[s]; ok {
			*res += cnts
		}
	}

	// update count
	count[state]++

	// Recursion
	for _, nei := range next[cur] {
		nxt, ch := nei.node, nei.char
		if nxt == parent {
			continue
		}
		dfs(nxt, cur, state ^ (1 << (ch - 'a')), next, count, res)
	}
}

type Pair struct {
	node int
	char byte
}