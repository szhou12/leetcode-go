# [2976. Minimum Cost to Convert String I](https://leetcode.com/problems/minimum-cost-to-convert-string-i/description/)

## Solution idea
### Floyd-Warshall Algorithm
1. 破题：`original[]`, `changed[]`, `cost[]` 三者想象成有向图。通过 Floyd-Warshall 算法，求出任意两点之间的最短路径，即可得到答案。
2. 注：我的解法使用 Floyd 算法模版的Variation。与标准答案 (标准模版) 稍微不同。

Time complexity = $O(n^3)$

## Resource
[【每日一题】LeetCode 2976. Minimum Cost to Convert String I](https://www.youtube.com/watch?v=I2-r1cHY-74&ab_channel=HuifengGuan)