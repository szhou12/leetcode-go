# [2335. Minimum Amount of Time to Fill Cups](https://leetcode.com/problems/minimum-amount-of-time-to-fill-cups/)

## Solution idea

### Priority Queue
**Key Oberservation**
题目要求，一秒最多取两种水，或者只取一种水。那么，为了总耗时最少，每一秒要尽可能一次性取两种水。

如何做到呢？就是每一秒优先取水量第一，第二高的两种水。

实现：`max heap`

Time complexity = $O(n\log 3)$ where $n$ is max amount among all types of water.

### Majority Vote/Pigeonhole Principle - Optimal Solution
There are two cases:

Case 1: if there is one type of water `A` who is 绝对多数，meaning its amount $\geq$ total amount of water:

那么，每一轮都是`A`和一种弱势群体的一种水配对，最后剩下`A`自己，每一轮取`A`自己

Case 2: if there is no type of water that is 绝对多数, meaning we can pair two types in each round:

那么，每一轮都可以由两种水配对。如果总水量是奇数，最后一轮会剩下一个单独的。如果偶数，所用时间就是总水量/2.

Time complexity = $O(1)$

## Resources
[【每日一题】LeetCode 2335. Minimum Amount of Time to Fill Cups](https://www.youtube.com/watch?v=cuZBEzojuOw)