# Greedy Algorithm

## Scheduling 类型题

* 排除 Overlapping Invertals: [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)
    * **重点**: 如果要loop from left to right, 则需要sort by **right end** in increasing order

* 计算 Overlapping Invertals: [452. Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)
    * **重点**: 如果要loop from left to right, 则需要sort by **right end** in increasing order

* Overlapping Intervals归为同一个 partition, 计算partitions: [763. Partition Labels](https://leetcode.com/problems/partition-labels/)
    * **重点**: loop from left to right, BUT sort by **left end** in increasing order.
    * 这是因为sort by left end 可以迅速确认每个partition最小左边界 (每个partition第一个遇到的interval), 这样, 只需要在loop中不断找partition的最大右边界就行。

## 未归类

* [738. Monotone Increasing Digits](https://leetcode.com/problems/monotone-increasing-digits/description/)