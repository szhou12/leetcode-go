# [200. Number of Islands](https://leetcode.com/problems/number-of-islands/submissions/)

## Solution idea
### DFS


### BFS


### Union Find
#### 思路总结
1. 遍历每一个节点，如果节点=1，看它上下左右四个方向的邻居是否=1，如果是，就Union起来
2. 技巧: 1-D index compression

Time Complexity = O(rows*cols)