# [647. Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/)

## Solution idea

### DP

思路与其他回文串的题，尤其是[5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)很像。每找到一个合法的回文串，就在 `res` 上 `+1`.

Time complexity = $O(n^2)$

Space complexity = $O(n^2)$

### Two Pointers

**中心开花**

* 找以`[i]`为中心的合法回文串有多少个

* 找以`[i, i+1]`为中心的合法回文串有多少个

* 把所有的个数加起来

Time complexity = $O(n^2)$

Space complexity = $O(1)$

## Resource

[代码随想录-647. 回文子串](https://programmercarl.com/0647.%E5%9B%9E%E6%96%87%E5%AD%90%E4%B8%B2.html)