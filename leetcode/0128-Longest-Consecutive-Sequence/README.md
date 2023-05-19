# [128. Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/description/)

## Solution idea
### Union Find
#### 思路总结
1. 如何转化为 Union Find 解题？
    * 对于每一个元素 `a`，如果它能与参与组成 consecutive sequence，那么说明它的邻居 `a-1` 或者 `a+1` 也在数组中。
    * 所以，我们都检查 `a-1` 和 `a+1` 是否存在，如果存在，就将 `a` 与 `a-1` 或 `a+1` 进行合并。
2. 注意：在记录每个家族有哪些成员时，要用 `map[int]bool`，不能用`[]int`，因为数组中会出现重复的元素, `[]int`会count多次

Time Complexity = $O(n)$

## Resource
[【每日一题】128. Longest Consecutive Sequence, 09/30/2019](https://www.youtube.com/watch?v=QnBcLxgeeGs&ab_channel=HuifengGuan)