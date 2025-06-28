# [36. Valid Sudoku](https://leetcode.com/problems/valid-sudoku/description/)

## Solution idea
### Hash Map + Matrix Calculation
1. Matrix calculation: segregate the NxN board into 3x3 sub-boxes. Sub-boxes are indexed by `(r/3)*3 + (c/3)` 

Time complexity = $O(N^2)$

Space complexity = $O(N^2)$