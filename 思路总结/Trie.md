# Trie

## 目录
* [Trie模版](#模版)
* [经典题](#经典题)

## 模版
* 可参考 [208. Implement Trie (Prefix Tree)](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0208-Implement-Trie-(Prefix-Tree))
```go
type TrieNode struct {
    char byte // Optional: letter stored in the current TrieNode
    isEnd bool // Optional: whether the current TrieNode denotes the end of a word
    children [26]*TrieNode
}

func newTrieNode(char byte) *TrieNode {
    node := TrieNode{
        char: char,
        isEnd: false,
    }

    for i := 0; i < 26; i++ {
        node.children[i] = nil
    }
    return &node
}

func buildTrie(words []string) *TrieNode {
    root := newTrieNode(' ')
    for _, word := range words {
        node := root

        for _, char := range word {
            index := int(char - 'a')
            if node.children[index] == nil {
                node.children[index] = newTrieNode(byte(char))
            }
            node = node.children[index]
        }
        node.isEnd = true
    }
    return root
}
```

* Resource: [Implementing A Trie Data structure In Golang(Search engine)](https://medium.com/@itachisasuke/implementing-a-search-engine-in-golang-trie-data-structure-c45152ddda24)

## 经典题
* :red_circle: 共享前缀的总数量: [2416. Sum of Prefix Scores of Strings](https://leetcode.com/problems/sum-of-prefix-scores-of-strings/description/)
    * 比较直接的运用**Trie**的题目

* :red_circle: input中可以由其他单词做前缀组成的最长单词: [720. Longest Word in Dictionary](https://leetcode.com/problems/longest-word-in-dictionary/description/)
    * Trie + DFS

* :red_circle: 自动补充的提示词: [1268. Search Suggestions System](https://leetcode.com/problems/search-suggestions-system/)
    * Trie + DFS
    * 思路类似 **720**

* :red_circle: 矩阵中搜词II: [212. Word Search II](https://leetcode.com/problems/word-search-ii/description/)
    * Trie + DFS in matrix

* :red_circle: 分词II: [140. Word Break II](https://leetcode.com/problems/word-break-ii/description/)
    * Trie + 区间型 DFS