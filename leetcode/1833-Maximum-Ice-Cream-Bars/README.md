# [1833. Maximum Ice Cream Bars](https://leetcode.com/problems/maximum-ice-cream-bars/description/)

## Solution idea
### Greedy Algorithm
We want to buy as many ice cream bars as possible with limited coins. Therefore, it is always optimal to buy the cheapest ice cream bar first. So greedily, we sort the ice cream bars in ascending order and keep buying them until we run out of coins.

Time complexity = $O(n\log n)$