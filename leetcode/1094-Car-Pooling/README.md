# [1094. Car Pooling](https://leetcode.com/problems/car-pooling/description/)

## Solution idea

### 1-D 差分数组

* 难点是这道题可以利用**差分数组**来解。解法套用模版基本上就能解决。

* 我们将每段trip拆解为上车和下车两部分。我们记录下每个站点的上车和下车人数，利用“差分”来维护每个站点的实际载客数，判断是否超载。

Time complexity = $O(n)$

## Resource

[小而美的算法技巧：差分数组](https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-48c1d/xiao-er-me-c31c8/)