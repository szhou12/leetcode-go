# [130. Surrounded Regions](https://leetcode.com/problems/surrounded-regions/)

## Solution idea

**Key 1**: 找到所有贴着边界的 O区域, 转换成 A区域; 再把剩下的 O区域转换成 X区域, A区域恢复成 O区域

### DFS / BFS
用DFS或者BFS找贴着边界的 O区域

### Union Find
1. 用 Union Find 先联合所有的连通图。
2. 然后找出包含边界元素的连通图。
3. 最后把不包含边界元素的连通图的元素转换成 X。


Time complexity = O(row * col)

## Resource

有很好的图示：

[代码随想录-130. 被围绕的区域](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0130.%E8%A2%AB%E5%9B%B4%E7%BB%95%E7%9A%84%E5%8C%BA%E5%9F%9F.md)
