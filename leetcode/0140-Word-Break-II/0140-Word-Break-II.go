package leetcode

func wordBreak(s string, wordDict []string) []string {

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