# Trie

## 目录
* [Trie模版](#模版)
* [经典题](#经典题)

## 模版
```go
type TrieNode struct {
    char byte
    children [26]*TrieNode
}

func newTrieNode(char byte) *TrieNode {
    node := TrieNode(char: char)
    for i := 0; i < 26; i++ {
        node.children[i] = nil
    }
    return &node
}
```

* [Implementing A Trie Data structure In Golang(Search engine)](https://medium.com/@itachisasuke/implementing-a-search-engine-in-golang-trie-data-structure-c45152ddda24)

## 经典题
* :red_circle: 共享前缀的总数量: [2416. Sum of Prefix Scores of Strings](https://leetcode.com/problems/sum-of-prefix-scores-of-strings/description/)
    * 比较直接的运用**Trie**的题目

* :red_circle: input中可以由其他单词做前缀组成的最长单词: [720. Longest Word in Dictionary](https://leetcode.com/problems/longest-word-in-dictionary/description/)
    * Trie + DFS