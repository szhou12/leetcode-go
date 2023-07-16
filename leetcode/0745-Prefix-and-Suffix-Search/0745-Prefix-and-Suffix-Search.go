package leetcode

type TrieNode struct {
	isEnd bool
	children [27]*TrieNode // 26 letters + '{'
}

func newTrieNode() *TrieNode {
	node := TrieNode{isEnd: false}
	for i := 0; i < 27; i++ {
		node.children[i] = nil
	}
	return &node
}

type WordFilter struct {

}

func Constructor(words []string) WordFilter {
    
}


func (this *WordFilter) F(pref string, suff string) int {
    
}


/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
