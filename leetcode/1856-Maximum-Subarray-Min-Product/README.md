# [1856. Maximum Subarray Min-Product](https://leetcode.com/problems/maximum-subarray-min-product/description/)

## Solution iea
### Stack - prevSmaller, nextSmaller + Prefix Sum
1. 解法如果在已经刷过[907](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0907-Sum-of-Subarray-Minimums)和[2104](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2104-Sum-of-Subarray-Ranges)的情况下，还是很直观的。
2. 题目关键词:
    1. subarray sum: 容易想到 prefix sum
    2. subarray minimum: 容易想到 使用递增stack寻找 prevSmaller, nextSmaller
    3. 求maximize: prevSmaller, nextSmaller 刚好拓展到subarray最远的边界，符合我们要最大化的要求。即，给定一个元素`nums[i]`，它能给出的最大值就是夹在`prevSmaller[i]`和`nextSmaller[i]`之间的subarray sum乘上`nums[i]`.
3. sanity check：注意prevSmaller可能越界给出-1，所以prefix sum需要sanity check

Time complexity = $O(n)$

## Resource
[【每日一题】1856. Maximum Subarray Min-Product, 5/12/2021](https://www.youtube.com/watch?v=sLYUi_0HOBM&ab_channel=HuifengGuan)