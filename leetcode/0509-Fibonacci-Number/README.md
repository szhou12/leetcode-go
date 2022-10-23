# [509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)

## Solution idea

### DP

* 比较经典 & 简单的 DP 题
```
Definition:
    dp[i] = F(i) from 0 to i
Base Cases:
    dp[0] = 0
    dp[1] = 1
Recurrence:
    dp[i] = dp[i-2] + dp[i-1]
```
* 也可以用 **滚动数组** 只保留前两个元素，来优化算法

Time complexity = $O(n)$