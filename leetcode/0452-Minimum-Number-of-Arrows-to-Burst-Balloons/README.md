# [452. Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)

## Solution idea

### Greedy Algorithm 

* 思路与 [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/) 一致，都是通过Greedy Algorithm 找 overlapping invertals

* Sort by end so that we can loop from left to right

Time complexity = $O(n\log n)$