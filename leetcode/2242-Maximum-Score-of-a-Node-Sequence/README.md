# [2242. Maximum Score of a Node Sequence](https://leetcode.com/problems/maximum-score-of-a-node-sequence/)

## Solution idea

突破点：

valid node sequence with a length of 4 

$\Rightarrow$ Brute-force

$\Rightarrow$ for any `a-b` (middle edge) in `i-a-b-j`, select neighbors `i, j` such that it gives max scores

$\Rightarrow$ for any node's neighbors, we only keep top 3 neighbors with highest scores (Greedy comes to play)

Time complexity = $O(E + V^2\log V)$ ?

## Resource
[【每日一题】LeetCode 2242. Maximum Score of a Node Sequence](https://www.youtube.com/watch?v=7eVhDmozOTE&ab_channel=HuifengGuan)