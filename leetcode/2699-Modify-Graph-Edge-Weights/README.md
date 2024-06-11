# [2699. Modify Graph Edge Weights](https://leetcode.com/problems/modify-graph-edge-weights/description/)

## Solution idea
### Dijkstra
#### 思路总结
1. 题目要求所有权重为-1的边一定要修改成正数。所以，第一步先把所有负数边的权重改成最小正数1，跑一遍Dijkstra，看看最短路径的总消耗是否等于target。如果 > target，显然找不到任何可行解；只有 <= target才可能存在解。但是，当 < target时，同时还需要满足最短路径途经至少一条“修改”边，否则也不存在可行解。

## Resource
[【每日一题】LeetCode 2699. Modify Graph Edge Weights](https://www.youtube.com/watch?v=gchifH4Ftjc&list=PLwdV8xC1EWHrtgsYCcDTXIMVaHSlsnLzL&index=14&ab_channel=HuifengGuan)