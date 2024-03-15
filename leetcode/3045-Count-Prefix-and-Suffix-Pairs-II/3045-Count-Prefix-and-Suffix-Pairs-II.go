package leetcode

// 1. pair (i, j) where i < j => loop backwards: for each i, find # of j's in Trie
// 2. KMP Step 1 find longest, 2nd longest, ... suffix-prefix for each word
// 3. Trie stores all lsp for each word
func countPrefixSuffixPairs(words []string) int64 {
	root := newTrieNode()

	res := 0
	for i := len(words) - 1; i >= 0; i-- {
		// step 1: find # of longest suffix-prefix that match current word
		res += find(root, words[i]) // O(L)

		// step 2: find all longest suffix-prefix for current word
		n := len(words[i])
		lspLen := make(map[int]bool) // store len of all longest suffix-prefix

		lsp := preprocess(words[i]) // O(L)
		l := lsp[n-1] // len of longest suffix-prefix
		for l > 0 {
			lspLen[l] = true
			l = lsp[l-1] // len of next longest suffix-prefix
		}
		lspLen[n] = true // current word itself can be a longest suffix-prefix in this problem bc we are matching another word with cur word rather than self-match

		// step 3: add current word's all longest suffix-prefix to Trie
		add(root, words[i], lspLen) // O(L)
	}

	return int64(res)
}

/* add all longest suffix-prefix of word to Trie */
func add(root *TrieNode, word string, set map[int]bool) {
	node := root

	for i, ch := range word {
		chIdx := int(ch-'a')
		if node.children[chIdx] == nil {
			node.children[chIdx] = newTrieNode()
		}
		node = node.children[chIdx]
		// i+1 = len of word[0:i] (inclusive)
		if _, ok := set[i+1]; ok {
			node.count++
		}
	}
}

/* find matches in Trie */
func find(root *TrieNode, word string) int {
	node := root

	for _, ch := range word {
		chIdx := int(ch-'a')
		if node.children[chIdx] == nil {
			return 0
		}
		node = node.children[chIdx]
	}

	return node.count
}

/* KMP Step 1 */
// dp[i] := max length j s.t. needle[0:j-1] == needle[i-j+1:i] (inclusive)
func preprocess(needle string) []int {
	n := len(needle)
	dp := make([]int, n)

	// base case: require 'true' prefix and suffix
	dp[0] = 0

	// recurrence
	for i := 1; i < n; i++ {
		j := dp[i-1]
		for j > 0 && needle[j] != needle[i] {
			j = dp[j-1]
		}
		if needle[j] == needle[i] {
			dp[i] = j+1
		} else {
			dp[i] = j
		}
	}
	return dp
}

/* Trie Structure */
type TrieNode struct {
	count int // # of words ending at this node
	children [26]*TrieNode
}

func newTrieNode() *TrieNode {
	node := TrieNode{count:0}
	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}
	return &node
}