package leetcode

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

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

type MagicDictionary struct {
	root *TrieNode
}

func Constructor() MagicDictionary {
	trie := MagicDictionary{root: newTrieNode()}
	return trie
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {

		node := this.root // MUST start from root again for each word

		for _, char := range word {
			letter := int(char - 'a')
			if node.children[letter] == nil {
				node.children[letter] = newTrieNode()
			}
			node = node.children[letter]
		}
		node.isEnd = true
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {

	for i, char := range searchWord {
		for j := 0; j < 26; j++ {
			newWord := searchWord[:i] + string(j+'a') + searchWord[i+1:]

			// NOTE: MUST change one letter
			if char == rune(j+'a') {
				continue
			}

			if canFind(this.root, newWord) {
				return true
			}
		}
	}

	return false
}

func canFind(root *TrieNode, word string) bool {
	node := root
	for _, char := range word {
		letter := int(char - 'a')
		if node.children[letter] == nil {
			return false
		}
		node = node.children[letter]
	}

	return node.isEnd
}


