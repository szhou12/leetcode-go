# [134. Gas Station](https://leetcode.com/problems/gas-station/)

## Solution idea:

### Greedy Solution

Key observations:

1. If we can go from `i` to `i+1`, then `tank = tank_i + gas[i] - cost[i] > 0` where `tank_i >= 0`.

2. At starting point, `tank = 0`.

3. Pre-processment: if `sum(gas[...] - cost[...]) < 0`, definitely cannot complete the circle. (如果总油量小于总的消耗，那肯定是没办法环游所有站点的).

Proof of Correctness:

Claim: If at $i$ as starting point, we can't just(恰好) reach $j$, then for all $k$ where $i < k < j$, we can't reach $j$ starting at $k$.

Proof:

If at $i$ as starting point, we can't just(恰好) reach $j$, this means $tank_j < 0$. 

If at $i$ as starting point, we can't just(恰好) reach $j$, this also means we CAN reach any $k<j$ from $i$. i.e. $tank_k > 0$.

Since even under the situation where $tank_k > 0$, we can't reach $j$, it's even more impossible to reach $j$ from $k$ if treating $k$ as the starting point because if $k$ is starting point, then $tank_k=0$.

Thus, the claim is correct.

The claim implies that $i$ and any $k<j$ can't be legal starting point candidates, which means if a legal starting point exits, it must exist at $j$ or beyond.

### Note:
Why in the algorithm when we reset `start = i + 1`, we don't need to worry about `0...i` part?

i.e. Why is it sufficient to only check `tank` between `i+1...n-1`?

Because we have shown that `i` is the point that can't just(恰好) be reached from `start=0`.

This means that for any `j` where `0<j<i`, we can reach `j` from `start=0` ($tank_j > 0$).

Therefore, if we can reach from `i+1` to `n-1`, we definitely can reach from `i+1` to `j`

(`[0 ... j ... i, i+1, ... n-1]`)

### References
[当老司机学会了贪心算法](https://labuladong.github.io/algo/3/27/101/)

[贾考博 LeetCode 134. Gas Station 加不起油怎么算?](https://www.youtube.com/watch?v=bkXokc5hh14)
