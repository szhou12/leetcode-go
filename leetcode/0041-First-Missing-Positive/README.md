# [41. First Missing Positive](https://leetcode.com/problems/first-missing-positive/description/)

## Solution idea
1. for any array whose length is `n`, the smallest missing positive must be in range `[1,...,n+1]`, so we only have to care about those elements in this range and remove the rest.

## Resource
[Golang concise solution with some explanation](https://leetcode.com/problems/first-missing-positive/solutions/17157/golang-concise-solution-with-some-explanation)