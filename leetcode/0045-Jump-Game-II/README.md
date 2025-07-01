# [45. Jump Game II](https://leetcode.com/problems/jump-game-ii/)

## Solution idea

### My solution - DP
Define $DP[i]$ as minimun jumps from index $0$ to index $i$.

Base case:

$DP[0] = 0$ as need $0$ jumps at index $0$.

Recurrence:

$DP[i] = \min(DP[j])$ for $0 \leq j < i$ and $nums[j] \geq i-j$ (i.e. look at all $j$'s who can reach $i$)

Time Complexity = $O(n^2)$

### Greedy
1. 想象成BFS, 每一层=可以跳到的一段连续区间

Time complexity = $O(n)$

## Resource 
[【每日一题】45. Jump Game II, 9/29/2020](https://www.youtube.com/watch?v=Ua_Vqtdd61E&ab_channel=HuifengGuan)