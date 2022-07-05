# [2318. Number of Distinct Roll Sequences](https://leetcode.com/problems/number-of-distinct-roll-sequences/)

## Solution idea

### DFS
Since it asks for **total number** of distinct sequences possible, the most intuitive solution is DFS.

1) Number of levels = $n$ ($n$ = input)

2) For each level, there are at most $6$ branches because 6-sided dice

3) Pruning: At each level, continue DFS only if two required conditions are satisfied

Time complexity = $O(6^n)$. Clearly, too expensive.