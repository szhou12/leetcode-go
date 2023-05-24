# [1584. Min Cost to Connect All Points](https://leetcode.com/problems/min-cost-to-connect-all-points/description/)

## Solution idea
### MST = Union Find + Kruskal
#### 思路总结
1. 比较直观的关于MST的题目：算出每条边的weight后，使用Kruskal算法。
2. 小技巧：每个节点是二维坐标，UnionFind怎么储存节点信息？使用每个节点的index，而不是像矩阵里那样把二维坐标压缩为一维。

Time complexity = $O(n^2) + O(n^2\log n^2) + O(n^2) = O(n^2\log n^2)$

## Resource
[【每日一题】1584. Min Cost to Connect All Points, 9/13/2020](https://www.youtube.com/watch?v=ZVh2WTcE8EY&ab_channel=HuifengGuan)