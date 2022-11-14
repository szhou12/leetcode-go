# [452. Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)

## Solution idea

### Greedy Algorithm 

* 思路与 [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/) 一致，都是通过Greedy Algorithm 找 overlapping invertals

* Sort by end so that we can loop from left to right

* 区间调度问题思路三步走:
    1. 从区间集合 `intervals` 中选择一个区间 `x`，这个 `x` 是在当前所有区间中**结束最早的**（end 最小）
    2. 把所有与 `x` 区间相交的区间从区间集合 `intervals` 中删除。
    3. 重复 Step 1 和 Step 2，直到 `intervals` 为空为止。之前选出的那些 `x` 就是最大不相交子集。

Time complexity = $O(n\log n)$