# [3459. Length of Longest V-Shaped Diagonal Segment](https://leetcode.com/problems/length-of-longest-v-shaped-diagonal-segment/description/)

## Solution idea
### DFS + Memoization
1. 理解题意：从一个cell=1作起点，朝一个方向(左上，右上，右下，左下)走，走的路径必须满足`2 -> 0 -> 2 -> 0 -> ...`。每走一步加一分。中间可以有一次顺时针90度拐弯的机会(e.g.从右下转到左下)。问最长路径分值。
2. 思路：没有思路的时候先想想用DFS暴力解。看到constraints: `1 <= n, m <= 500`，那么每个cell遍历一遍是`O(n*m) = 10^5`，可以接受。然后想到，每个cell的“状态“ - 一个
cell可以从四个方向中的其中一个方向到达，所以有四种状态。同时，到达一个cell时，可以有用了”拐弯“和没用过两种状态。所以，一共是`O(n*m*4*2) = 10^6`，刚刚好。那么，解法就是使用DFS。
3. 实现：
    - 定义: `dfs(i, j, d, turn) =` 站在`cell(i, j)`，朝着`d`方向走，还剩`turn`次“拐弯机会”，接下来所能得到的最长路径(最多得分)。
    - 两条分支:
        1. 继续沿着`d`方向走，只要下一步`cell(di, dj)`没有出界 并且 遵守了`1 -> 2 -> 0 -> 2 -> 0 ...`的pattern，那么就可以往下走：`1 + dfs(di, dj, d, turn)`
        2. 如果还没有使用“拐弯”机会，可以尝试拐个弯。同样的，只要下一步`cell(di, dj)`没有出界 并且 遵守了`1 -> 2 -> 0 -> 2 -> 0 ...`的pattern，那么就可以往下走：`1 + dfs(di, dj, newD, 0)`
    - 最后，往回返的时候取两条分支中的最大值。
4. 优化: 每个cell的状态可以用 `memo = [500][500][4][2]int` 来记录。如果当前`cell(i, j)`查memo发现已经更新过，就不用再算一遍。如果没有走过，两条分支走完后，更新当前`cell(i, j)`的memo。

Time complexity = $O(n*m*4*2)$

## Resource
[【每日一题】LeetCode 3459. Length of Longest V-Shaped Diagonal Segment](https://www.youtube.com/watch?v=5CvMx5TKsX0&t=199s&ab_channel=HuifengGuan)