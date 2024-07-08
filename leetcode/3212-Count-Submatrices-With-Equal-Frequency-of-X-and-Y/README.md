# [3212. Count Submatrices With Equal Frequency of X and Y](https://leetcode.com/problems/count-submatrices-with-equal-frequency-of-x-and-y/description/)

## Solution idea
### 2-D Prefix Sum
1. Definition
    * `PrefixSumX[][]` and `PrefixSumY[][]` of shape `(m+1, n+1)`
        * i.e. `PrefixSum[i+1][j+1]` 对应 `grid[i][j]`
    * `PrefixSumX[i+1][j+1] :=` the number of `X` in the submatrix `grid[0][0] -> grid[i][j]`
    * `PrefixSumY[i+1][j+1] :=` the number of `Y` in the submatrix `grid[0][0] -> grid[i][j]`
    * 边界条件：`PrefixSumX[0][0:n+1] = 0`, `PrefixSumY[0:m+1][0] = 0` (0这里表示无意义，只为了方便recurrence计算)
2.  Recurrence: Venn diagram (inclusive-exclusive principle)
    1. `PrefixSumX[i+1][j+1] = (PrefixSumX[i][j+1] + PrefixSumX[i+1][j] - PrefixSumX[i][j]) + (grid[i][j] == 'X')`
    2. `PrefixSumY[i+1][j+1] = (PrefixSumY[i][j+1] + PrefixSumY[i+1][j] - PrefixSumY[i][j]) + (grid[i][j] == 'Y')`
3. `(0, 0) -> (i, j)` is a valid submatrix if:
    1. `PrefixSumX[i+1][j+1] > 0`: has at least one X
    2. `PrefixSumX[i+1][j+1] == PrefixSumY[i+1][j+1] > 0`

Time complexity = $O(m*n)$

### 一图胜千言
![AutoDraw Aug 7 2024](https://github.com/szhou12/leetcode-go/assets/35708194/5aea912c-1229-4805-b518-de03626afc1f)


## Resource
[Go / Python Solution | Beat 100% | Prefix Sum](https://leetcode.com/problems/count-submatrices-with-equal-frequency-of-x-and-y/solutions/5433966/go-python-solution-beat-100-prefix-sum/)
