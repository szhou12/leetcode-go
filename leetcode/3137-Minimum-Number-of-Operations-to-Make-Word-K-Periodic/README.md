# [3137. Minimum Number of Operations to Make Word K-Periodic](https://leetcode.com/problems/minimum-number-of-operations-to-make-word-k-periodic/description/)

## Solution idea
1. Any indices `i, j` to pick are divisible by `k`. In other words, `i, j` are multiples of `k`. Every substring has fixed length `k`. Thus, we use a map to record each substring of length `k` and its frequency. Find the substring of the highest frequency `m`. The minimum number of operations = `n/k - m`.

Time complexity = $O(n)$
