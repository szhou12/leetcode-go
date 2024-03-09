# [3049. Earliest Second to Mark Indices II](https://leetcode.com/problems/earliest-second-to-mark-indices-ii/description/)

## Solution idea
### Binary Search 猜答案 + Regret Greedy (MinHeap)
1. 理解题意: 与[3048](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3048-Earliest-Second-to-Mark-Indices-I)不同的是，本题中mark操作也需要占用一单位时间。每一个时刻t可以做三件事情:
    1. decrement: 消耗一单位时间，挑选任意一个`nums[]`中一个非0元素减1
    2. set 0: 直接把`nums[changeIndices[t]]`一次性清0
    3. mark: 消耗一单位时间，挑选任意一个`nums[]`中一个0元素做马克

## Resource
[【每日一题】LeetCode 3049. Earliest Second to Mark Indices II](https://www.youtube.com/watch?v=VA6TLsOVMa4&ab_channel=HuifengGuan)