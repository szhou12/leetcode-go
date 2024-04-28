# [2771. Longest Non-decreasing Subarray From Two Arrays](https://leetcode.com/problems/longest-non-decreasing-subarray-from-two-arrays/description/)

## Solution idea
### 2-D DP
1. 比较容易的题。
    1. max subarray length: “回头看，另起炉灶”
    2. DP第二维表示两种情况分类讨论：i位取较小的值；i位取较大的值。

Time complexity = $O(n)$