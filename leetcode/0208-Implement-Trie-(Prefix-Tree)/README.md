# [208. Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/description/)

## Solution idea
### Trie
#### 思路总结
1. 设计题: 实现 Trie. 
    * 其实这里的 Trie class 就可以认为是 TrieNode class，不需要额外在 Trie class 里面再定义 TrieNode class
    * 但为了方便理解，定义 TrieNode class 作为 Trie class 的内部类/变量
2. Trie类型的基本功操作包括：
    1. Trie类型的定义：
        * 成员变量: TrieNode
            * 包括两个成员变量: 
                1. `isEnd bool`: 是否当前节点标记了一个单词的结尾
                2. `children [26]*TrieNode`: 26个子节点对应26个字母
    2. Trie的构造函数
    3. 在Trie树中找指定的完整单词（需要找到叶子节点），check if `isEnd == true`
    4. 在Trie树中找指定的前缀（不需要找到叶子节点）

## Resource
[wisdompeak/LeetCode/Trie/208.Implement-Trie--Prefix-Tree/](https://github.com/wisdompeak/LeetCode/tree/master/Trie/208.Implement-Trie--Prefix-Tree)

[halfrost/LeetCode-Go/leetcode/0208.Implement-Trie-Prefix-Tree/](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0208.Implement-Trie-Prefix-Tree)