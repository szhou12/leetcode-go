# [3116. Kth Smallest Amount With Single Denomination Combination](https://leetcode.com/problems/kth-smallest-amount-with-single-denomination-combination/description/)

## Solution idea
### Binary Search (Guess k) + Grosper's Hack (Combinatorics) + Exclusion-Inclusion Principle + Divisors
1. 理解题意: 将所有组合数 $a_0coins[0] + a_1coins[1] + \cdots + a_{n-1}coins[n-1]$ where $a_0, a_1, \cdots, a_{n-1} = 0\cdots \infty$ and not all $a_i = 0$ at the same time 从小到大排列，找到第 $k$ 个数。


## Resrouce
[【每日一题】LeetCode 3116. Kth Smallest Amount With Single Denomination Combination](https://www.youtube.com/watch?v=R63BH6UDQXQ&ab_channel=HuifengGuan)