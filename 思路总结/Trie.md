# Trie

## 目录
* [Trie模版](#模版)
* [经典题](#经典题)
* [Design题](#design题)

## 模版
* 字典树介绍: [数据结构与算法： 字典树（Trie）](https://aimuke.github.io/algorithm/2019/07/01/algorithm-trie/)
* 可参考: [208. Implement Trie (Prefix Tree)](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0208-Implement-Trie-(Prefix-Tree))
```go
type TrieNode struct {
    isEnd bool // Whether the current TrieNode denotes the end of a word
    char byte // [Optional] letter stored in the current TrieNode
    count int // [Optional] Count # of words in the text that ends with the current TrieNode's letter
    children [26]*TrieNode
}

func newTrieNode(char byte) *TrieNode {
    node := TrieNode{
        isEnd: false,
        char: char,
        count: 0
    }

    for i := 0; i < 26; i++ {
        node.children[i] = nil
    }
    return &node
}

func buildTrie(root *TrieNode, words []string) {
    for _, word := range words {

        node := root // Start over from root for each word

        for _, char := range word {
            index := int(char - 'a')
            if node.children[index] == nil {
                node.children[index] = newTrieNode(byte(char))
            }
            node = node.children[index]
            node.count++
        }
        node.isEnd = true
    } 
}
```

* Resource: [Implementing A Trie Data structure In Golang(Search engine)](https://medium.com/@itachisasuke/implementing-a-search-engine-in-golang-trie-data-structure-c45152ddda24)



## 经典题
* :red_circle: 共享前缀的总数量: [2416. Sum of Prefix Scores of Strings](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2416-Sum-of-Prefix-Scores-of-Strings)
    * 比较直接的运用**Trie**的题目
    * `count`: 文本词频统计

* :red_circle: input中可以由其他单词做前缀组成的最长单词: [720. Longest Word in Dictionary](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0720-Longest-Word-in-Dictionary)
    * Trie + DFS

* :red_circle: 自动补充的提示词: [1268. Search Suggestions System](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1268-Search-Suggestions-System)
    * Trie + DFS
    * 思路类似 **720**

* :red_circle: 矩阵中搜词II: [212. Word Search II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0212-Word-Search-II)
    * Trie + 矩阵 DFS

* :red_circle: 分词II: [140. Word Break II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0140-Word-Break-II)
    * Trie + 区间型 DFS

* :yellow_circle: 用词根替换: [648. Replace Words](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0648-Replace-Words)
    * Trie


## Design题
* :red_circle: 实现Trie: [208. Implement Trie (Prefix Tree)]()

* :red_circle: 前后缀搜索: [745. Prefix and Suffix Search]()

* :red_circle: :secret: 查询字符流后缀: [1032. Stream of Characters]()
    * 看到是 查询suffix: 想到 逆序建Trie


* :yellow_circle: 实现Trie, 实现添加和查询功能(包含通用符): [211. Design Add and Search Words Data Structure]()
    * Search feature: Recursion

* :yellow_circle: 搜索一字母之差的字典: [676. Implement Magic Dictionary]()
