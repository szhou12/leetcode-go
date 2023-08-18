# [2812. Find the Safest Path in a Grid](https://leetcode.com/problems/find-the-safest-path-in-a-grid/)

## Solution idea
### BFS + Binary Search 猜答案
#### 思路总结
#### 物理意义
重新定义: `grid[i][j] =` the distance from cell (i, j) to nearest thief + 1
#### 代码结构
```
Step 1:
    - BFS: 从每个thief格子出发，计算每个格子到最近的thief的距离
Step 2:
    - Binary Searh: 猜测一个最大的安全距离 (max safe factor)
        - BFS: 检查在当前给定的安全距离下，是否有一条从起点到终点的路径
```

#### Time complexity & Space complexity

## Resource
[【每日一题】LeetCode 2812. Find the Safest Path in a Grid](https://www.youtube.com/watch?v=OjzQ6TmMh6k&ab_channel=HuifengGuan)