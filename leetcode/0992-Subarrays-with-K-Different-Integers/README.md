# [992. Subarrays with K Different Integers](https://leetcode.com/problems/subarrays-with-k-different-integers/description/)

## Solution idea
### Sliding Window (Flex)
#### 思路总结
1. 题目要求滑窗内框住 eaxctly k 个不同数字。我的思路是：对于每一个左边界
    1. 确定lower bound: 最短滑窗(第一个index)可以框住eaxctly k 个不同数字。
    2. 确定upper bound: 最长滑窗(最后一个index)可以框住eaxctly k 个不同数字。
    3. 对于当前左边界可以框住 exactly k 个不同数字的subarray个数 = upper bound - lower bound + 1
2. 优化解法：
    1. 对于当前左边界可以框住 exactly k 个不同数字的subarray个数 = upperBound(k) - upperBound(k - 1)
    2. 这样可以只写一个 helper function

Time complexity = $O(n)$

## Resource
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Two_Pointers/992.Subarrays-with-K-Different-Integers)