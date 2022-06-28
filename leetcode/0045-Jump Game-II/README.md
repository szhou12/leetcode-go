# [45. Jump Game II](https://leetcode.com/problems/jump-game-ii/)

## Solution idea

### My solution - DP
Define $DP[i]$ as minimun jumps from index $0$ to index $i$.

Base case:

$DP[0] = 0$ as need $0$ jumps at index $0$.

Recurrence:

$DP[i] = \min(DP[j])$ for $0 \leq j < i$ and $nums[j] \geq i-j$ (i.e. look at all $j$'s who can reach $is$)

Time Complexity = $O(n^2)$