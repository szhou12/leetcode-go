package leetcode

type TrieNode struct {
	index int
	children [27]*TrieNode // 26 letters + '{'
}

func newTrieNode() *TrieNode {
	node := TrieNode{index: -1}
	for i := 0; i < 27; i++ {
		node.children[i] = nil
	}
	return &node
}

// Add a new word to Trie
func insert(root *TrieNode, word string, id int) {
	node := root
	for _, char := range word {
		letter := int(char - 'a')
		if node.children[letter] == nil {
			node.children[letter] = newTrieNode()
		}
		node = node.children[letter]

		node.index = id
	}
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */

type WordFilter struct {
	root *TrieNode
}

func Constructor(words []string) WordFilter {
    trie := WordFilter{root: newTrieNode()}
	for id, word := range words {
		for i := 0; i < len(word); i++ {
			newWord := word[i:] + "{" + word
			insert(trie.root, newWord, id)
		}
		insert(trie.root, "{" + word, id)
	}
	return trie
}


func (this *WordFilter) F(pref string, suff string) int {
    node := this.root

	concat := suff + "{" + pref
	for _, char := range concat {
		letter := int(char - 'a')
		if node.children[letter] == nil {
			return -1
		}
		node = node.children[letter]
	}

	return node.index
}



