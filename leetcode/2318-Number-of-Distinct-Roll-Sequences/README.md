# [2318. Number of Distinct Roll Sequences](https://leetcode.com/problems/number-of-distinct-roll-sequences/)

## Solution idea

### DFS
Since it asks for **total number** of distinct sequences possible, the most intuitive solution is DFS.

1) Number of levels = $n$ ($n$ = input)

2) For each level, there are at most $6$ branches because 6-sided dice

3) Pruning: At each level, continue DFS only if two required conditions are satisfied

Time complexity = $O(6^n)$. Clearly, too expensive.

### 3-D DP (Optimal Solution)
Define $DP[i][b][a]$ as total number of possible sequences when rolling $i$ times where the $i-1$-th value is $b$ and $i$-th value is $a$. (`X X X X b a`)

Reccurence:

$DP[i][b][a] = \sum DP[i-1][c][b]$ where $1\leq a,b,c \leq 6$ and $gcd(a,b)=1$ and $a\neq b$ and $a\neq c$.

Base cases:

if $n=1$, only $6$ possible ways to roll a dice once.

if $n=2$, all pairs $(x,y)$ where $x \neq y$ and $gcd(x,y) == 1$.

Result:

$\sum DP[n][X][X]$ where $1\leq X \leq 6$.

Time complexity = $O(6^3\times n)$

### WHY 3-D DP?
`DP[i]` 表示 扔 i 次 的所有可行方法。 但好像不够，怎么知道第i次扔到的数能不能满足条件呢？`DP[i]` and `DP[i-1]` 没法确定是不是合法的

那么，这样呢？

`DP[i][a]` 表示 扔 i 次而且第i次扔到a 的所有可行方法, 这样我们就可以显性地建立 i-1 和 i 次的联系: `DP[i][a]` and `DP[i-1][b]`

但是，题目还要求第i次扔到的数和第i-2次扔到的数不能相等。似乎，二维DP又不够用了。那么，三维呢？

`DP[i][b][a]` 表示 扔 i 次 而且 第i-1次扔到b & 第i次扔到a 的所有可行方法, 这样我们就可以显性地建立 i-1 和 i 次的联系 又能 考虑到 i-2 和 i的要求： `DP[i][b][a]` and `DP[i-1][c][b]`.

### Resource
[【每日一题】LeetCode 2318. Number of Distinct Roll Sequences](https://www.youtube.com/watch?v=7SZBLIga_-s&ab_channel=HuifengGuan)