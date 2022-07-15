# [2335. Minimum Amount of Time to Fill Cups](https://leetcode.com/problems/minimum-amount-of-time-to-fill-cups/)

## Solution idea

### Priority Queue
**Key Oberservation**
题目要求，一秒最多取两种水，或者只取一种水。那么，为了总耗时最少，每一秒要尽可能一次性取两种水。

如何做到呢？就是每一秒优先取水量第一，第二高的两种水。

实现：`max heap`

Time complexity = $O(n\log 3)$ where $n$ is max amount among all types of water.

## Resources
[【每日一题】LeetCode 2335. Minimum Amount of Time to Fill Cups](https://www.youtube.com/watch?v=cuZBEzojuOw)