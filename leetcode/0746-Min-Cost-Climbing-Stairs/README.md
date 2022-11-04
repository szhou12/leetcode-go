# [746. Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/description/)

## Solution idea

### DP

* Definition
```
dp[i] = min cost to reach i-th cell ending at i
```
* Base Cases
```
dp[0] = cost[0]
dp[1] = cost[1]
```
* Recurrence
```
dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
```

Time complexity = $O(n)$, Space complexity = $O(n)$

## Resource
[代码随想录-746. 使用最小花费爬楼梯](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0746.%E4%BD%BF%E7%94%A8%E6%9C%80%E5%B0%8F%E8%8A%B1%E8%B4%B9%E7%88%AC%E6%A5%BC%E6%A2%AF.md)