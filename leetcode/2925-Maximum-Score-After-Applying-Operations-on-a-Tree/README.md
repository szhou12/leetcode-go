# [2925. Maximum Score After Applying Operations on a Tree](https://leetcode.com/problems/maximum-score-after-applying-operations-on-a-tree/description/)

## Solution idea
### DFS
1. 首先最重要的是把题目理解清楚!
    1. operation: 相当于每个node有一个不同数额的金币，一次取走了就没了。
    2. healthy: 每一条path至少要有一个node是不能取走金币的。
    3. 目的: 最大化取走的金币总数额。
2. Tree :arrow_right: 天然的Recursive结构 :arrow_right: 95% Tree类结构考虑使用Recursion解题 :arrow_right: DFS :arrow_right: 把Graph里DFS的代码思路 应用到 Tree里DFS
3. Generalization: 如何把 root的所求 泛化到 树上任何node为root的subtree的所求？
    * 考虑 树上任何node为root的subtree
    * 如果要保证这棵subtree每一条到leaf的path都"healthy"，那么可以归结为两种情况：
        1. 不取走 subtree root 金币：那么除了root之外的subtree的所有node都可以取走金币。
        2. 取走 subtree root 金币：那么root连结的所有subtree自己想办法决定保留哪个node的金币。相当于又回到了从这两种情况中选择最优解。这里用到Recursion。
    * 最后，这棵 subtree 的 max score 就是这两种情况中的最大值。

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 2925. Maximum Score After Applying Operations on a Tree](https://www.youtube.com/watch?v=oxcUMxxcPcU&ab_channel=HuifengGuan)