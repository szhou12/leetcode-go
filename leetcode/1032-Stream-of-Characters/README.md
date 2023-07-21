# [1032. Stream of Characters](https://leetcode.com/problems/stream-of-characters/description/)

## Solution idea
### Trie
#### 思路总结
1. 理解题意: 每次`query`，要把当前的 字符 添加在之前所有`query`组成的字符流的末尾
2. 一开始的想法: 建Trie时每个dictionary word**按照顺序插入**。查询时，遍历所有以当前 `query`带来的字符为结尾的所有子串，挨个检查一遍。但是这样做**会超时**，因为 Time Complexity = $O(1 + 2 + 3 + \cdots + Q) = O(Q^2)$ where $Q$ is the number of queries.
```
// e.g. stream queries: a -> b -> c -> d
a b c d
      -
    ---
  -----
-------
```
3. 优化解法: 
    * 题目要求查询 suffix: 暗示 建Trie时每个dictionary word应该**按照逆序插入**。
    * `query`带来的字符流 也按照 逆序组成 字符串: 当前`query`带来的字符放在头部。`d c b a`
    * 爬 Trie 时，走到一个节点的 `isEnd == true` 就算找到了。
    * 这样做只用走一遍Queries。Time Complexity = $O(Q)$ where $Q$ is the number of queries.

Time complexity = $O(mN) + O(Q)$ (build Trie + query) where $m = $ # of dictionary words, $N = $ worst-case length of a dictionary word, $Q = $ # of queries.

## Resource
* [小姐姐刷题-Leetcode 1032 字符流 Trie](https://www.bilibili.com/video/BV1eg411P7ej/?spm_id_from=333.337.search-card.all.click&vd_source=0c02ef6f6e7a2b0959d7dd28e9e49da4)
* [花花酱 LeetCode Weekly Contest 133 (1029,1030,1031,1032)](https://www.youtube.com/watch?v=3A98vh5zsqw&ab_channel=HuaHua)