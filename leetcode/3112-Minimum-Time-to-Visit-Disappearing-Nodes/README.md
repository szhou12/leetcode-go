# [3112. Minimum Time to Visit Disappearing Nodes](https://leetcode.com/problems/minimum-time-to-visit-disappearing-nodes/description/)

## Solution idea
### Dijkstra
1. 典型题。from node 0确定是single-source (唯一起点) + 所有边权都是正数 = 确认使用Dijkstra
2. 题目中`disappear[]`的意义：每次pop current node时，如果最短路径到达时已经超过了"时限"，则不能更新到达的最短路径，同时，无法从current node进行expand。

Time complexity = $O(E \log E)$

## Resource
[【每日一题】LeetCode 3112. Minimum Time to Visit Disappearing Nodes](https://www.youtube.com/watch?v=JPo1wy8FPTo&ab_channel=HuifengGuan)