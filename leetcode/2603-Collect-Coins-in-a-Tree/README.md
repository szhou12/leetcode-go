# [2603. Collect Coins in a Tree](https://leetcode.com/problems/collect-coins-in-a-tree/description/)

## Solution idea

### BFS, Topological Sort

#### 要点总结
1. 解题逻辑链条:
    * 靠近中间有coin的node比较好收集，边缘的难收集 (难收集指要花更多的步数走更远)。所以，算法集中在怎么设计策略收集边缘的coins?
    * :arrow_right: 注意到，边缘路径上没有coin的node都可以剪枝砍掉
        * 实现：“剥洋葱 (BFS)” 从外到里，一层一层地把度数=1并且没有coin的node砍掉
    * :arrow_right: 砍掉后，叶子节点一定是含有coin的node
    * :arrow_right: 又注意到，题目说coin在两步以内就可以收集，说明还可以再“砍”, i.e. 边缘有coin的node其实不用实际上到达
    * :arrow_right: 定义 **深度**: 从 node i 到达最远的一个叶子节点的步数。
        * e.g. 最外层node (leaf node) 深度 = 1； 次外层node深度 = 2; ...
        * 实现：“剥洋葱 (BFS)” 从外到里，一层一层地计算深度
    * :arrow_right: 要访问的节点 = node 深度 $\geq 3$
    * :arrow_right: 起点node定在哪里？其实无所谓，因为反正题目要求需要最终原路返回起点
    * :arrow_right: 总步数 = `(m - 1) x 2` where `m` is number of nodes whose depth $\geq 3$



## Resource
[【每日一题】LeetCode 2603. Collect Coins in a Tree](https://www.youtube.com/watch?v=JAB27NlOfq0)