# [1061. Lexicographically Smallest Equivalent String](https://leetcode.com/problems/lexicographically-smallest-equivalent-string/description/)

## Solution idea
### Union Find
#### 要点总结
1. 思路比较直接的题目。题目已经明确给出 transitive 的性质，说明使用 Union Find 来解。字典序最小的显然就是每个 字符 所属component里的root。

Time complexity = $O(n)$ where `n = max{s1.length, s2.length, baseStr.length}`
