# Breadth First Search & Dijkstra

## 目录
 * [最经典最基础的BFS应用: Tree - Level Order Traversal](#最经典最基础的-bfs-应用-tree---level-order-traversal)
 * [岛屿沉没类](#岛屿沉没类---找neighbor)
 * [Dijkstra](#dijkstra---find-shortest-path)
 * [BFS + PQ](#bfs--pq-思路想bfs-代码结构像dijkstra)
 * [Topological Sort](#topological-sort)

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
    * 两种Return方式:
        1. Early Return: Dijkstra loop内到达终点即return; loop结束还没return说明无法到达终点
        2. 填表 Return: Dijkstra把所有结果填入一个表 (array, matrix) 中, loop结束后把所求node的结果按要求从表中取出return

* :red_circle: 从Grid左上角走到右下角所花最短时间: [2577. Minimum Time to Visit a Cell In a Grid](https://leetcode.com/problems/minimum-time-to-visit-a-cell-in-a-grid/description/)
    * **矩阵走格子类型题**

* :red_circle: 从Grid左上角走到右下角所花最少cost: [1368. Minimum Cost to Make at Least One Valid Path in a Grid](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/description/)
    * **矩阵走格子类型题**

* :red_circle: 从两个起点到达同一个终点所需最短路径: [2203. Minimum Weighted Subgraph With the Required Paths](https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths/description/)
    * 难点: "拉链" - 如何找到"拉链"node i

* :red_circle: :secret: 到达终点的路径总数: [1976. Number of Ways to Arrive at Destination](https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/description/)
    * **Dijkstra + DFS + Deduplication**
    * 题目求方法总数, 最自然的是使用DFS
    * 使用Memo避免DFS重复走走过的路径
    * *1976 与 1786 两题逻辑完全相同*

* :red_circle: :secret: 从起点走到终点的特殊路径总数: [1786. Number of Restricted Paths From First to Last Node](https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/description/)
    * **Dijkstra + DFS + Deduplication**
    * 题目求方法总数, 最自然的是使用DFS
    * 使用Memo避免DFS重复走走过的路径
    * *1976 与 1786 两题逻辑完全相同*

* :red_circle: 从起点可到达的nodes总数: [882. Reachable Nodes In Subdivided Graph](https://leetcode.com/problems/reachable-nodes-in-subdivided-graph/)
    * 难点1: 如何把解题思路与Dijkstra做关联 --> 把 小node 看做 edge weight
    * 难点2: Dijkstra expand时添加限制条件

* :red_circle: :secret: K 站中转内最便宜的航班路径: [787. Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/description/)
    * **node储存双状态**: node 储存 位置信息+已中转次数信息
    * Dijkstra结果储存进 二维矩阵 = (# of nodes, # of stops used so far)
    * Early Return: Dijkstra loop内到达终点即return
    * *0787 与 2093 两题逻辑相似*

* :lock: :yellow_circle: :secret: [2093. Minimum Cost to Reach City With Discounts](https://leetcode.ca/2021-12-16-2093-Minimum-Cost-to-Reach-City-With-Discounts/#2093-minimum-cost-to-reach-city-with-discounts)
    * 每次Expand一个邻居node, push "两条边": toll edge; toll/2 edge
    * **node储存双状态**: node 储存 位置信息+已使用discounts次数信息
    * Dijkstra结果储存进 二维矩阵 = (# of nodes, # of discounts used so far)
    * Early Return: Dijkstra loop内到达终点即return
    * *0787 与 2093 两题逻辑相似*

* :green_circle: 最高概率的路径: [1514. Path with Maximum Probability](https://leetcode.com/problems/path-with-maximum-probability/)

* :green_circle: 从source node到所有node所需最短时间: [743. Network Delay Time](https://leetcode.com/problems/network-delay-time/description/)
    * Dijkstra求最短路径的基础题: 求从source node到任意node i的最短路径
    * 可用来作为以上题目的练手题

* :lock: :yellow_circle: :secret: 走迷宫II / 走冰系道馆: [505. The Maze II](https://leetcode.ca/2017-04-18-505-The-Maze-II/)
    * 难点: 地面是冰面, 只能维持一个方向前进直到撞墙
    * Early Return: Dijkstra loop内到达终点即return

* :lock: :red_circle: :secret: 走迷宫III / 走冰系道馆+掉洞: [499. The Maze III](https://leetcode.ca/2017-04-12-499-The-Maze-III/)
    * **node储存双状态**: node 储存 位置信息+到达node的方向指令信息
    * Priority Queue 再定义:
        1. 用一个结构体容纳不同variable type (int, string)
        2. `Less(i, j int)`: 双排序 - 先按照 path cost 排序, 再按照 指令的lexicographical order排序
    * 维持一个方向滑冰时, 如果掉洞，直接break，不再往前滑
    * Early Return: Dijkstra loop内到达终点即return

## BFS + PQ: 思路想BFS, 代码结构像Dijkstra
* 一个query可以淹没多少格子, 求淹没总数: [2503. Maximum Number of Points From Grid Queries](https://leetcode.com/problems/maximum-number-of-points-from-grid-queries/)

* 升高海平面中的最短路线: [778. Swim in Rising Water](https://leetcode.com/problems/swim-in-rising-water/description/)

* 接雨水II (海水倒灌II): [407. Trapping Rain Water II](https://leetcode.com/problems/trapping-rain-water-ii/description/)
    * 题目设置很像 [417. Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)



## Topological Sort

* :star: **解题思路:**
    * **农村包围城市/剥洋葱**: 优先访问那些入度最低的节点。删去第一批最外围的节点后，再继续访问此时入度更新为最低的节点。依次类推。
    * 使用的数据结构: BFS
        * 注: 拓扑排序也可以用DFS实现，但个人感觉BFS的思路更直观更易模版化且算法在空间上也更高效，遂练习时着重使用BFS解题
    * 使用条件: **Acyclic** (在无向图中就意味着必须是一个Tree)
    * 入度/degree 的定义: 一般的, **degree = the number of incoming edges**.
        * 注: degree 的定义也可能根据具体题目的要求进行调整，不一定死板地遵守一般定义
    * 外围 (低阶级) vs 内层 (高阶级): 一般的，最外围的节点 = 入度的值最低的节点
        * 注: 最外围节点的定义也可能根据具体题目的要求增添额外的判定条件，但入度的值一般是其中一条必要的判定条件
    * 有向图 vs 无向图:
        * 最低入度值 (**最低入度值是非常重要的判定外围node的条件！**):
            * **有向图**中node入度最低 = 0
            * **无向图**中node入度最低 = 1
        * Check for visited (防止从内层重新走回外围的机会):
            * **有向图**无需 check for visited, 因为有向就意味着从内层走回外围的机会根本不可能发生。
            * **无向图**需要 check for visited, 除非next move时直接过河拆桥, 把从内层走回外围的edge删除了 (e.g. [2603. Collect Coins in a Tree](https://leetcode.com/problems/collect-coins-in-a-tree/description/))
        * visited 另一作用 - 给非环节点染色:
            * **有向图**中visited一般用于检测环, visited起到"染色"的作用 - 即，visited==1说明节点不属于环的一部分，visited==0说明节点属于一个环。因为，能被访问说明节点入度可变为0，不能访问说明节点入度不可变为0，也就是，它在一个环中。
            * 注：染色作用的 visited 实际上只是命名为 visited，其他名字也可以，visited只是比较方便理解
    * Adjacency List `next` 的 data structure 选择:
        1. `[][]int`: slice of slices
        2. `[]map[int]bool`: slice of maps
        3. `map[string]map[string]bool`: map of maps (json)
        * 总结: 怎么方便怎么来。但是，要注意！选择 map 相关的结构时 (第2, 3种)，要额外考虑 duplicated edges 的情况，重复的edge要跳过，因为，`next`不会重复添加，但是`degree`会不小心多+1。选择 slice 时 (第1种)，就允许添加重复的edge，因为，`next`和`degree`都会对应增加。

### 有向图类型

* :yellow_circle: 课程表: [207. Course Schedule](https://leetcode.com/problems/course-schedule/)
    * 拓扑排序的基础题，一定要熟练掌握

* :yellow_circle: 课程表II: [210. Course Schedule II](https://leetcode.com/problems/course-schedule-ii/)
    * 拓扑排序的基础题，一定要熟练掌握

* :red_circle: 按照约束条件填表: [2392. Build a Matrix With Conditions](https://leetcode.com/problems/build-a-matrix-with-conditions/description/)

* :red_circle: DAG图中每个node的祖先: [2192. All Ancestors of a Node in a Directed Acyclic Graph](https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/description/)

* :red_circle: :secret: 舔狗链-圆桌会议可以邀请最多的互相喜欢的人数: [2127. Maximum Employees to Be Invited to a Meeting](https://leetcode.com/problems/maximum-employees-to-be-invited-to-a-meeting/description/)
    * trick: 计算节点`depth`

* :yellow_circle: 找出所有可以用给出的材料组成的食谱: [2115. Find All Possible Recipes from Given Supplies](https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/description/)

* red_circle: :secret: 平行课程III: [2050. Parallel Courses III](https://leetcode.com/problems/parallel-courses-iii/description/)
    * 转化思想：每门课的完成时间 等价于 计算节点`depth`
    * trick: 计算节点`depth`

### 无向图类型
* :red_circle: :secret: tree上捡硬币: [2603. Collect Coins in a Tree](https://leetcode.com/problems/collect-coins-in-a-tree/description/)
    * trick: 计算节点`depth`

* :red_circle: 分树-每个子树值相同: [2440. Create Components With Same Value](https://leetcode.com/problems/create-components-with-same-value/description/)