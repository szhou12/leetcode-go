# [2370. Longest Ideal Subsequence](https://leetcode.com/problems/longest-ideal-subsequence/)

## Solution idea

### My solution - easy to come up, correct, BUT TLE!

Define $DP[i] = $ longest subsequence of `s[:i]` ending at `i`

(Prepend `s` with a placeholder `#` so that indices match with $DP$)

Base cases:

$DP[0] = $ empty string

$DP[1] = s[1]$

Recurrence:

$DP[i] = DP[j] + s[i]$ for $1\leq j \leq i-1$ such that $|s[i] - DP[j][len(DP[j]-1)]| \leq k$ and $\max_{1\leq j \leq i-1} len(DP[j])$

$DP[i]=s[i]$ otherwise.