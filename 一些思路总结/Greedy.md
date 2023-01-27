# Greedy Algorithm

## Scheduling 类型题

* 区间调度问题思路三步走:
    1. 从区间集合 `intervals` 中选择一个区间 `x`，这个 `x` 是在当前所有区间中**结束最早的**（end 最小）
    2. 把所有与 `x` 区间相交的区间从区间集合 `intervals` 中删除。
    3. 重复 Step 1 和 Step 2，直到 `intervals` 为空为止。之前选出的那些 `x` 就是最大不相交子集。

* 排除 Overlapping Invertals: [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)
    * **重点**: 如果要loop from left to right, 则需要sort by **right end** in increasing order

* 计算 Overlapping Invertals: [452. Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)
    * **重点**: 如果要loop from left to right, 则需要sort by **right end** in increasing order

* Overlapping Intervals归为同一个 partition, 计算partitions: [763. Partition Labels](https://leetcode.com/problems/partition-labels/)
    * **重点**: loop from left to right, **BUT** sort by **left end** in increasing order.
    * 这是因为sort by left end 可以迅速确认每个partition最小左边界 (每个partition第一个遇到的interval), 这样, 只需要在loop中不断找partition的最大右边界就行。

## 找规律

* 单调递增位数: [738. Monotone Increasing Digits](https://leetcode.com/problems/monotone-increasing-digits/description/)
    * 比当前number小 并且每个digit单调递增的 最大number 是 (X-1)9...9
    * e.g. 93 -> 89 = (9-1)(3->9)

* 最大化k次取反后的array sum: [1005. Maximize Sum Of Array After K Negations](https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/description/)
    * 将数组按照绝对值大小**从大到小**排序，注意要**按照绝对值的大小**

* 找第k小array sum: [1439. Find the Kth Smallest Sum of a Matrix With Sorted Rows](https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/description/)
    * 逐行merge, 每merge完一行取前k小的array sum

## 分清楚情况讨论

* 找零的情况: [860. Lemonade Change](https://leetcode.com/problems/lemonade-change/description/)