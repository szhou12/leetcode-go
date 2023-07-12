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
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

type WordDictionary struct {
    root *TrieNode
}


func Constructor() WordDictionary {
    trie := WordDictionary{root: newTrieNode()}
	return trie
}


func (this *WordDictionary) AddWord(word string)  {
    node := this.root
	for _, char := range word {
		letter := int(char - 'a')
		if node.children[letter] == nil {
			node.children[letter] = newTrieNode()
		}
		node = node.children[letter]
	}
	node.isEnd = true
}


func (this *WordDictionary) Search(word string) bool {
    return canFind(this.root, word)
}

func canFind(node *TrieNode, word string) bool {
	// base case
	if len(word) == 0 {
		return node.isEnd
	}

	found := false
	if word[0] != '.' {
		letter := int(word[0] - 'a')
		if node.children[letter] != nil {
			res := canFind(node.children[letter], word[1:])
			found = found || res
		}
	} else {
		for i := 0; i < 26; i++ {
			if node.children[i] != nil {
				res := canFind(node.children[i], word[1:])
				found = found || res
			}
		}
	}

	return found
}
