# [456. 132 Pattern](https://leetcode.com/problems/132-pattern/)

## Solution idea

Key point: $A[j]$

For each $A[j]$:

1. want to find the smallest element on its left side in $A[0\cdots j-1]$

$\Rightarrow$ loop forward + memoization

2. want to find the largest element < $A[j]$ on its right side in $A[j+1 \cdots n]$

$\Rightarrow $ loop backward + monotonic stack

So that: $A[i] < A[k] < A[j]$ for $i<j<k$