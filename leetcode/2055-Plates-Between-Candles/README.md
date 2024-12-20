# [2055. Plates Between Candles](https://leetcode.com/problems/plates-between-candles/description/)

## Solution idea
### Greedy + State Machine + Prefix Sum
1. 突破点：这道题到底在问什么？题目要求指定一段区间内被两个candle包裹住的最多plates。实际上，这道题就是在问：这段区间里面，最外层的两个candle包裹住的盘子。有了这个理解，就比较好想解决方案了。
2. 如何高效求出区间内最外层的两个candle？i.e. 左边界 和 有边界
    - ”状态机“思想：对于任意i，找出离i最近的右边的candle (找不到就标记为n)，为`next[i]`(“左边界”)；找出离i最近的左边的candle (找不到就标记为-1)，为`last[i]`(“右边界”)。
    - 对于任意一个query，左端点找出`next[a]`，右端点找出`last[b]`。
    - 坑：我们要检查找出的这段由左右边界夹住的区间是有效的。i.e. `next[a] <= last[b]`
3. 如何高效求出区间内plates总数？i.e. range sum
    - range sum 很明显可以用prefix sum来解决。
    - 任意i，计算从0到i的所有plates数量。

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 2055. Plates Between Candles](https://www.youtube.com/watch?v=bYBP-Z0gmcg&ab_channel=HuifengGuan)