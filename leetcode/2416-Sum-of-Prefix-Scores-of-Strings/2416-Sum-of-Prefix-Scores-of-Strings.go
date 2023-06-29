package leetcode

func sumPrefixScores(words []string) []int {
	// Step 1: build Trie tree
	root := newTrieNode()
	for _, word := range words {
		node := root
		for _, ch := range word {
			letter := int(ch - 'a')
			if node.children[letter] == nil {
				node.children[letter] = newTrieNode()
			}
			node = node.children[letter]
			node.count++
		}
	}

	// Step 2: traverse Trie tree + collect scores
	res := make([]int, 0)
	for _, word := range words {
		node := root
		score := 0
		for _, ch := range word {
			letter := int(ch - 'a')
			if node.children[letter] == nil {
				break
			}
			node = node.children[letter]
			score += node.count
		}
		res = append(res, score)
	}

	return res
}

type TrieNode struct {
	count int
	children [26]*TrieNode
}

func newTrieNode() *TrieNode {
	node := TrieNode{count: 0}
	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}
	return &node
}