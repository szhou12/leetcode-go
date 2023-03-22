# [882. Reachable Nodes In Subdivided Graph](https://leetcode.com/problems/reachable-nodes-in-subdivided-graph/)

## Solution idea
### Dijkstra

#### 要点总结
![leetcode882](https://user-images.githubusercontent.com/35708194/226774392-a7f47a7f-77a7-463c-8c3a-fd89d5140bc9.png)
1. `cnt`转化成weight -> 可用Dijkstra求大点之间至少需要多少步
2. 刨去Dijkstra求得的从起点到每个大点之间所用最少步数之后, 任意两个大点之间还可以走多少步 (一步就是一个小点): `(maxMoves-Dij[x])+(maxMoves-Dij[y])`
    * 我的最初想法: 是把小点当作大点看待, 然后从起点开始用BFS计算最远可以到哪里. 但是问题有两个: 1. 小点要重新label和储存; 2. 题目给出的小点的量级为$10^4$, 用BFS太贵
    * 转换思路: 如上图, 任意两个大点`x`和`y`之间可以走到的小点个数分别是从`x`出发还可以走多少步 (`i`)+从`y`出发还可以走多少步 (`j`), 这样就可以利用上Dijkstra, 再遍历每条边计算小点就可以


Time Complexity = $O(E\log E)$


## Resource
[【每日一题】882. Reachable Nodes In Subdivided Graph, 2/21/2021](https://www.youtube.com/watch?v=TYYkcdeVTi8&t=1356s&ab_channel=HuifengGuan)
