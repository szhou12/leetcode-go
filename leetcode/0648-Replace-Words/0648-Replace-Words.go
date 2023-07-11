package leetcode

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	root := newTrieNode()
	// Step 1: build Trie from dictionary
	for _, word := range dictionary { // O(m * Height)
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

	// Step 2: search word root and replace if exists
	splits := strings.Split(sentence, " ")
	res := ""
	for _, word := range splits { // O(n * Height)
		wordRoot := searchRoot(root, word)
		if wordRoot == "" {
			res += word + " "
		} else {
			res += wordRoot + " "
		}
	}
	return res[:len(res) - 1]
}

// Test case 1: no root in Trie
// Test case 2: partial letters in Trie but not a complete root
// e.g. Trie = {bat}. case 1: "cow", case 2: "bottle", case 3: "battle"
func searchRoot(root *TrieNode, word string) string {
	wordRoot := ""

	node := root
	for _, char := range word {
		letter := int(char - 'a')
		if node.children[letter] != nil {
			node = node.children[letter]
			wordRoot += string(char)
			if node.isEnd {
				return wordRoot
			}
		} else {
			break
		}
	}

	return ""
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
