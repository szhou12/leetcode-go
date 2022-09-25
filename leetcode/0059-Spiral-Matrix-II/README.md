# [59. Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/)

## Solution idea

从外到里一层一层地填，每一层填的顺序为上->右->下->左. **注意**每一次填完，下一次避开上一次填的最后一个格子

**Tip**: 对照着题目给的图写code会比较清晰要规避哪些corner case

Time complexity = O(n*n)

## Resource
[代码随想录 - 59.螺旋矩阵II](https://www.programmercarl.com/0059.%E8%9E%BA%E6%97%8B%E7%9F%A9%E9%98%B5II.html)