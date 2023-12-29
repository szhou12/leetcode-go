package leetcode

import "math"

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	// STEP 1: Assign node ID + Build Trie
	/* reverse each string in original & changed */
	for i, s := range original {
		original[i] = reverseStr(s)
	}
	for i, s := range changed {
		changed[i] = reverseStr(s)
	}
	/* END */

	// deduplication
	strs := make(map[string]bool)
	for i, _ := range original {
		org := original[i]
		chd := changed[i]
		if _, ok := strs[org]; !ok {
			strs[org] = true
		}
		if _, ok := strs[chd]; !ok {
			strs[chd] = true
		}
	}

	// assign ID + build Trie
	m := len(strs)                  // total number of nodes
	nodeMap := make(map[string]int) // {string : node ID}
	id := 0
	root := newTrieNode()
	for str, _ := range strs { // O(V * h)
		nodeMap[str] = id

		/* build Trie */
		node := root
		for _, ch := range str {
			letterIdx := int(ch - 'a')
			if node.children[letterIdx] == nil {
				node.children[letterIdx] = newTrieNode()
			}
			node = node.children[letterIdx]
		}
		node.id = id
		/* END */

		id++
	}

	// STEP 2: Floyd
	// dist[i][j] := the minimum cost to travel from node i to node j
	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i != j {
				dist[i][j] = math.MaxInt / 3
			}
		}
	}

	for k, _ := range original { // O(n^3)
		from := nodeMap[original[k]]
		to := nodeMap[changed[k]]
		weight := cost[k]
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				dist[i][j] = min(dist[i][j], dist[i][from]+weight+dist[to][j])
			}
		}
	}

	// STEP 3: DP
	// dp[i] := the minimum cost to convert source[:i] to target[:i]
	n := len(source)
	dp := make([]int, n+1)
	source = "#" + source
	target = "#" + target
	// base case
	dp[0] = 0
	// recurrence
	for i := 1; i <= n; i++ { // O(n^2 * h)
		dp[i] = math.MaxInt / 3
		if source[i] == target[i] {
			dp[i] = dp[i-1]
		}

		/* reversely loop j + traverse Trie */
		srcNode := root
		tarNode := root
		for j := i; j >= 1; j-- {
			if srcNode.children[int(source[j]-'a')] == nil {
				break
			}
			if tarNode.children[int(target[j]-'a')] == nil {
				break
			}
			srcNode = srcNode.children[int(source[j]-'a')]
			tarNode = tarNode.children[int(target[j]-'a')]
			if srcNode.id != -1 && tarNode.id != -1 {
				dp[i] = min(dp[i], dp[j-1]+dist[srcNode.id][tarNode.id])
			}
		}
		/* END */
	}

	if dp[n] >= math.MaxInt/3 {
		return -1
	}
	return int64(dp[n])

}

func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type TrieNode struct {
	id       int // work as isEnd
	children [26]*TrieNode
}

func newTrieNode() *TrieNode {
	node := TrieNode{id: -1}
	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}
	return &node
}
