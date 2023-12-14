# [2871. Split Array Into Maximum Number of Subarrays](https://leetcode.com/problems/split-array-into-maximum-number-of-subarrays/description/)

## Solution idea
### Greedy + Bit Operation
1. AND 一个重要性质: AND越多元素，结果的值越小。
    * 因为只要一个bit位上出现0，AND就会使这个bit位永远为0。AND 的元素越多使得每个bit位上出现0的概率越大。
2. 应用 AND的性质在这道题上：所有元素全部 AND 起来，就是可以得到的最小值。

## Resource
[【每日一题】LeetCode 2871. Split Array Into Maximum Number of Subarrays](https://www.youtube.com/watch?v=IKoy5gQltE4&list=PLwdV8xC1EWHrtgsYCcDTXIMVaHSlsnLzL&index=20)

