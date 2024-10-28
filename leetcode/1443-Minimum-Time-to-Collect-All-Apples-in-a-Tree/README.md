# [1443. Minimum Time to Collect All Apples in a Tree](https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/description/)

## Solution idea
### DFS
站在一个node上，怎么知道左孩子和右孩子值不值得往下探？

一个孩子节点值得探 if: 1) 这个孩子节点本身有苹果 OR 2) 以这个孩子节点为根的子树中有苹果

定义 DFS function: 返回值为走当前节点为子树往下探索，摘取子树中所有苹果的最小时间。如果子树中没有苹果，返回0。

所以，当前节点的任意一个孩子节点是否值得探索？看这个孩子节点返上来的时间是否 > 0 或者 这个孩子节点本身是否有苹果。

Time complexity = $O(n)$

## Resource
[Minimum Time to Collect All Apples in a Tree - Leetcode 1443 - Python](https://www.youtube.com/watch?v=Xdt5Z583auM&ab_channel=NeetCodeIO)