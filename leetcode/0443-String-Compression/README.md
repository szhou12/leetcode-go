# [443. String Compression](https://leetcode.com/problems/string-compression/)

## Solution idea

### Two Pointers

细节**非常多**的一道双指针问题

* slow物理意义：slow挡板左边不包括slow 表示compressed string
* 遇到相同字母时, fast移动, slow不动, 并计算遇到的相同字母个数
* 遇到不同字母时：
    1. 把当时在看的字母 抛给 slow, slow走一格
    2. 把字母个数append到后面挨着 (count==1 则跳过)，slow走count多个格子. 此时，slow和fast在同一个格子

Time complexity = $O(n^2)$