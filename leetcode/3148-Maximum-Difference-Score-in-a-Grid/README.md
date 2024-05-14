# [3148. Maximum Difference Score in a Grid](https://leetcode.com/problems/maximum-difference-score-in-a-grid/description/)

## Solution idea
### 2-D DP Min-Max Strategy
1. 突破点: For simplicity, let $c_k = (i, j)$. The score of any path that passes through $c_1, c_2, \cdots, c_k$ can be represented as $(c_k - c_{k-1}) + (c_{k-1}-c_{k-2}) + \cdots + (c_3-c_2) + (c_2-c_1) = c_k-c_1$. Therefore, the score of any path depends ONLY on the start and end cells. Since a move can only go right or do down, if the end cell is $(i, j)$, we only need to find the smallest cell in its top-left area (the rectangle of `grid[0...i][0...j]` not including $(i, j)$ itself because we must make one move) as the start cell.
2. 2-D DP:
    * Define `dp[i][j] = ` minimum value in the top-left area of $(i, j)$.
    * Recurrence: `dp[i][j] = min(dp[i-1][j], dp[i][j-1], grid[i][j])`
    * NOTE: CANNOT set base case as `dp[0][0] = grid[0][0]` before update the result because we must make one move.

Time complexity = $O(m*n)$

## Resource
* [[Java/C++/Python] DP, Minimum on Top-Left](https://leetcode.com/problems/maximum-difference-score-in-a-grid/solutions/5145704/java-c-python-dp-minimum-on-top-left/)
* [[C++] tricky problem, but simple solution DP](https://leetcode.com/problems/maximum-difference-score-in-a-grid/solutions/5145625/c-tricky-problem-but-simple-solution-dp/)
    * Its figures are helpful.