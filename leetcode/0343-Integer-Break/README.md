# [343. Integer Break](https://leetcode.com/problems/integer-break/)

## Solution idea:

Define $DP[i]$ the max product of breaking integer $i$

Base cases:

$DP[0] = 0, DP[1] = 1$

Recurrence:

$DP[i] = \max (j\times (i-j), j\times DP[i-j])$ for $1\leq j < i$