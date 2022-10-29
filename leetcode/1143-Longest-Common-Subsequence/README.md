# [1143. Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/)

## Solution idea

### DP - 经典题: LCS型

* Definition:
```
dp[i][j] = length of longest common subseq of text1[1..i] and text2[1...j]
```
    * NOTE: prepend text1, text2 with a placeholder so that characters starts at index 1 instead of 0
* Base Cases:
```
dp[0][0] = 0
dp[0][j] = 0 for all j
dp[i][0] = 0 for all i
```
* Recurrence:
```
dp[i][j] = dp[i-1][j-1] + 1             if text[i] == text[j]
         = max(dp[i-1][j], dp[i][j-1])  otherwise
```

Time complexity = $O(m*n)$ where $m$ length of text1, $n$ length of text2.

## Resource
[代码随想录-1143.最长公共子序列](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/1143.%E6%9C%80%E9%95%BF%E5%85%AC%E5%85%B1%E5%AD%90%E5%BA%8F%E5%88%97.md)