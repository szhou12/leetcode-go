# [51. N-Queens](https://leetcode.com/problems/n-queens/)

## Solution idea
Pruning: check if the column to place a new Queen is:

1. vertical to any previous Queens' column; 

2. diagonal to any previous Queens (slope == 1)

3. reversely diagonal to any previous Queens (slope == -1)

Time = $n\times T(n-1) + O(n)$ where O(n) is for check valid = $O(n!)$ OR $O(n^n)$ ???

Space = $O(n)$