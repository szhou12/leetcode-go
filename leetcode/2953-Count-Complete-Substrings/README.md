# [2953. Count Complete Substrings](https://leetcode.com/problems/count-complete-substrings/description/)

## Solution idea
### Sliding Window: Flex + Fix

Time complexity = $O(n*26*26)$

Explanation: The whole array is splited into `m` chunks (each chunk has max size of `n` while the total size of all chunks together is `n`). Each chunk iterates at most 26 times of window size while each window size iterates at most 26 times to check each character's frequency.

## Recource
[【每日一题】LeetCode 2953. Count Complete Substrings](https://www.youtube.com/watch?v=4DYbP4shsno&ab_channel=HuifengGuan)