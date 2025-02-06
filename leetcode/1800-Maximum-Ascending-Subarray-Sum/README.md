# [1800. Maximum Ascending Subarray Sum](https://leetcode.com/problems/maximum-ascending-subarray-sum/description/)

## Solution idea
用一个缓存`cur`记录当前的递增subarray的和，如果当前元素比上一个大，累加`cur`。否则，`cur`断开，另起炉灶。

Time complexity = $O(n)$
