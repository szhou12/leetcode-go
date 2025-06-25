# [3593. Minimum Increments to Equalize Leaf Paths](https://leetcode.com/problems/minimum-increments-to-equalize-leaf-paths/description/)

## Solution idea
### DFS on undirected tree with edges represented as adj-list.
2. Intuition: 为了最小化increment的操作次数，我们会将所有root-to-leaf的路径与Cost最大的那条路径对齐。

Time complexity = $O(n)$

Space complexity = $O(n)$