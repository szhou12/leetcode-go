# [684. Redundant Connection](https://leetcode.com/problems/redundant-connection/)

## Solution idea

### UnionFind 基础题

* 遍历每一条边
    * 如果这条边两端的节点已经有同一个祖宗, 说明已经连接过, 当前的边会导致成环
    * 如果不是, 两端节点联姻 (ie. 加入同一集合/同一根节点)

Time complexity = $O(n)$

## Resource

[【每日一题】684. Redundant Connection, 10/01/2019](https://www.youtube.com/watch?v=8u-sjzyHjDg)

[代码随想录-684.冗余连接](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0684.%E5%86%97%E4%BD%99%E8%BF%9E%E6%8E%A5.md)

[time complexity analysis](https://massivealgorithms.blogspot.com/2018/05/leetcode-684-redundant-connection.html)