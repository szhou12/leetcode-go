# [685. Redundant Connection II](https://leetcode.com/problems/redundant-connection-ii/)

## Solution idea
### Union Find
#### 思路总结
1. 解题之前先问自己一个问题：一棵树随意多加一笔会出现什么情况？提示：一棵树中root节点入度=0，其他节点入度=1
    * 答案：有一个节点会多加1个入度。如果这个节点是root，那么每个节点的入度都是1；如果是其他的节点，那么这个节点的入度就变成2，也就是它会有两个父节点。

## Resource
[【每日一题】685. Redundant Connection II, 10/02/2019](https://www.youtube.com/watch?v=DGwJtS-8iG8&ab_channel=%E5%82%85%E7%A0%81%E7%88%B7)