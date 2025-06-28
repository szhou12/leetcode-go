# [2200. Find All K-Distant Indices in an Array](https://leetcode.com/problems/find-all-k-distant-indices-in-an-array/description/)

## Solution idea
1. For an index `j` where `nums[j] == key`, all valid indices `i` are in the range `[max(j-k, 0), min(j+k, n-1)]`

Time complexity = $O(n)$

Space complexity = $O(1)$