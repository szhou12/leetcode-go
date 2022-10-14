# [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)

## Solution idea

### Greedy Algorithm

* 难点：应该sort by start time? 还是 sort by end time?

    * **Sort by end time**: 就要从左向右遍历，因为右边界越小越好，只要右边界越小，留给下一个区间的空间就越大，所以从左向右遍历，优先选右边界小的。
    * **Sort by start time**: 就要从右向左遍历，因为左边界数值越大越好（越靠右），这样就给前一个区间的空间就越大，所以可以从右向左遍历。

Time complexity = $O(n\log n)$

正确性证明：参考MPCS - Algorithm课讲义

## Resrouce

[代码随想录-435. 无重叠区间](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0435.%E6%97%A0%E9%87%8D%E5%8F%A0%E5%8C%BA%E9%97%B4.md)