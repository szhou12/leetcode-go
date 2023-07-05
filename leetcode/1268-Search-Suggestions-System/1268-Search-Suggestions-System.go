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