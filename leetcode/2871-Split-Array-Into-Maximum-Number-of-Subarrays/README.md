# [2871. Split Array Into Maximum Number of Subarrays](https://leetcode.com/problems/split-array-into-maximum-number-of-subarrays/description/)

## Solution idea
### Greedy + Bit Operation
1. AND 一个重要性质: AND越多元素，结果的值越小。
    * 因为只要一个bit位上出现0，AND就会使这个bit位永远为0。AND 的元素越多使得每个bit位上出现0的概率越大。
2. 应用AND的性质在这道题上：所有元素全部 AND 起来，就是可以得到的最小值。
3. 但是，不一定必须 AND 所有的元素才可以得到最小值。如果一个subarray的元素AND起来得到0，那么剩下的元素也可以得到最小值。这样就给split array提供了契机。i.e., 0 + 0 + ... + 0 + min = min
4. 解法：根据上面的思路，split的过程就是**贪心地**寻找 AND 起来得0的subarray的过程。最不济的，就是符合这个条件的subarray一个都没找到，那么，整个array就是唯一的subarray。

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 2871. Split Array Into Maximum Number of Subarrays](https://www.youtube.com/watch?v=IKoy5gQltE4&list=PLwdV8xC1EWHrtgsYCcDTXIMVaHSlsnLzL&index=20)

