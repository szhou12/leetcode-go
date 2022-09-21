# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## Solution idea

### 2-D DP

$DP[i][j] = $ whether (boolean) substring `s[i...j]` is a palindrome

Base Case 1:

$DP[i][j] = T$ if $i==j$

Base Case 2:

$DP[i][j]$ = F$ if $i>j$

Recurrence:

$DP[i][j] = T$ if ($s[i]==s[j]$ & j-i+1==2) OR ($s[i]==s[j]$ & $DP[i+1][j-1] == T$)

Result = the longest of `s[i...j]` among all $DP[i][j] == T$

Time complexity = $O(n^2)$

### Two Pointers

**核心思路**：从**中心**向两端扩散的双指针技巧

**中心** 有两种情况：1) 以每个字母为中心 (e.g. `...a...`)；2）以连续的两个字母为中心 (e.g. `...aa...`)

即：如果回文串的长度为奇数，则它有一个中心字符；如果回文串的长度为偶数，则可以认为它有两个中心字符。

Time complexity = $O(n^2)$
