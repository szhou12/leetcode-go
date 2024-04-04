package leetcode

import "math"

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	root := newTrieNode(-1, math.MaxInt)
	// build Trie
	for i, word := range wordsContainer {
		// root node needs to mark (idx, length) that meet criteria as well
		length := len(word)
		// if current word is shorter, update root
		// if equal length, no need to update bc previously stored idx is natually earlier
		if root.length > length {
			root.idx = i
			root.length = length
		}
		buildTrie(root, word, i)
	}

	res := make([]int, 0)
	for _, word := range wordsQuery {
		res = append(res, searchTrie(root, word))
	}
	return res
}

func searchTrie(root *TrieNode, word string) int {
	node := root
	n := len(word)
	for i := n-1; i>=0; i-- {
		letter := int(word[i] - 'a')
		if node.children[letter] == nil {
			return node.idx
		}
		node = node.children[letter]
	}
	return node.idx
}

// Add one word to Trie
func buildTrie(root *TrieNode, word string, idx int) {
	node := root
	n := len(word)
	// iterate characters reversely
	for i := n-1; i>=0; i-- {
		letter := int(word[i] - 'a')
		if node.children[letter] == nil {
			node.children[letter] = newTrieNode(idx, n)
		}
		node = node.children[letter]

		// if current word is shorter than stored, update node's info
		// if equal length, no need to update bc previously stored idx is natually earlier
		if node.length > n {
			node.idx = idx
			node.length = n
		}
	}
	node.isEnd = true
}

type TrieNode struct {
	isEnd    bool // not necessary in this problem
	idx      int  // index of the shortest word whose suffix reaches this node; earliest index if equal legnth
	length   int  // length of the shortest word whose suffx reaches this node; or length of earliest word if equal length
	children [26]*TrieNode
}

func newTrieNode(idx int, length int) *TrieNode {
	node := TrieNode{
		idx:    idx,
		length: length,
	}
	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}
	return &node
}
