# [2385. Amount of Time for Binary Tree to Be Infected](https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/description/)

## Solution idea
### BFS
1. Use post-order traversal to convert tree to adj-list representation
2. Conduct BFS to calculate the shortest time to reach all nodes. **NOTE:** tree structure is undirected graph, so we need to add edge in both directions (parent -> child and child -> parent).

Time complexity = $O(n)$

Space complexity = $O(n)$. When converting the tree to a graph, we require $O(n)$ extra space for the map. We also require $O(n)$ space for the queue and $O(n)$ space for the visited set during the BFS.