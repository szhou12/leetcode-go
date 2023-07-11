# [648. Replace Words](https://leetcode.com/problems/replace-words/description/)

## Solution idea
### Trie
#### 思路总结
1. 参考答案的思路是 **Trie + DFS**。但是我的解法不用到 DFS 也能过。
2. 我的思路: 
    1. 构建 Trie
    2. 把 sentence 拆分成一个一个word，每个word爬一次Trie树，遇到一个完整的“词根”就直接返回，否则返回空串。爬树的实现思路类似于 [140. Word Break II](https://leetcode.com/problems/word-break-ii/description/)

Time complexity = $O(m * H + n * H)$ where $m = $ number of words in the dictionary, $n = $ number of words in the sentence, $H = $ height of the Trie tree.

## Resource
[wisdompeak/LeetCode/Trie/648.Replace-Words/](https://github.com/wisdompeak/LeetCode/tree/master/Trie/648.Replace-Words)