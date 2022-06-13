# [377. Combination Sum IV](https://leetcode.com/problems/combination-sum-iv/)

## Solution Idea:

This problem looks like a DFS problem but is not easily solvable by DFS.

It's actually a DP problem in disguise.

Let $DP(i)$ be the number of combination ways that sum up to $i$.

Define the recurrence:

Base case: $DP(0) = 1$ since one way (empty array) will sum up to $0$.

Ind. case: $DP(i) = \sum_{1\leq j \leq len(A), i-A[j]\geq 0}DP(i-A[j])$
