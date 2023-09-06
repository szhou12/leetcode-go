# Breadth First Search & Dijkstra

## 目录
* [基础题: 层级遍历](#基础题-层级遍历)
* [常规题](#常规题)
* [综合题](#综合题)
* [岛屿沉没类/连通图个数](#岛屿沉没类---找neighbor)
* [Dijkstra](#dijkstra---find-shortest-path)
* [BFS + PQ](#bfs--pq-思路想bfs-代码结构像dijkstra)
* [Topological Sort](#topological-sort)
    * [解题思路](#解题思路)
    * [有向图类型](#有向图类型)
    * [无向图类型](#无向图类型)

## 基础题: 层级遍历

* 基础: [102. Binary Tree Level Order Traversal](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0102-Binary-Tree-Level-Order-Traversal)

* 进阶: [103. Binary Tree Zigzag Level Order Traversal](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0103-Binary-Tree-Zigzag-Level-Order-Traversal)

* 基础: [116. Populating Next Right Pointers in Each Node](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0116-Populating-Next-Right-Pointers-in-Each-Node)

* 基础: [513. Find Bottom Left Tree Value](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0513-Find-Bottom-Left-Tree-Value)



## 常规题
* :red_circle: :secret: 在一个图中按要求给节点分组，求最大分组数: [2493. Divide Nodes Into the Maximum Number of Groups](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2493-Divide-Nodes-Into-the-Maximum-Number-of-Groups)

* :green_circle: 引爆最多炸弹数: [2101. Detonate the Maximum Bombs](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2101-Detonate-the-Maximum-Bombs)



## 综合题

* :red_circle: :red_circle: :secret: 逃离野火: [2258. Escape the Spreading Fire](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2258-Escape-the-Spreading-Fire)
    * Solution 1: BFS + BFS + Reasoning

* :red_circle: :secret: **Grid中寻找安全路径:** [2812. Find the Safest Path in a Grid](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2812-Find-the-Safest-Path-in-a-Grid)
   * BFS (类Topological Sort式剥洋葱) + Binary Search 猜答案 + 嵌套BFS找路径

* :red_circle: **从Grid左上角走到右下角所花最少cost:** [1368. Minimum Cost to Make at Least One Valid Path in a Grid](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1368-Minimum-Cost-to-Make-at-Least-One-Valid-Path-in-a-Grid)
    * BFS 嵌套 BFS

## 岛屿沉没类/连通图个数 - 找neighbor
* :red_circle: **连通岛屿的个数:** [200. Number of Islands](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0200-Number-of-Islands)

* :red_circle: **四面八方, 潮水涌来:** [417. Pacific Atlantic Water Flow](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0417-Pacific-Atlantic-Water-Flow)

* :red_circle: **找出不包含边界元素的连通图:** [130. Surrounded Regions](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0130-Surrounded-Regions)

* :red_circle: **飞地的个数:** [1020. Number of Enclaves](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1020-Number-of-Enclaves)

* Roblox - Candy Crush

## Dijkstra - Find Shortest Path

* :star: **适用题目的特征:**
    1. **single-source**: 只给一个起点, 求从起点(single source)到图上任意一个node的最短路径
    2. **non-negative weight**: edge的权重非负
    * 这样, 就可以求图上任意一点到起点single-source的最短距离
    * Time complexity = $O(E\log E)$
    * 两种Return方式:
        1. Early Return: Dijkstra loop内到达终点即return; loop结束还没return说明无法到达终点
        2. 填表 Return: Dijkstra把所有结果填入一个表 (array, matrix) 中, loop结束后把所求node的结果按要求从表中取出return

* :red_circle: **调整边的权重得到最短路径:** [2699. Modify Graph Edge Weights](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2699-Modify-Graph-Edge-Weights)

* :red_circle: **有特殊路径选择权的最短路径:** [2662. Minimum Cost of a Path With Special Roads](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2662-Minimum-Cost-of-a-Path-With-Special-Roads)

* :red_circle: **从Grid左上角走到右下角所花最短时间:** [2577. Minimum Time to Visit a Cell In a Grid](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2577-Minimum-Time-to-Visit-a-Cell-In-a-Grid)
    * **矩阵走格子类型题**

* :red_circle: **从Grid左上角走到右下角所花最少cost:** [1368. Minimum Cost to Make at Least One Valid Path in a Grid](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1368-Minimum-Cost-to-Make-at-Least-One-Valid-Path-in-a-Grid)
    * **矩阵走格子类型题**

* :red_circle: **求必经最少障碍物的路径:** [2290. Minimum Obstacle Removal to Reach Corner](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2290-Minimum-Obstacle-Removal-to-Reach-Corner)
    * **矩阵走格子类型题**

* :red_circle: **从两个起点到达同一个终点所需最短路径:** [2203. Minimum Weighted Subgraph With the Required Paths](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2203-Minimum-Weighted-Subgraph-With-the-Required-Paths)
    * 难点: "拉链" - 如何找到"拉链"node i

* :red_circle: :secret: **到达终点的路径总数:** [1976. Number of Ways to Arrive at Destination](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1976-Number-of-Ways-to-Arrive-at-Destination)
    * **Dijkstra + DFS + Deduplication**
    * 题目求方法总数, 最自然的是使用DFS
    * 使用Memo避免DFS重复走走过的路径
    * *1976 与 1786 两题逻辑完全相同*

* :red_circle: :secret: **从起点走到终点的特殊路径总数:** [1786. Number of Restricted Paths From First to Last Node](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1786-Number-of-Restricted-Paths-From-First-to-Last-Node)
    * **Dijkstra + DFS + Deduplication**
    * 题目求方法总数, 最自然的是使用DFS
    * 使用Memo避免DFS重复走走过的路径
    * *1976 与 1786 两题逻辑完全相同*

* :red_circle: **从起点可到达的nodes总数:** [882. Reachable Nodes In Subdivided Graph](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0882-Reachable-Nodes-In-Subdivided-Graph)
    * 难点1: 如何把解题思路与Dijkstra做关联 --> 把 小node 看做 edge weight
    * 难点2: Dijkstra expand时添加限制条件

* :red_circle: :secret: **K站中转内最便宜的航班路径:** [787. Cheapest Flights Within K Stops](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0787-Cheapest-Flights-Within-K-Stops)
    * **node储存双状态**: node 储存 位置信息+已中转次数信息
    * Dijkstra结果储存进 二维矩阵 = (# of nodes, # of stops used so far)
    * Early Return: Dijkstra loop内到达终点即return
    * *0787 与 2093 两题逻辑相似*

* :yellow_circle: :secret: :lock: **最小花费到达目的地:** [2093. Minimum Cost to Reach City With Discounts](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2093-Minimum-Cost-to-Reach-City-With-Discounts)
    * 每次Expand一个邻居node, push "两条边": toll edge; toll/2 edge
    * **node储存双状态**: node 储存 位置信息+已使用discounts次数信息
    * Dijkstra结果储存进 二维矩阵 = (# of nodes, # of discounts used so far)
    * Early Return: Dijkstra loop内到达终点即return
    * *0787 与 2093 两题逻辑相似*

* :green_circle: **最高概率的路径:** [1514. Path with Maximum Probability](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1514-Path-with-Maximum-Probability)

* :green_circle: **从source node到所有node所需最短时间:** [743. Network Delay Time](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0743-Network-Delay-Time)
    * Dijkstra求最短路径的基础题: 求从source node到任意node i的最短路径
    * 可用来作为以上题目的练手题

* :yellow_circle: :secret: :lock: **走迷宫II / 走冰系道馆:** [505. The Maze II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0505-The-Maze-II)
    * 难点: 地面是冰面, 只能维持一个方向前进直到撞墙
    * Early Return: Dijkstra loop内到达终点即return

* :red_circle: :secret: :lock: **走迷宫III / 走冰系道馆+掉洞:** [499. The Maze III](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0499-The-Maze-III)
    * **node储存双状态**: node 储存 位置信息+到达node的方向指令信息
    * Priority Queue 再定义:
        1. 用一个结构体容纳不同variable type (int, string)
        2. `Less(i, j int)`: 双排序 - 先按照 path cost 排序, 再按照 指令的lexicographical order排序
    * 维持一个方向滑冰时, 如果掉洞，直接break，不再往前滑
    * Early Return: Dijkstra loop内到达终点即return



## BFS + PQ: 思路像BFS, 代码结构像Dijkstra
* :red_circle: **一个query可以淹没多少格子, 求淹没总数:** [2503. Maximum Number of Points From Grid Queries](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2503-Maximum-Number-of-Points-From-Grid-Queries)

* :red_circle: **升高海平面中的最短路线:** [778. Swim in Rising Water](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0778-Swim-in-Rising-Water)

* :red_circle: **接雨水II (海水倒灌II):** [407. Trapping Rain Water II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0407-Trapping-Rain-Water-II)
    * 题目设置很像 [417. Pacific Atlantic Water Flow](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0417-Pacific-Atlantic-Water-Flow)



## Topological Sort

### 解题思路
* :bulb: **农村包围城市/剥洋葱**: 优先访问那些入度最低的节点。删去第一批最外围的节点后，再继续访问此时入度更新为最低的节点。依次类推。
    * 编者注：想像成“蜘蛛网”，我们从边缘(入度最低)出发，逐步向中心(入度最高)侵略
* :bulb: 使用的算法/数据结构: BFS
    * 注: 拓扑排序也可以用DFS实现，但个人感觉BFS的思路更直观更易模版化且算法在空间上也更高效，遂练习时着重使用BFS解题
* :bulb: 使用条件: **Acyclic** (在无向图中就意味着必须是一个Tree)
* :bulb: 入度/degree 的定义: 一般的, **degree = the number of incoming edges**.
    * 注: degree 的定义也可能根据具体题目的要求进行调整，不一定死板地遵守一般定义
* :bulb: 外围 (低阶级) vs 内层 (高阶级): 一般的，最外围的节点 = 入度的值最低的节点
    * 注: 最外围节点的定义也可能根据具体题目的要求增添额外的判定条件，但入度的值一般是其中一条必要的判定条件
* :bulb: 有向图 vs 无向图:
    * 最低入度值 (**最低入度值是非常重要的判定外围node的条件！**):
        * **有向图**中node入度最低 = 0
        * **无向图**中node入度最低 = 1
    * Check for visited (防止从内层重新走回外围的机会):
        * **有向图**无需 check for visited, 因为有向就意味着从内层走回外围的机会根本不可能发生。
        * **无向图**需要 check for visited, 除非next move时直接过河拆桥, 把从内层走回外围的edge删除了 (e.g. [2603. Collect Coins in a Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2603-Collect-Coins-in-a-Tree))
    * visited 另一作用 - 给非环节点染色:
        * **有向图**中visited一般用于检测环, visited起到"染色"的作用 - 即，visited==1说明节点不属于环的一部分，visited==0说明节点属于一个环。因为，能被访问说明节点入度可变为0，不能访问说明节点入度不可变为0，也就是，它在一个环中。
        * 注：染色作用的 visited 实际上只是命名为 visited，其他名字也可以，visited只是比较方便理解
* :mag: Adjacency List `next` 的 data structure 选择:
    1. `[][]int`: slice of slices
    2. `[]map[int]bool`: slice of maps
    3. `map[string]map[string]bool`: map of maps (json)
    * 总结: 怎么方便怎么来。但是，要注意！选择 map 相关的结构时 (第2, 3种)，要额外考虑 duplicated edges 的情况，重复的edge要跳过，因为，`next`不会重复添加，但是`degree`会不小心多+1。选择 slice 时 (第1种)，就允许添加重复的edge，因为，`next`和`degree`都会对应增加。
* :mag: **常涉及的小技巧:** 一层一层剥洋葱的时候会需要一个 "继承"变量。"继承"变量可以是节点的深度、祖辈/父辈的某个信息。
    * 怎么设计这个"继承"变量的通常思路是 `DP`
* :bulb: Topological Sort 的排序数列 什么情况下是唯一确定的？
    * 每一层入度=0的节点只有一个
    * i.e., `queue` contains ONLY one element at any time (size == 1 at all times)

### 有向图类型

* :yellow_circle: **课程表:** [207. Course Schedule](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0207-Course-Schedule)
    * 拓扑排序的基础题，一定要熟练掌握

* :yellow_circle: **课程表II:** [210. Course Schedule II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0210-Course-Schedule-II)
    * 拓扑排序的基础题，一定要熟练掌握

* :red_circle: :secret: **课程表IV:** [1462. Course Schedule IV](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1462-Course-Schedule-IV)
    * "继承"变量 trick: 下一步节点要继承`cur`以及`cur`的所有prerequisites

* :red_circle: **按照约束条件填表:** [2392. Build a Matrix With Conditions](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2392-Build-a-Matrix-With-Conditions)

* :red_circle: **DAG图中每个node的祖先:** [2192. All Ancestors of a Node in a Directed Acyclic Graph](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2192-All-Ancestors-of-a-Node-in-a-Directed-Acyclic-Graph)

* :red_circle: :secret: **舔狗链-圆桌会议可以邀请最多的互相喜欢的人数:** [2127. Maximum Employees to Be Invited to a Meeting](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2127-Maximum-Employees-to-Be-Invited-to-a-Meeting)
    * "继承"变量 trick: 计算节点`depth`

* :yellow_circle: **找出所有可以用给出的材料组成的食谱:** [2115. Find All Possible Recipes from Given Supplies](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2115-Find-All-Possible-Recipes-from-Given-Supplies)


* :red_circle: :secret: **有向图中最多的颜色:** [1857. Largest Color Value in a Directed Graph](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1857-Largest-Color-Value-in-a-Directed-Graph)
    * "继承"变量 trick: 计算节点`depth`

* :red_circle: :secret: **矩阵以rank表示:** [1632. Rank Transform of a Matrix](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1632-Rank-Transform-of-a-Matrix)
    * Topological Sort + Union Find

* :red_circle: :secret: **奇怪打印机II:** [1591. Strange Printer II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1591-Strange-Printer-II)


* :red_circle: :secret: **平行课程III:** [2050. Parallel Courses III](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2050-Parallel-Courses-III)
    * 转化思想：每门课的完成时间 等价于 计算节点`depth`
    * "继承"变量 trick: 计算节点`depth`

* :yellow_circle: :lock: **平行课程:** [1136. Parallel Courses](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1136-Parallel-Courses)
    * "继承"变量 trick: 计算节点`depth`

* :red_circle: **排序有组别的元素:** [1203. Sort Items by Groups Respecting Dependencies](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1203-Sort-Items-by-Groups-Respecting-Dependencies)
    * 难点不在于解法思路，而在于实现

* :red_circle: :secret: **寻找可以走到终点的起点:** [802. Find Eventual Safe States](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0802-Find-Eventual-Safe-States)
    * 突破点: 在于关注节点的"出度"。也就是要把graph整个逆向，计算每个节点的出度，"洋葱"的最外层是 出度=0 的节点。

* :red_circle: :secret: :lock: **拓扑排序唯一解:** [444. Sequence Reconstruction](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0444-Sequence-Reconstruction)
    * 知识点: 拓扑排序有唯一解 = 每一时刻入度=0的节点只有一个 = `queue`每一时刻都只会有一个元素

* :red_circle: :secret: :lock: **推断字典序:** [269. Alien Dictionary](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0269-Alien-Dictionary)
    * trick: 利用 lexicographical order 来构建 adjacency-list representation


### 无向图类型
* :red_circle: :secret: **tree上捡硬币:** [2603. Collect Coins in a Tree]()
    * "继承"变量 trick: 计算节点`depth`

* :red_circle: **分树-每个子树值相同:** [2440. Create Components With Same Value]()

* :red_circle: :secret: **找到所有最矮树的根节点:** [310. Minimum Height Trees]()

* :red_circle: :secret: :lock: **所有节点到到环的最短距离:** [2204. Distance to a Cycle in Undirected Graph]()
