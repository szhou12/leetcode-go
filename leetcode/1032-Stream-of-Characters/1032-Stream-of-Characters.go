package leetcode

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

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */

type StreamChecker struct {
    root *TrieNode
	stream string
}


func Constructor(words []string) StreamChecker {
    trie := StreamChecker{
		root: newTrieNode(),
		stream: "",
	}

	// Built Trie on reversed words
	for _, word := range words {
		node := trie.root
		revWord := reverse(word)

		for _, char := range revWord {
			letter := int(char - 'a')
			if node.children[letter] == nil {
				node.children[letter] = newTrieNode()
			}
			node = node.children[letter]
		}
		node.isEnd = true
	}

	return trie
}

func reverse(word string) string {
	res := ""
	// prepend each char
	for _, char := range word {
		res = string(char) + res
	}
	return res
}


func (this *StreamChecker) Query(letter byte) bool {
    // add current letter to this.stream to form reversed order
	this.stream = string(letter) + this.stream

	node := this.root
	for _, char := range this.stream {
		letter := int(char - 'a')
		if node.children[letter] == nil {
			return false
		}
		node = node.children[letter]
		if node.isEnd {
			return true
		}
	}

	return false

}
