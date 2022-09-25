# [63. Unique Paths II](https://leetcode.com/problems/unique-paths-ii/)

## Solution idea

### DP

思路和[62. Unique Paths](https://leetcode.com/problems/unique-paths/)基本上一样，唯一区别是，如果某格子是obstacle, 那么能走到这个格子的方法数是 0.

同时，**注意！！！** Base case 如何设值!

如果某格子是obstacle, 那么这个格子和之后所有格子都无法到达。

Time complexity = $O(m*n)$