# [2911. Minimum Changes to Make K Semi-palindromes](https://leetcode.com/problems/minimum-changes-to-make-k-semi-palindromes/description/)

## Solution idea
### DP
#### 思路总结
1. DP 典型类型题 - 切绳子: 一个大字符串，分成 k 个区间，要求每个区间满足一些条件，求整体的一个最优规划。
    1. Definition: `dp[i][p]` = minimun number of changes in make `s[0:i]` (inclusive) so it has **p** semi-palindrome substrings.
        * 怎么回头看? `X X X j [X X i]` 前 j 个元素分成 p-1 个区间所需最小变动 (for all `j < i`) + `[X X i]` 当前这段区间所需最小变动
```go
// Recurrence 模版
for i := 0; i < n; i++ {
    for p := 2; p <= k; p++ { // 注意：这里 p 从2开始。p=1 表示一整个大区间，不做任何分割。p=0 没有意义。
        dp[i][p] = math.MaxInt / 2
        for j := 0; j < i; j++ {
            dp[i][p] = min(dp[i][p], dp[j][p-1]+cost[j+1][i])
        }
    }
}
return dp[n-1][k]
```

2. 啥是 semi-palindrome?
    1. 数学描述: A string with a length of `len` is considered a **semi-palindrome** if there exists a positive integer `d` such that `1 <= d < len` and `len % d == 0`, and if we take indices that have the same modulo by `d`, they form a palindrome.
    2. 例子:
        * `"adbgad"` is a **semi-palindrome**
        ```
           0 1 2 3 4 5
           a d b g a d (len = 6)
        %1 0 0 0 0 0 0 (abdgad)
        %2 0 1 0 1 0 1 (aba, dgd) <- semi-palindrome
           | | | | | |
           a | b | a |
             d   g   d
        %3 0 1 2 0 1 2 (ag, da, bd)
        %4 0 1 2 3 0 1 (aa, dd, b, g) <- invalid! 6%4 != 0
        %5 0 1 2 3 4 5 (a, d, b, g, a, d) <- invalid! 6%5 != 0
        ```

Time Complexity = $O(n^3 + n^3) = O(n^3)$
## Resource
[【每日一题】LeetCode 2911. Minimum Changes to Make K Semi-palindromes](https://www.youtube.com/watch?v=We7N0gmmkLg&ab_channel=HuifengGuan)