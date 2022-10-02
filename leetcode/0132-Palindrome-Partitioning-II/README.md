# [132. Palindrome Partitioning II](https://leetcode.com/problems/palindrome-partitioning-ii/)

## Solution idea

### DP
这是一道分两步走，并且每一步都可以用DP实现的题目。

1. 用dp来计算任意两个字符之间是否是回文的。转移方程是：
```
isPal[i][j] = 1 if isPal[i+1][j-1]==1 and s[i]==s[j]
            = 1 if i >= j (Base Cases)
            = 0 o/w
```

2. 用dp来计算第二个问题。令dp[i]表示从0到i的字符串可被拆分为最少的回文数的个数。则, `dp[i] = min(dp[i], dp[j]+1 for isPal[j+1][i]==1)`

Time complexity = $O(n^2)$

## Resource

[代码随想录-132. 分割回文串 II](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0132.%E5%88%86%E5%89%B2%E5%9B%9E%E6%96%87%E4%B8%B2II.md)