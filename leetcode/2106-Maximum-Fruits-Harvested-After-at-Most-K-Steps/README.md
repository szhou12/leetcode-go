# [2106. Maximum Fruits Harvested After at Most K Steps](https://leetcode.com/problems/maximum-fruits-harvested-after-at-most-k-steps/description/)

## Solution idea
### Sliding Window (非模版)
#### 思路总结
1. 首先需要明确一条规律: 调头只会调一次。因为如果需要第二次调头，完全可以第一次调头的时候就可以把第二次调头的操作也做了。
3. 求区间 `[A, B]` 的总数 => 前缀和: `prefix sum([0, B]) - prefix sum([0, A])`


## Resource
[【每日一题】LeetCode 2106. Maximum Fruits Harvested After at Most K Steps](https://www.youtube.com/watch?v=X3Mp8SAU9gI&ab_channel=HuifengGuan)