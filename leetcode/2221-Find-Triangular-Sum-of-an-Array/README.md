# [2221. Find Triangular Sum of an Array](https://leetcode.com/problems/find-triangular-sum-of-an-array/description/)

## Solution idea
### Simulation
因为数据量不大(<1000)，按照题意进行全模拟就行：每个`nums[i]`被`(nums[i]+nums[i+1])%10`替换，一共进行n-1轮，最后返回头元素。

Time complexity = $O(n^2)$

## Resource
[【每日一题】LeetCode 2221. Find Triangular Sum of an Array](https://www.youtube.com/watch?v=RzoHl7M9xvM&ab_channel=HuifengGuan)