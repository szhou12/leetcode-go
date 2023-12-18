# [2959. Number of Possible Sets of Closing Branches](https://leetcode.com/problems/number-of-possible-sets-of-closing-branches/description/)

## Solution idea
### Floyd-Warshall Algorithm
1. 题目要求：找到所有合法的subsets。根据Constraint, 可以直接brute force所有的subsets。
2. 题目要求：关闭一些节点后，任意两个开放的节点之间的最短路径不超过一个上限。因为是求任意两个节点之间的最短距离 并且 可以brute force，使用Floyd-Warshall算法。

Time complexity = $O(2^n \times E \times n^2)$

## Resrouce
[【每日一题】LeetCode 2959. Number of Possible Sets of Closing Branches](https://www.youtube.com/watch?v=vJHNiWY4go4&ab_channel=HuifengGuan)