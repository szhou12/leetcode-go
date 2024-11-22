# [2193. Minimum Number of Moves to Make Palindrome](https://leetcode.com/problems/minimum-number-of-moves-to-make-palindrome/description/)

## Solution idea
### Greedy + Simulation
1. 构想机制：有没有一个系统的方法进行转换操作？
    - 从外往内，找一个配对，“拉到”最外层，“消掉”。重复此过程，像“剥洋葱”一样。

Time complexity = $O(n^2)$

## Resource
- [【每日一题】LeetCode 2193. Minimum Number of Moves to Make Palindrome](https://www.youtube.com/watch?v=IUB4G_hp8ug&ab_channel=HuifengGuan)
    - [00:00 - 28:40] Greedy + Simulation
    - [28:40 - ] Greedy + Reverse Pairs
