# [2024. Maximize the Confusion of an Exam](https://leetcode.com/problems/maximize-the-confusion-of-an-exam/description/)

## Solution idea

### Two Pointers
* **模版Flex**: 写法与 [1004. Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/description/) 一样, 固定左边界, 右边界不断延伸至最长. 唯一不一样是本题需要 two passes, 也就是分别求 变成 T 的最长 和 变成 F 的最长, 然后取较大的.

Time complexity = $O(n)$

## Resource
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Two_Pointers/2024.Maximize-the-Confusion-of-an-Exam)