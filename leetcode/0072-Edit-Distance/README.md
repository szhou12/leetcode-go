# [72. Edit Distance](https://leetcode.com/problems/edit-distance/)

## Solution idea

## DP

**经典中的经典DP**

$DP[i][j]$ = min cost to transform first `i` chars of `word1` to first `j` chars of `word2`

e.g. $DP[0][2]$ = `word1` is empty string (0 char), `word2` has 2 chars

Base Cases:

* $DP[0][0] = 0$

* $DP[i][0] = i$

* $DP[0][j] = j$

Recurrence:

$DP[i][j] = DP[i-1][j-1]$ if `word1[i]` == `word2[j]`

$DP[i][j] = \min\{RC, IC, DC\}$ otherwise

* RC: replace cost = $d(w1_i, w2_j) + DP[i-1][j-1]$ ($d(w1_i, w2_j)$ is the cost to replace ith char of w1 with jth char of w2)

* DC : delete cost = $d(w1_i, _) + DP[i-1][j]$ ($d(w1_i, _)$ is the cost to delete ith char of w1)

* IC: insert cost = $d(_, w2_j) + DP[i][j-1]$ ($d(_, w2_j)$ is the cost to insert jth char of w2)


**Key 1**: prepend '#' as placeholder. 意义在于，使得 `dp[i][j]` 中 `i/j` 为0时，分别代表各自的string为 empty string. 如果不这样做，`dp[0][j]` 中对 `word1`的表述是代表了 `word1` 第一个字母，而不是`word1`为空串的情况，这样`dp`的index与实际 `word1`, `word2`中的index 产生了偏移 (相差1)， 实际coding中容易犯迷糊，所以，前置一个占位符.

Time complexity = $O(m* n)$