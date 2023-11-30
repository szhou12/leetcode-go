# [2930. Number of Strings Which Can Be Rearranged to Contain Substring](https://leetcode.com/problems/number-of-strings-which-can-be-rearranged-to-contain-substring/description/)

## Solution idea
### DP
#### Key Idea: 大多数数学中的组合问题都是通过 DP 来推导得到的
* 组合公式: `C(n, m) = C(n-1, m) + C(n-1, m-1)`
* 转换成 DP 等价于: `dp[n][m] = dp[n-1][m] + dp[n-1][m-1]`
* 数学意义：n中选m个, 要问的是: 第n个取不取？如果不取，就转化成从前n-1个中取m个；如果取，就转化成选定第n个，再从前n-1个中取m-1个
#### 解题思路：在由 n 个字母组成的字符串中，挖掉4个位子放入 'l', 'e', 'e', 't' (order doesn't matter)，有多少种组合的可能？
#### DP Solution
1. Definition: `DP[i][a][b][c] =` number of strings of i letters that have at least `a` number of 'l', `b` number of 'e', `c` number of 't' where `a = 0, 1`, `b = 0, 1, 2`, `c = 0, 1`.
2. Base Case: `DP[0][0][0][0] = 1` 满足条件的 (没有 'l', 'e', 't') 空串有且只有一个
3. Recurrence: 应用公式 `C(n, m) = C(n-1, m) + C(n-1, m-1)`
    * `DP[i][1][b][c] += DP[i-1][0][b][c]` 如果第i个字母是 'l'，且取第i个使 `DP[i]` 包含至少1个'l' (`a==1`)
    * `DP[i][a][1][c] += DP[i-1][a][0][c]` 如果第i个字母是 'e'，且取第i个使 `DP[i]` 包含至少1个'e' (`b==1`)
    * `DP[i][a][2][c] += DP[i-1][a][1][c]` 如果第i个字母是 'e'，且取第i个使 `DP[i]` 包含至少2个'e' (`b==2`)
    * `DP[i][a][b][1] += DP[i-1][a][b][0]` 如果第i个字母是 't'，且取第i个使 `DP[i]` 包含至少1个't' (`c==1`)
    * `DP[i][a][b][c] += DP[i-1][a][b][c]` 所有其他的情况, 不取第i个 == `C(n-1, m)`
### Math: PIE (inclusion-exclusion)
#### Intuition
$A =$ the string has $\leq 1$ 'e'

$B =$ the string has no 'l'

$C =$ the string has no 't'

$|A\cup B\cup C| =$ the number of strings that have <= 1 'e' or no 'l' or no 't' = the total number of BAD strings

By PIE (inclusion-exclusion), $|A\cup B\cup C| = |A| + |B| + |C| - (|A\cap B|+|A\cap C|+|B\cap C|) + |A\cap B \cap C|$

$|A| = n\cdot 25^{n-1} + 25^n$
1. no 'e': $25^n$
2. exactly one 'e': ${n \choose 1} \cdot 25^{n-1}$ (选一个位子放'e'，其余位子放除了'e'任意字母)

$|B| = 25^n$

$|C| = 25^n$

$|B\cap C| = 24^n$

$|A\cap B| = n\cdot 24^{n-1} + 24^n$
1. no 'e' + no 'l': $24^n$
2. exactly one 'e' + no 'l': ${n \choose 1} \cdot 24^{n-1}$ (选一个位子放'e'，其余位子放除了'e'和'l'任意字母)

$|A\cap C| = n\cdot 24^{n-1} + 24^n$
1. no 'e' + no 't': $24^n$
2. exactly one 'e' + no 't': ${n \choose 1} \cdot 24^{n-1}$ (选一个位子放'e'，其余位子放除了'e'和't'任意字母)

$|A\cap B\cap C| = n\cdot 23^{n-1} + 23^n$
1. no 'e' + no 'l' + no 't': $23^n$
2. exactly one 'e' + no 'l' + no 't': ${n \choose 1} \cdot 23^{n-1}$ (选一个位子放'e'，其余位子放除了'e'、'l'、't'任意字母)

Good strings = Total strings - Bad strings

i.e., $26^n - [|A\cup B\cup C| = |A| + |B| + |C| - (|A\cap B|+|A\cap C|+|B\cap C|) + |A\cap B \cap C|]$

* :exclamation: 编者注:
    * DP 的code确实能通过测试
    * 我也有尝试推导本题的数学表达式，得出的是: $C(n, 4) \times \frac{P(4,4)}{P(2,2)} \times 26^{n-4} = C(n, 4) \times 12 \times 26^{n-4}$
    * 即, 从n个中挑出4个位子放入 'l', 'e', 'e', 't'进行排列，其余位子可以放入任意26个字母中的一个
    * 但是！这个表达式的计算结果对导致重复计算。例如: n=5时, 表达式得到 1560，而 DP 得到 1460。当 n > 4时，表达式的结果总会大于 DP 的结果
    * 网上的参考答案有从逆向角度来解题，使用 PIE (inclusion-exclusion)，虽然繁琐，但可以避免重复计算

Time Complexity = $O(2\times 3\times 2\times 26\times n) = O(n)$

## Resource
- DP: [每日一题】LeetCode 2930. Number of Strings Which Can Be Rearranged to Contain Substring](https://www.youtube.com/watch?v=0V95_GZH6DM&ab_channel=HuifengGuan)
- PIE (inclusion-exclusion): [Inclusion and exclusion principle, detailed explanation, faster than 100%](https://leetcode.com/problems/number-of-strings-which-can-be-rearranged-to-contain-substring/solutions/4277290/inclusion-and-exclusion-principle-detailed-explanation-faster-than-100/)