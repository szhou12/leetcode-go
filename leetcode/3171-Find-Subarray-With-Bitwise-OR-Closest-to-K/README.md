# [3171. Find Subarray With Bitwise OR Closest to K](https://leetcode.com/problems/find-subarray-with-bitwise-or-closest-to-k/description/)

## Solution idea
### Binary Search + Bitwise OR 单调性 + Segment Tree

Time complexity = $O(n\cdot \log n \cdot \log n)$ where the first $\log n$ is for binary search, the second $\log n$ is for segment tree querying range.

## Resrouce
[【每日一题】LeetCode 3171. Find Subarray With Bitwise AND Closest to K](https://www.youtube.com/watch?v=DWhQ9efbT84&ab_channel=HuifengGuan)