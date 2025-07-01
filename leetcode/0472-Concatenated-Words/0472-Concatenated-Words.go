package leetcode

import "sort"

type TrieNode struct {
	IsEnd    bool
	Children [26]*TrieNode
}

func NewTrieNode() *TrieNode {
	node := &TrieNode{IsEnd: false}

	for i := 0; i < 26; i++ {
		node.Children[i] = nil
	}

	return node
}

func findAllConcatenatedWordsInADict(words []string) []string {
	// sort by length NOT by lexicographical order!!!
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	root := NewTrieNode()

	res := make([]string, 0)

	for _, word := range words {
		if word != "" && checkValid(root, word) {
			res = append(res, word)
		}

		// build word into Trie
		node := root
		for _, ch := range word {
			chIdx := int(ch - 'a')
			if node.Children[chIdx] == nil {
				node.Children[chIdx] = NewTrieNode()
			}
			node = node.Children[chIdx]
		}
		node.IsEnd = true
	}

	return res
}

// How to design Memo in dfs?
// Original dfs(cur int, word string) has ONLY two args and word doesn't vary.
// So we design memo for cur, meaning that a specific index cur can be seen multiple times
// but the first time seen already tells us that if word[cur:] can be "word break" or not,
// we can record the first-time result in memo for later recurrence, and early return if encounter again
func checkValid(root *TrieNode, word string) bool {
	// memoization
	visited := make(map[int]bool)

	// whether word[cur:] can be "word break"
	// X X X X [y y y]
	// NOTE: every recursion call, start from root again bc we are to see if word[cur:] has a prefix in Trie
	var dfs func(cur int, word string, visited map[int]bool) bool

	dfs = func(cur int, word string, visited map[int]bool) bool {
		// base case
		if cur == len(word) {
			return true
		}

		// memo
		if _, ok := visited[cur]; ok {
			return visited[cur]
		}

		node := root
		for i := cur; i < len(word); i++ {
			charIdx := int(word[i] - 'a')
			if node.Children[charIdx] != nil { // word[i] exists in Trie
				node = node.Children[charIdx]

				// node.IsEnd = true means word[cur:i] (inclusive) exists in Trie, cut it!
				// test if word[i+1:] can be "word break"
				if node.IsEnd && dfs(i+1, word, visited) {
					visited[cur] = true
					return true
				}
			} else {
				break
			}
		}

		visited[cur] = false

		return false
	}

	return dfs(0, word, visited)

}
