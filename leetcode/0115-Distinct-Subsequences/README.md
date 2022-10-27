# [115. Distinct Subsequences](https://leetcode.com/problems/distinct-subsequences/)

## Solution idea

## DP

$DP[i][j] = $ # of distinct subseq of `s[1...i]` = `t[1...j]`

Base Cases:

* $DP[0][0] = 1$

* $DP[i][0] = 1$ 因为空串永远是`s[1...i]`的一个subsequence

* $DP[0][j] = 0$

Recurrence:

```
dp[i][j] = dp[i-1][j-1] + dp[i-1][j] if s[i] == t[j]
         = dp[i-1][j]                if s[i] != t[j]
```

**难点!!!** 
1. 当 `s[i] == t[j]`时，`dp[i-1][j-1]`比较容易想到，`dp[i-1][j]`难以想到。为什么要考虑用 `s[1...i-1]` 来匹配 `t[1...j]`? 举个例子：`s = bagg`, `t = bag`, 此时, 我们可以用`s`的末位 g 来匹配; 也可以不用, 用倒数第二位的 g 来匹配.

2. 当 `s[i] != t[j]`时, 显然，此时`s`的末位无法用来匹配, 那么只能回头看，尝试用 `s[1...i-1]` 来匹配 `t[1...j]`

Time complexity = $O(m*n)$ where $m$ is length of `s`, $n$ is length of `t`

## Resource

[代码随想录-115.不同的子序列](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0115.%E4%B8%8D%E5%90%8C%E7%9A%84%E5%AD%90%E5%BA%8F%E5%88%97.md)