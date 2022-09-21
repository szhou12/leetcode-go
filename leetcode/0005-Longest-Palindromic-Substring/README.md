# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## Solution idea

### 2-D DP

$DP[i][j] = $ whether (boolean) substring `s[i...j]` is a palindrome

Base Case 1:

$DP[i][j] = T$ if $i==j$

Base Case 2:

$DP[i][j] = F$ if $i>j$

Recurrence:

$DP[i][j] = T$ if ($s[i]==s[j]$ & j-i+1==2) OR (($s[i]==s[j]$ & $DP[i+1][j-1] == T$))

Result = the longest of `s[i...j]` among all $DP[i][j] == T$

Time complexity = $O(n^2)$
