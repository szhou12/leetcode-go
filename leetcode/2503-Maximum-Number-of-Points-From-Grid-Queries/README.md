# [2503. Maximum Number of Points From Grid Queries](https://leetcode.com/problems/maximum-number-of-points-from-grid-queries/)

## Solution idea

### BFS + PQ

#### 要点总结
1. 像这类题目给一系列 queries, **最常用的技巧**是: 分析query彼此之间是否存在某种联系? 是否每个query独立处理？能否在解决一个query的同时对解决另一个有帮助？query之间能否有一些信息共享？
    * 对于本题: 我们发现, query值越小, 能走的格子越少; query值越大, 能走的格子越多。而且！大query走的格子包含了小query走的格子, 也就是说, 小query能走的格子是大query能走的格子的子集。
    * 这个规律暗示我们: 处理顺序 小query $\Rightarrow $ 大query，走过的格子数累加起来就行。"在小query走的格子基础上又能走多少格子"。这个规律就是简单的BFS的遍历。
2. 既然BFS的总体思路确定了，那么，具体策略怎么做?
    * 注意到，对于一个query的值，哪些格子首当其冲先背"淹没"？显然是 **格子的值越小的那些先被牺牲**。 这句话就是在暗示: 这道题BFS的queue中的元素(格子)需要排一个顺序 - 值越小的元素排在前面
    * 策略: BFS在expand时(从queue中Pop时)总是值最小的先出来, 直到queue中所有元素都大于当前query的值, 这样, 我们就找到了所有可以“淹没”的格子
3. 那么怎么实现？怎么让一个queue维护元素的顺序?
    * 显然, 用到 Priority Queue

Time complexity = $O(mn\log mn)$

## Resource
[【每日一题】LeetCode 2503. Maximum Number of Points From Grid Queries](https://www.youtube.com/watch?v=G7Gg9w5_KWk&t=5s&ab_channel=HuifengGuan)