# [2662. Minimum Cost of a Path With Special Roads](https://leetcode.com/problems/minimum-cost-of-a-path-with-special-roads/)

## Solution idea
### Dijkstra
#### 思路总结
1. 曼哈顿距离的性质 - 单增坐标情况下曼哈顿距离的计算
    * 假设: 某路径为 `source -> A -> B -> C -> target`，并且，每个节点的横纵坐标的值**依次单调递增** (i.e. for any node `i < j`, `xi <= xj` and `yi <= yj`)
    * 那么: `|target - source| = |A - source| + |B-A| + |C-B| + |target - C|`
```
(1, 1) -> (3, 5)
dist = |3 - 1| + |5 - 1| = 2 + 4 = 6

(1, 1) -> (2, 4) -> (3, 5)
dist = (|2 - 1| + |4 - 1|) + (|3 - 2| + |5 - 4|)
     = (2 - 1) + (4 - 1) + (3 - 2) + (5 - 4)
     = (-1, -1) + (2, 4) + (-2, -4) + (3, 5)
     = (-1, -1) + 0 + (3, 5)
     = 6

(1, 1) -> (7, 4) -> (3, 5) 不行!
dist = (|7 - 1| + |4 - 1|) + (|3 - 7| + |5 - 4|)
     = (7 - 1) + (4 - 1) + (7 - 3) + (5 - 4)
     = (-1, -1) + (7, 4) + (7, -4) + (-3, 5)
    != (-1, -1) + (7, 4) + (-7, -4) + (3, 5)
```
2. 本题中，从起点`source`到终点`target`的最短距离有两种情况:
    1. 途中不pick任何一条special road, 纯走曼哈顿距离
    2. 途中pick一条/多条special road
3. 上述情况一在实现中需要用到转化思想: 
    * 题目给出条件: `startX <= x1i, x2i <= targetX`, `startY <= y1i, y2i <= targetY`
    * 意味着: 横纵坐标值单调递增，就可以利用到上述曼哈顿距离的性质
    * 代码中: PQ中一开始除了要放入`source`节点, 还要放入从起点到每一个special road末端节点(B点)的曼哈顿距离
4. Dijkstra Loop 中，第一次Pop `source`，并且expand之后，PQ中就会同时存在从起点到任意一条special road末端节点(B点)的两种情况: 纯曼哈顿距离，pick本条路的距离

## Resource
[【每日一题】LeetCode 2662. Minimum Cost of a Path With Special Roads](https://www.youtube.com/watch?v=kQkJzCVQj-o&ab_channel=HuifengGuan)