# [213. House Robber II](https://leetcode.com/problems/house-robber-ii/)

## Solution idea

### DP

* **突破口**：这道题是环形的，所以, **首尾房间不能同时被抢**, 那么就会产生三种不同情况:
    1. 要么都不被抢
    2. 要么第一间房子被抢, 最后一间不抢
    3. 要么最后一间房子被抢, 第一间不抢
    * 而情况二 和 情况三 都包含了情况一了，所以只考虑情况二和情况三就可以了

* 具体的DP实现与[198. House Robber](https://leetcode.com/problems/house-robber/)一致

Time complexity = $O(n)$

## Resource

[代码随想录-213.打家劫舍II](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0213.%E6%89%93%E5%AE%B6%E5%8A%AB%E8%88%8DII.md)