# [3093. Longest Common Suffix Queries](https://leetcode.com/problems/longest-common-suffix-queries/description/)

## Solution idea
### Trie
1. Longest common suffix: 单词逆序建Trie树转化成Longest common prefix
2. 每个Trie node存储最符合criteria的单词的信息 (index, length), 更新信息如果新单词更符合criteria。
    - criteria: 单词总长度最短。如果等长，单词的index最小。
3. Guan神的解法是先按照criteria的优先级排序, Trie node储存的信息会少一些。

Time complexity = $O(n\bar L)$ where n is the number of words, $\bar L$ is the average length of words.

## Resource
[【每日一题】LeetCode 3093. Longest Common Suffix Queries](https://www.youtube.com/watch?v=ZEwzjwfIIAQ&ab_channel=HuifengGuan)