package leetcode

type TrieNode struct {
	isEnd bool
	children [27]*TrieNode // 26 letters + '{'
}

func newTrieNode() *TrieNode {
	node := TrieNode{isEnd: false}
	for i := 0; i < 27; i++ {
		node.children[i] = nil
	}
	return &node
}

func buildTrie(root *TrieNode, word string) {
	node := root
	for _, char := range word {
		letter := int(char - 'a')
		if node.children[letter] == nil {
			node.children[letter] = newTrieNode()
		}
		node = node.children[letter]
	}
	node.isEnd = true
}

type WordFilter struct {
	root *TrieNode
}

func Constructor(words []string) WordFilter {
    trie := WordFilter{root: newTrieNode()}
	return trie
}


func (this *WordFilter) F(pref string, suff string) int {
    
}


/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
