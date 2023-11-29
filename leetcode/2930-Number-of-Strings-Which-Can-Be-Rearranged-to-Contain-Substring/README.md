# [2930. Number of Strings Which Can Be Rearranged to Contain Substring](https://leetcode.com/problems/number-of-strings-which-can-be-rearranged-to-contain-substring/description/)

## Solution idea
* **Key Idea**: 大多数数学中的组合问题都是通过 DP 来推导得到的
    * 组合公式: `C(n, m) = C(n-1, m) + C(n-1, m-1)`
    * 转换成 DP: `dp[n][m] = dp[n-1][m] + dp[n-1][m-1]`
    * 数学意义：n中选m个, 要问的是: 第n个取不取？如果不取，就转化成从前n-1个中取m个；如果取，就转化成选定第n个，再从前n-1个中取m-1个
* 解题思路：在由 n 个字母组成的字符串中，挖掉4个位子放入 'l', 'e', 'e'. 't' (order doesn't matter)，有多少种组合的可能？

Time Complexity = $O(n)$

## Resource
[每日一题】LeetCode 2930. Number of Strings Which Can Be Rearranged to Contain Substring](https://www.youtube.com/watch?v=0V95_GZH6DM&ab_channel=HuifengGuan)