# [211. Design Add and Search Words Data Structure](https://leetcode.com/problems/design-add-and-search-words-data-structure/description/)

## Solution idea
### Trie
#### 思路总结
1. Trie的设计题: 比较典型的Trie题目。`bool search(word)` 我采用在树上 Top-down Recursion 的方式搜索， 每一次层只看头字母，剩下的交给recursive call

Time complexity = $O(n)$ where $n$ is the length of the word (`void addWord(word)`)

Time complexity = $O(26 * k)$ where $k$ is the height of input searching word (`bool search(word)`)

## Resource
[wisdompeak/LeetCode/Trie/211.Add-and-Search-Word/](https://github.com/wisdompeak/LeetCode/tree/master/Trie/211.Add-and-Search-Word)