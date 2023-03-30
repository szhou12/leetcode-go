# Breadth First Search & Dijkstra

## 最经典最基础的 BFS 应用: Tree - Level Order Traversal

* 基础: [102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)

* 进阶: [103. Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/)

* 基础: [116. Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)

* 基础: [513. Find Bottom Left Tree Value](https://leetcode.com/problems/find-bottom-left-tree-value/)

## 岛屿沉没类 - 找neighbor
* Roblox - Candy Crush

* [200. Number of Islands](https://leetcode.com/problems/number-of-islands/submissions/)

* 四面八方, 潮水涌来 [417. Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)

* [130. Surrounded Regions](https://leetcode.com/problems/surrounded-regions/)

* [1020. Number of Enclaves](https://leetcode.com/problems/number-of-enclaves/)

## Dijkstra - Find Shortest Path

* :star: **适用题目的特征:**
    1. **single-source**: 只给一个起点, 求从起点(single source)到图上任意一个node的最短路径
    2. **non-negative weight**: edge的权重非负
    * 这样, 就可以求图上任意一点到起点single-source的最短距离
    * Time complexity = $O(E\log E)$

* :red_circle: 从Grid左上角走到右下角所花最短时间: [2577. Minimum Time to Visit a Cell In a Grid](https://leetcode.com/problems/minimum-time-to-visit-a-cell-in-a-grid/description/)
    * **矩阵走格子类型题**

* :red_circle: 从Grid左上角走到右下角所花最少cost: [1368. Minimum Cost to Make at Least One Valid Path in a Grid](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/description/)
    * **矩阵走格子类型题**

* 从两个起点到达同一个终点所需最短路径: [2203. Minimum Weighted Subgraph With the Required Paths](https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths/description/)

* :red_circle: :secret: 到达终点的路径总数: [1976. Number of Ways to Arrive at Destination](https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/description/)
    * **Dijkstra + DFS + Deduplication**
    * 1976 与 1786 两题逻辑完全相同
    * 题目求方法总数, 最自然的是使用DFS
    * 使用Memo避免DFS重复走走过的路径

* :red_circle: :secret: 从起点走到终点的特殊路径总数: [1786. Number of Restricted Paths From First to Last Node](https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/description/)
    * **Dijkstra + DFS + Deduplication**
    * 1976 与 1786 两题逻辑完全相同
    * 题目求方法总数, 最自然的是使用DFS
    * 使用Memo避免DFS重复走走过的路径

* 从起点可到达的nodes总数: [882. Reachable Nodes In Subdivided Graph](https://leetcode.com/problems/reachable-nodes-in-subdivided-graph/)

* :red_circle: :secret: K 站中转内最便宜的航班路径: [787. Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/description/)
    * **node储存双状态**: node 储存 位置信息+已中转次数信息
    * Early Return: Dijkstra loop内return

* :green_circle: 最高概率的路径: [1514. Path with Maximum Probability](https://leetcode.com/problems/path-with-maximum-probability/)

* :green_circle: 从source node到所有node所需最短时间: [743. Network Delay Time](https://leetcode.com/problems/network-delay-time/description/)
    * Dijkstra求最短路径的基础题: 求从source node到任意node i的最短路径
    * 可用来作为以上题目的练手题

* :lock: :secret: 走迷宫II / 走冰系道馆: [505. The Maze II](https://leetcode.ca/all/505.html)
    * 难点: 地面是冰面, 只能维持一个方向前进直到撞墙

## BFS + PQ: 思路想BFS, 代码结构像Dijkstra
* 一个query可以淹没多少格子, 求淹没总数: [2503. Maximum Number of Points From Grid Queries](https://leetcode.com/problems/maximum-number-of-points-from-grid-queries/)

* 升高海平面中的最短路线: [778. Swim in Rising Water](https://leetcode.com/problems/swim-in-rising-water/description/)

* 接雨水II (海水倒灌II): [407. Trapping Rain Water II](https://leetcode.com/problems/trapping-rain-water-ii/description/)
    * 题目设置很像 [417. Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)
