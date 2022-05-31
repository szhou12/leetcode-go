# [2272. Substring With Largest Variance](https://leetcode.com/problems/substring-with-largest-variance/)

## Note:

Neet to use two DP arrays:

DP1[i]: the largest sum subarray ending at i (left-to-right)

DP2[i]: the largest sum subarray starting at i (right-to-left)

result = DP1[i] + DP2[i] - nums[i]