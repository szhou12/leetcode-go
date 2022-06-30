# [674. Longest Continuous Increasing Subsequence](https://leetcode.com/problems/longest-continuous-increasing-subsequence/)

## Solution idea
Define $DP[i] = $ longest length of the subarray ending at i

Base case: $DP[0] = 1$

Recurrence:

$DP[i] = DP[i-1] + 1$ if $nums[i] > nums[i-1]$

$DP[i] = 1$ otherwise

Time complexity = $O(n)$