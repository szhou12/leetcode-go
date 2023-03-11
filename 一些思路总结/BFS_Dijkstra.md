# Breadth First Search

## Tree - Level Order Traversal

### 最经典最基础的 BFS 应用

* 基础: [102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)

* 进阶: [103. Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/)

* 基础: [116. Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)

* 基础: [513. Find Bottom Left Tree Value](https://leetcode.com/problems/find-bottom-left-tree-value/)

### 岛屿沉没类 - 找neighbor
* Roblox - Candy Crush

* [200. Number of Islands](https://leetcode.com/problems/number-of-islands/submissions/)

* 四面八方, 潮水涌来 [417. Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)

* [130. Surrounded Regions](https://leetcode.com/problems/surrounded-regions/)

* [1020. Number of Enclaves](https://leetcode.com/problems/number-of-enclaves/)

### Dijkstra - find shortest path

* 适用题目的特征:
    1. **single-source**: 只给一个起点, 求从起点(single source)到图上任意一个node的最短路径
    2. **non-negative weight**: edge的权重非负
    * 这样, 就可以求图上任意一点到起点single-source的最短距离
    * Time complexity = $O(E\log E)$

* 从Grid左上角走到右下角所花最短时间: [2577. Minimum Time to Visit a Cell In a Grid](https://leetcode.com/problems/minimum-time-to-visit-a-cell-in-a-grid/description/)

* 从两个起点到达同一个终点所需最短路径: [2203. Minimum Weighted Subgraph With the Required Paths](https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths/description/)

* 到达终点的路径总数: [1976. Number of Ways to Arrive at Destination](https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/description/)
    * **Dijkstra + DFS + Deduplication**
    * 1976 与 1786 两题逻辑完全相同

* 从起点走到终点的特殊路径总数: [1786. Number of Restricted Paths From First to Last Node](https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/description/)
    * **Dijkstra + DFS + Deduplication**
    * 1976 与 1786 两题逻辑完全相同

