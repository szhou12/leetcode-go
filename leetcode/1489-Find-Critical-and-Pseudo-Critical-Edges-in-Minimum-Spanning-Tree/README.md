# [1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree](https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/description/)

## Solution idea
### Union Find + Kruskal
#### 思路总结
1. 这道题要求找MST中两种有特殊性质的边，容易想到类似的题目 - 找redundant edge。怎么做？大致想法是：先用 Kruskal 算法出全局最小的MST，然后逐条边检验，与全局最小的 MST做比较，找出符合要求的边。

## Resource
[【每日一题】1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree, 9/14/2020](https://www.youtube.com/watch?v=ozlYJy-2ZEY&ab_channel=happygirlzt)