# [2305. Fair Distribution of Cookies](https://leetcode.com/problems/fair-distribution-of-cookies/)

## Solution idea

### DFS

DFS暴力解

每一层是当前cookie。 每一个分支是一个人，总共`k`个分支。

Time complexity = $O(k^n)$

### Binary Search + DFS

**思路**: 如果**直接找答案** (DFS暴力解) 不够efficient，那么，就反其道而行之，**猜答案**.

**猜答案**的方法：**Binary Search** - 对每个人最多能拿到的数量进行猜测，如果当前数量可行，可能是答案，再往低了猜。

## Resource

[【每日一题】LeetCode 2305. Fair Distribution of Cookies](https://www.youtube.com/watch?v=tAIzu_MWa8U&t=840s)