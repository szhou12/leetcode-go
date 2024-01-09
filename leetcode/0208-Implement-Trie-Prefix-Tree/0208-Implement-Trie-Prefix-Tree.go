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
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
    root *TrieNode
}


func Constructor() Trie {
    trie := Trie{root: newTrieNode()}
    return trie
}


func (this *Trie) Insert(word string)  {
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


func (this *Trie) Search(word string) bool {
    node := this.root
    for _, char := range word {
        letter := int(char - 'a')
        if node.children[letter] == nil {
            return false
        }
        node = node.children[letter]
    }
    return node.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
    node := this.root
    for _, char := range prefix {
        letter := int(char - 'a')
        if node.children[letter] == nil {
            return false
        }
        node = node.children[letter]
    }
    return true
}