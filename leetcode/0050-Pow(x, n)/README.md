# [50. Pow(x, n)](https://leetcode.com/problems/powx-n/)

## Solution idea

简单的 Recursion. 注意：提前算好 `pow(x, n/2)`, 避免重复计算.

Time complexity = $O(\log n)$