# [990. Satisfiability of Equality Equations](https://leetcode.com/problems/satisfiability-of-equality-equations/description/)

## Solution idea
### Union Find
#### 思路总结
1. 思路比较直接的题目。先把过一遍 equations 中的 `a==b`，Union起来，然后再过一遍 equations 中的 `a!=b`，如果遇到相同的祖宗，说明找到一组矛盾的等式。

Time complexity = $O(n)$ where $n = $ length of equations