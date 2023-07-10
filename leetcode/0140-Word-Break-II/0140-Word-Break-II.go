package leetcode

func wordBreak(s string, wordDict []string) []string {
	 root := newTrieNode()
	 // Step 1: build Trie from wordDict
	 for _, word := range wordDict {
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

	 // Step 2: DFS
	 var res []string
	 dfs(s, root, 0, "", &res)
	 return res
}

func dfs(s string, root *TrieNode, startIndex int, sentence string, res *[]string) {
	
}

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