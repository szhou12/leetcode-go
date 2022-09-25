# [62. Unique Paths](https://leetcode.com/problems/unique-paths/)

## Solution idea

### DFS 暴力解 - 超时!!!

每一格子有向下、向右两种，一共有 $m \times n$ 格子,

Time complexity = $O(2^{m+n})$

### DP

**注意！！！** Base case 如何设值!

$DP[i][j]$ = number of unique paths from `grid[0][0]` to `grid[i][j]`

Base Cases:

$DP[0][0] = 1$ 站在原点只有一种path

$DP[0][j] = 1$ 当棋盘只有一行，到达任意格子都只有不停向右走一种path

$DP[i][0] = 1$ 当棋盘只有一列，到达任意格子都只有不停向下走一种path

Recurrence:

$DP[i][j]= DP[i-1][j] + DP[i][j-1]$ 当前格子由前一个格子向右走来，或者向下走来

Time complexity = $O(m*n)$

## Resource

[贾考博 LeetCode 62. Unique Paths](https://www.youtube.com/watch?v=L6dWXuh8BuE&ab_channel=%E8%B4%BE%E8%80%83%E5%8D%9A)