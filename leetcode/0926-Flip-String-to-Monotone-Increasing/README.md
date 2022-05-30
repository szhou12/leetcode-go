# [926. Flip String to Monotone Increasing](https://leetcode.com/problems/flip-string-to-monotone-increasing/)

## Idea: DP

DP[i] := minimal number of flips ending at i-th element

Base case:

DP[0] = 0

Recurrence:

case 1: if s[i] == 1, then no need to flip i-th element, DP[i] same as DP[i-1]

case 2: if s[i] == 0, then either flip current element and inherit DP[i-1] or flip all previous zeros

$DP[i] = DP[i-1] if s[i] == 1$

= min(DP[i-1] + 1, previous ones) if s[i] == 0