package leetcode

import "sort"

// WARNING: Trie will TLE too!
func palindromePairs_trie(words []string) [][]int {
	n := len(words)
	dict := make(map[string]int)
	for i, word := range words {
		dict[word] = i
	}

	// sort by length WHY?
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	res := make([][]int, 0)
	root := newTrieNode()

	for i := 0; i < n; i++ {
		suffix := words[i]
		for k := 0; k <= len(suffix); k++ {
			s1 := suffix[:k]
			s2 := suffix[k:]
			// s2's1s2
			match := reverse(s2)
			if root.search(match) && isPalidrome(s1) {
				res = append(res, []int{dict[match], dict[suffix]})

			}

			// s1s2s1'
			match = reverse(s1)
			if root.search(match) && isPalidrome(s2) {
				res = append(res, []int{dict[suffix], dict[match]})
			}
		}
		root.insert(suffix)
	}
	return res
}

type TrieNode struct {
	isEnd    bool
	children [26]*TrieNode
}

func newTrieNode() *TrieNode {
	node := TrieNode{isEnd: false}
	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}
	return &node
}

func (node *TrieNode) insert(word string) {

	for _, ch := range word {
		letter := int(ch - 'a')
		if node.children[letter] == nil {
			node.children[letter] = newTrieNode()
		}
		node = node.children[letter]
	}
	node.isEnd = true
}

func (node *TrieNode) search(word string) bool {
	for _, ch := range word {
		letter := int(ch - 'a')
		if node.children[letter] == nil {
			return false
		}
		node = node.children[letter]
	}
	return node.isEnd
}
