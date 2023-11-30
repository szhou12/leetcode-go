# [2930. Number of Strings Which Can Be Rearranged to Contain Substring](https://leetcode.com/problems/number-of-strings-which-can-be-rearranged-to-contain-substring/description/)

## Solution idea
* **Key Idea**: 大多数数学中的组合问题都是通过 DP 来推导得到的
    * 组合公式: `C(n, m) = C(n-1, m) + C(n-1, m-1)`
    * 转换成 DP 等价于: `dp[n][m] = dp[n-1][m] + dp[n-1][m-1]`
    * 数学意义：n中选m个, 要问的是: 第n个取不取？如果不取，就转化成从前n-1个中取m个；如果取，就转化成选定第n个，再从前n-1个中取m-1个
* 解题思路：在由 n 个字母组成的字符串中，挖掉4个位子放入 'l', 'e', 'e', 't' (order doesn't matter)，有多少种组合的可能？
* DP
    1. Definition: `DP[i][a][b][c] =` number of strings of i letters that have at least `a` number of 'l', `b` number of 'e', `c` number of 't' where `a = 0, 1`, `b = 0, 1, 2`, `c = 0, 1`.
    2. Base Case: `DP[0][0][0][0] = 1` 满足条件的 (没有 'l', 'e', 't') 空串有且只有一个
    3. Recurrence: 应用公式 `C(n, m) = C(n-1, m) + C(n-1, m-1)`
        * `DP[i][1][b][c] += DP[i-1][0][b][c]` 如果第i个字母是 'l'，且取第i个使 `DP[i]` 包含至少1个'l' (`a==1`)
        * `DP[i][a][1][c] += DP[i-1][a][0][c]` 如果第i个字母是 'e'，且取第i个使 `DP[i]` 包含至少1个'e' (`b==1`)
        * `DP[i][a][2][c] += DP[i-1][a][1][c]` 如果第i个字母是 'e'，且取第i个使 `DP[i]` 包含至少2个'e' (`b==2`)
        * `DP[i][a][b][1] += DP[i-1][a][b][0]` 如果第i个字母是 't'，且取第i个使 `DP[i]` 包含至少1个't' (`c==1`)
        * `DP[i][a][b][c] += DP[i-1][a][b][c]` 所有其他的情况, 不取第i个 == `C(n-1, m)`
* 编者注:
    * 我有尝试推导本题的数学表达式，得出的是: $C(n, 4) \times \frac{P(4,4)}{P(2,2)} \times 26^{n-4} = C(n, 4) \times 12 \times 26^{n-4}$
    * 即, 从n个中挑出4个位子放入 'l', 'e', 'e', 't'进行排列，其余位子可以放入任意26个字母中的一个
    * 但是！这个表达式的计算结果与 DP 的结果不一致。例如: n=5时, 表达式得到 1560，而 DP 得到 1460
    * 目前并没有明白为什么会算的不同。如果对 1460 进行分解质因数得: $1460 = 2^2 \times 5 \times 73$，并不是 12 (leet的所有排列数) 的倍数
    * 网上的参考答案有从逆向角度来解题，使用 PIE (inclusion-exclusion)，目前并没有完全理解

Time Complexity = $O(2*3*2*26*n) = O(n)$

## Resource
[每日一题】LeetCode 2930. Number of Strings Which Can Be Rearranged to Contain Substring](https://www.youtube.com/watch?v=0V95_GZH6DM&ab_channel=HuifengGuan)