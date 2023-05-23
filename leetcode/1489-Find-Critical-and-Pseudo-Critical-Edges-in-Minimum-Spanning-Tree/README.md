# [1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree](https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/description/)

## Solution idea
### Union Find + Kruskal
#### 思路总结
1. 这道题要求找MST中两种有特殊性质的边，容易想到类似的题目 - 找redundant edge。怎么做？
    * 大致想法是：先用 Kruskal 算法出全局最小的MST，然后逐条边检验，与全局最小的 MST做比较，找出符合要求的边。
2. 大致的框架有了，下一步是怎么找出符合要求的边？
    * 重中之重是明确 关键边 (critical edge) 和 伪关键边 (pseudo-critical edge) 的定义，并分析出两者之间的关系。
    * 关键边 (critical edge):
        * 题目中给的定义是: An MST edge whose deletion from the graph would cause the MST weight to increase.
        * 怎么**结合全局最小MST**进行理解？
            * Delete this edge, we will end up with a larger-valued MST.
            * 删除了这条边，我们无法得到 全局最小MST。
            * 隐含暗示：这条边必然是构建 全局最小MST 的边之一。全局最小MST一定用了这条边。
    * 伪关键边 (pseudo-critical edge):
        * 题目中给的定义是: a pseudo-critical edge is that which can appear in some MSTs but not all.
        * 什么意思？也就是说，可以找到“平替”
        * 怎么**结合全局最小MST**进行理解？
            * Delete this edge, we will end up with a same-valued MST.
            * 删除了这条边，我们依然用一个"平替"得到 全局最小MST。
            * 隐含暗示：这条边对于构建 全局最小MST 是可有可无的。全局最小MST可能用了这条边，也可能没用，但是一个可用的候补。
    * 明确了定义，我们可以总结出它们之间的关系：
        * 关键边 (critical edge) 和 伪关键边 (pseudo-critical edge) 是 **互斥** 的。
        * 一条边 如果是关键边，那它就一定不是 伪关键边；反之亦然。
        * 如果把 可以参与构建全局最小MST的边(所有的候补选手) 看作一个集合，那么 关键边集合 和 伪关键边集合 是互补关系。

## Resource
[【每日一题】1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree, 9/14/2020](https://www.youtube.com/watch?v=ozlYJy-2ZEY&ab_channel=happygirlzt)