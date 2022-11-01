# [1254. Number of Closed Islands](https://leetcode.com/problems/number-of-closed-islands/)

## Solution idea

### DFS - 岛屿沉没类

* [200. Number of Islands](https://leetcode.com/problems/number-of-islands/submissions/)的进阶题

* 与[1905. Count Sub Islands](https://leetcode.com/problems/count-sub-islands/)的突破点一样, 都是要先找到不符合题意的岛屿提前沉没掉, 再数剩下的符合要求的岛屿个数

* **突破点**: 哪些是不合题意的岛屿???
    * 依题意, closed islands是四面环水的岛屿, 也就是说, **靠边界的陆地不算作closed islands**
    * 提前沉没掉贴着上、下、左、右边界的岛屿, 剩下的就是符合题意的岛屿

Time complexity = $O(m*n)$

## Resource
[一文秒杀所有岛屿题目](https://labuladong.github.io/algo/4/31/108/)