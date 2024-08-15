# [3249. Count the Number of Good Nodes](https://leetcode.com/problems/count-the-number-of-good-nodes/description/)

## Solution idea
### BFS + DFS
1. BFS: BFS建树。因为只给了`edges [][]int`, 并且无法判断`[a, b]`谁是父节点谁是子节点
2. DFS: DFS统计以当前节点为根时，它下辖的每个分支的节点数

Time complexity = $O(V+E)$