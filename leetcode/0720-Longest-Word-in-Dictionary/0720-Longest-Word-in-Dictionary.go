package leetcode

func longestWord(words []string) string {
    root := newTrieNode()

	// Build the Trie tree
	for _, word := range words {
		node := root

		for _, char := range word {
			index := int(char - 'a')
			if node.children[index] == nil {
				node.children[index] = newTrieNode()
			}
			node = node.children[index]
		}
		node.isEnd = true
	}

	// DFS
	res := ""
	dfs(root, "", &res)
	return res
}

func dfs(node *TrieNode, cur string, res *string) {
	if node == nil {
		return
	}

	for i := 0; i < 26; i++ {
		if node.children[i] != nil && node.children[i].isEnd {
			cur += string(i + 'a')
			if len(cur) > len(*res) {
				*res = cur
			}
			dfs(node.children[i], cur, res)
			cur = cur[:len(cur) - 1]
		}
	}
}

type TrieNode struct {
	isEnd bool // indicates whether it is an end of a word in the input array
	children [26]*TrieNode
}

func newTrieNode() *TrieNode {
	node := TrieNode{isEnd: false}
	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}

	return &node
}