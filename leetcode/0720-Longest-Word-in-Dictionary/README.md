# [720. Longest Word in Dictionary](https://leetcode.com/problems/longest-word-in-dictionary/description/)

## Solution idea
### Trie
#### 思路总结
1. 构造Trie树，然后使用 DFS 找到最深处就是最长的单词。注意：用 `isEnd` 标注一个TrieNode是否是一个单词的结尾。只有 `isEnd` 为 `true` 的节点才能继续往深处走。

## Resource
[wisdompeak/BreadcrumbsLeetCode/Trie/1858.Longest-Word-With-All-Prefixes](https://github.com/wisdompeak/LeetCode/tree/master/Trie/1858.Longest-Word-With-All-Prefixes)

## NOTE
Same question as [1858 - Longest Word With All Prefixes](https://leetcode.ca/2021-07-14-1858-Longest-Word-With-All-Prefixes/)