# [2563. Count the Number of Fair Pairs](https://leetcode.com/problems/count-the-number-of-fair-pairs/description/)

## Solution idea
### Binary Search
#### 思路总结
1. 题目只要求计算数量，不要求列出具体的pair。所以，就可以先排序，然后使用Binary Search。
2. 题目要求 `lower <= nums[i] + nums[j] <= upper`. 也就是说，对于每个 `nums[i]` 对应的 `nums[j]` 的范围是 `[lower - nums[i], upper - nums[i]]`。这可以抽象成类型题：对于一个数`x`，求出满足条件的**区间**. 求区间的问题，可以使用`upperBound()`, `lowerBound()`来解决。

## Resource
[【每日一题】LeetCode 2563. Count the Number of Fair Pairs](https://www.youtube.com/watch?v=MmegIOi5Rrw&ab_channel=HuifengGuan)