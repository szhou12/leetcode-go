# [3341. Find Minimum Time to Reach Last Room I](https://leetcode.com/problems/find-minimum-time-to-reach-last-room-i/description/)

## Solution idea
### DFS
- `dfs(x, y, curTime) := ` it takes `curTime` to travel from `(0, 0)` to `(x, y)`.
### Dijkstra
- **矩阵走格子类型题**。唯一需要注意的是，下一个格子进入PQ时需要判断: 如果当前时间不够，需要等到moveTime再进入，同时，不管需不需要等待，走到下一格子都需要花费1s. i.e., `max(curTime, moveTime[dx][dy]) + 1`

Time complexity = $O(mn\log{mn})$