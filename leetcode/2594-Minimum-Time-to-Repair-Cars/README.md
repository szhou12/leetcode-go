# [2594. Minimum Time to Repair Cars](https://leetcode.com/problems/minimum-time-to-repair-cars/description/)

## Solution idea
### Binary Search 猜答案
1. 二分搜索一个时间t。尝试给定这个时间t，是否可以修完所有的cars。
2. 注意！不用考虑具体每一个机器修多少辆车。给定的时间t是一个上限，那么每台机器在这个上限的时间内，最多可以修`sqrt(t/r)`辆车。判定条件只需要所有机器最多可修车数量的总和 >= cars就说明给定的t是一个可行解。

Time complexity = $O(r \log t)$