# [56. Merge Intervals](https://leetcode.com/problems/merge-intervals/description/)

## Solution idea

### Greedy Algorithm

**Key 1**: 按照**start time**从小到大排序

**Key 2**: i-th end time 比较 (i+1)-th start time找重叠，merge时选择较大一个的 end time

Time complexity = $O(n\log n)$

**注意！！！** 这道题和find largest compatible subset 不一样，那道题sort end time, 这道题sort start time