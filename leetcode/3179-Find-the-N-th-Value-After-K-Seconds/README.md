# [3179. Find the N-th Value After K Seconds](https://leetcode.com/problems/find-the-n-th-value-after-k-seconds/description/)

## Solution idea
### Prefix Sum
1. 以`[1, 1, ..., 1]`作为种子，算prefixSum算k次
Time complexity = $O(k\cdot n)$

### Pascal's Triangle
1. Pascal's Triangle is a graphical representation of the binomial coefficients.
2. The binomial coefficient $\binom{n}{k}$, in Pascal's Triangle, is the element at row $n$ and position $k$ (starting from 0).
3. The value at position (n−1) after k seconds can be represented in Pascal's Triangle as: $\binom{n+k-1}{k}$ (why?)
Time complexity = $O(\binom{n+k-1}{k})$
