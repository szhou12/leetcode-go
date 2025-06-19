# [3355. Zero Array Transformation I](https://leetcode.com/problems/zero-array-transformation-i/description/)

## Solution idea
### Difference Array
1. The problem is basically asking for range decrement by 1 for each query. So we can make use of property 1 of difference array to conduct range update.

Time complexity = $O(n + q)$
Space complexity = $O(n)$
