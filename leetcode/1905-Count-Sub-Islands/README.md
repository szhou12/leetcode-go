# [1905. Count Sub Islands](https://leetcode.com/problems/count-sub-islands/)

## Solution idea

### DFS - 岛屿沉没类

* [200. Number of Islands](https://leetcode.com/problems/number-of-islands/submissions/)的进阶题

* 与[1254. Number of Closed Islands](https://leetcode.com/problems/number-of-closed-islands/)的突破点一样, 都是要先找到不符合题意的岛屿提前沉没掉, 再数剩下的符合要求的岛屿个数

* **突破点**: 哪些是不合题意的岛屿???
    * 依题意, 如果grid2中一个岛屿 B 是 grid1 中一个岛屿 A 的一个子岛屿 (sub-island), 那么, 岛屿 B 所有是陆地的地方, 岛屿 A 对应的地方一定也是陆地
    * **反过来说，如果岛屿 B 中存在一片陆地，在岛屿 A 的对应位置是海水，那么岛屿 B 就不是岛屿 A 的子岛**
    * 提前沉没掉贴这些岛屿, 剩下的就是符合题意的岛屿

Time complexity = $O(m*n)$

## Resource
[一文秒杀所有岛屿题目](https://labuladong.github.io/algo/4/31/108/)