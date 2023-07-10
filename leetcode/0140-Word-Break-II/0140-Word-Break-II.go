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
	// base case
	if startIndex == len(s) {
		sentence = sentence[:len(sentence) - 1]
		*res = append(*res, sentence)
		return
	}

	node := root
	for end := startIndex; end < len(s); end++ {
		letter := int(s[end] - 'a')
		if node.children[letter] != nil {
			node = node.children[letter]
			if node.isEnd {
				prevLen := len(sentence)
				sentence += s[startIndex:end+1] + " "
				dfs(s, root, end+1, sentence, res)
				sentence = sentence[:prevLen]
			}
		} else {
			break
		}
	}
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