# [1140. Stone Game II](https://leetcode.com/problems/stone-game-ii/description/)

## Solution idea
### DFS
#### 思路总结
1. 关键词: DFS, Prefix Sum, Memoization (2-D 区间型)
2. 为什么用于记忆化的 Memo 两个维度的长度都是 `len(piles)` ?
    * 第一维：比较好理解，代表起点位置，显然 piles 的任意位置都有可能是起点
    * 第二维 *(我的理解，不一定对)*：代表 M 的取值，任一 X 的取值都由 M 的值确定，也就是说，X 是一个关于 M 的函数 (`X = f(M)`)。X 代表长度，也就是，确定一个起点的offset。X 最长可以是整个 piles 的长度，也就是 `len(piles)`。M 确定 X 的上限。所以，第二维的长度也是 `len(piles)`。因为 `X <= 2M`，所以其实第二维的长度最长是 `len(piles) / 2`。但是，为了方便，直接用 `len(piles)`。
3. Definition:
    * `dfs(start, M)` := max # stones that cur player can get by first X piles (1 <= X <= 2*M) from `piles[start:]`

## Resource
[【每日一题】1140. Stone Game II, 8/27/2020](https://www.youtube.com/watch?v=liA938-cdfM&ab_channel=HuifengGuan)