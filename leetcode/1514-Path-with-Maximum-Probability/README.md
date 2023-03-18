# [1514. Path with Maximum Probability](https://leetcode.com/problems/path-with-maximum-probability/)

## Solution idea

### Dijkstra

#### 要点总结
1. 比较常规的Dijkstra找最短路径的题目, 唯一的变化是: 求"最大概率"的路径 - 也就是，1. 路径的值变成weight相乘; 2. 用 maxHeap 挑选 "最大概率"


Time complexity = $O(E\log E)$

## Resource
[【每日一题】LeetCode 1514. Path with Maximum Probability](https://www.youtube.com/watch?v=L6DsIFXDyNs&ab_channel=HuifengGuan)