# [392. Is Subsequence](https://leetcode.com/problems/is-subsequence/)

## Solution idea

**Greedy Algorithm**: 直接按顺序检查是否所有必要的characters都存在

Time complexity = $O(\max \{m, n\})$ where $m$ is length of `s`, n is length of `t`

**Binary Search**:

记录`t`中每个char的index，再对`s`的每个char进行binary search to find the smallest value > target

Time complexity = $O(m\log n)$

进阶题：[792. Number of Matching Subsequences](https://leetcode.com/problems/number-of-matching-subsequences/)