package leetcode

func suggestedProducts(products []string, searchWord string) [][]string {
	
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