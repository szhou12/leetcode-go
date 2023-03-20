# [1368. Minimum Cost to Make at Least One Valid Path in a Grid](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/description/)

## Solution idea

### Dijkstra

#### 要点总结
1. 怎么把题目抽象成Dijkstra求最短路径?
    * 关键: 在于把`grid[i][j]`里指向的方向看做权重=0, 没指向的方向权重=1. e.g. `grid[i][j]=1`意味着向右走cost=0, 向上走cost=1, 向下走cost=1, 向左走cost=1
2. 此题是 **矩阵走格子类型**, 所以, 走下一步时除了检查visited, 还要检查out-of-bound

Time complexity = $O(mn\log mn)$

## Resource
[【每日一题】1368. Minimum Cost to Make at Least One Valid Path in a Grid, 2/20/2021](https://www.youtube.com/watch?v=Xctb7zzF1I4&ab_channel=HuifengGuan)