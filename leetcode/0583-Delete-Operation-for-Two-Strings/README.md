# [583. Delete Operation for Two Strings](https://leetcode.com/problems/delete-operation-for-two-strings/)

## Solution idea

### DP method 1

* Definition:
```
DP[i][j] = min # of deletions needed to make word1[1...i] identical to word2[1...j]
```
    * Note: word1[0], word2[0] placed with a placeholder
* Base Cases:
```
DP[0][0] = 0
DP[i][0] = i for 1 <= i <= m
DP[0][j] = j for 1 <= j <= n
```
* Recurrence:
```
DP[i][j] = DP[i-1][j-1] if word1[i] == word2[j]
         = min(DP[i-1][j-1]+2, DP[i-1][j]+1, DP[i][j-1]+1) otherwise
```
    * 解释：
        * 当`word1[i]` 与 `word2[j]`不相同的时候，有三种情况：
            1. 删`word1[i]`，最少操作次数为`DP[i - 1][j] + 1`
            2. 删`word2[j]`，最少操作次数为`DP[i][j - 1] + 1`
            3. 同时删`word1[i]`和`word2[j]`，操作的最少次数为 `DP[i - 1][j - 1] + 2`

Time complexity = $O(m*n)$


### DP method 2 - Longest Common Subsequence
* [1143. Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/)的变形题
    * 求出LCS, 然后减去LCS长度即为需要删除的次数

## Resource
[代码随想录-583. 两个字符串的删除操作](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0583.%E4%B8%A4%E4%B8%AA%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%9A%84%E5%88%A0%E9%99%A4%E6%93%8D%E4%BD%9C.md)
