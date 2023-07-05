package leetcode

func suggestedProducts(products []string, searchWord string) [][]string {
	root := newTrieNode()

	// Step 1: build Trie tree
	for _, product := range products {
		node := root

		for _, char := range product {
			letter := int(char - 'a')
			if node.children[letter] == nil {
				node.children[letter] = newTrieNode()
			}
			node = node.children[letter]
		}
		node.isEnd = true
	}

	res := make([][]string, 0)
	startNode := root
	for i := 0; i < len(searchWord); i++ {
		// 1. tranverse to the node of ith char as the start node
		prefix := searchWord[:i+1]
		letter := int(searchWord[i] - 'a')
		if startNode.children[letter] == nil {
			for j := i; j < len(searchWord); j++ {
				res = append(res, []string{})
			}
			break
		}
		startNode = startNode.children[letter]

		// 2. use DFS to fetch first 3 words that begin with the start node
		words := make([]string, 0)
		dfs(startNode, prefix, &words)
		res = append(res, words)
		
	}

	return res
}

func dfs(node *TrieNode, prefix string, res *[]string) {
	// base case: if already found 3 words, return
	if len(*res) == 3 {
		return
	}

	if node.isEnd {
		*res = append(*res, prefix)
		// Do NOT return since we may be able to traverse deeper
	}

	for i := 0; i < 26; i++ {
		if node.children[i] != nil {
			prefix += string('a' + i)
			dfs(node.children[i], prefix, res)
			prefix = prefix[:len(prefix)-1]
		}
	}


}

type TrieNode struct {
	isEnd bool
	children [26]*TrieNode
}

func newTrieNode() *TrieNode {
	node := TrieNode{isEnd: false}
	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}

	return &node
}