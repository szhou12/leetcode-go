# [712. Minimum ASCII Delete Sum for Two Strings](https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/)

## Solution idea

### DP - LCS 类型

* Definition:
```
dp[i][j] = min sum of deleted chars to make s1[1...i] == s2[1...j]
```
    * Note: prepend s1, s2 with a placeholder so that meaningful charater of s1, s2 start at index 1
* Base cases:
```
dp[0][0] = 0
dp[i][0] = sum(s1[i]) for all i
dp[0][j] = sum(s2[j]) for all j
```
* Recurrence:
```
dp[i][j] = dp[i-1][j-1]                                                      if s1[i] == s2[j]
         = min(dp[i-1][j]+s1[i], dp[i][j-1]+s2[j], dp[i-1][j-1]+s1[i]+s2[j]) o/w
```

Time complexity = $O(m*n)$ where $m$ length of s1, $n$ length of s2.