# [70. Climbing Stairs]()

## Solution idea

### DP
Define $DP[i] = $ number of distinct ways to climb to $i$ 

Recurrence:

$DP[i] = DP[i-2] + DP[i-1]$

Base cases:

$DP[0] = 1$, $DP[1] = 1$

Time complexity = $O(n)$

This is Fibonacci Sequence...