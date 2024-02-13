# Greedy Algorithm

## Scheduling 类型题

* 区间调度问题思路三步走:
    1. 从区间集合 `intervals` 中选择一个区间 `x`，这个 `x` 是在当前所有区间中**结束最早的**（end 最小）
    2. 把所有与 `x` 区间相交的区间从区间集合 `intervals` 中删除。
    3. 重复 Step 1 和 Step 2，直到 `intervals` 为空为止。之前选出的那些 `x` 就是最大不相交子集。

* 排除 Overlapping Invertals: [435. Non-overlapping Intervals](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0435-Non-overlapping-Intervals)
    * **重点**: 如果要loop from left to right, 则需要sort by **right end** in increasing order

* 计算 Overlapping Invertals: [452. Minimum Number of Arrows to Burst Balloons](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0452-Minimum-Number-of-Arrows-to-Burst-Balloons)
    * **重点**: 如果要loop from left to right, 则需要sort by **right end** in increasing order

* Overlapping Intervals归为同一个 partition, 计算partitions: [763. Partition Labels](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0763-Partition-Labels)
    * **重点**: loop from left to right, **BUT** sort by **left end** in increasing order.
    * 这是因为sort by left end 可以迅速确认每个partition最小左边界 (每个partition第一个遇到的interval), 这样, 只需要在loop中不断找partition的最大右边界就行。

## 找规律

* :red_circle: 给Alice和Bob寻找合适位置II: [3027. Find the Number of Ways to Place People II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3027-Find-the-Number-of-Ways-to-Place-People-II)
    * 2-D坐标排序

* 单调递增位数: [738. Monotone Increasing Digits](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0738-Monotone-Increasing-Digits)
    * 比当前number小 并且每个digit单调递增的 最大number 是 (X-1)9...9
    * e.g. 93 -> 89 = (9-1)(3->9)

* 最大化k次取反后的array sum: [1005. Maximize Sum Of Array After K Negations](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1005-Maximize-Sum-Of-Array-After-K-Negations)
    * 将数组按照绝对值大小**从大到小**排序，注意要**按照绝对值的大小**

* 找第k小array sum: [1439. Find the Kth Smallest Sum of a Matrix With Sorted Rows](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1439-Find-the-Kth-Smallest-Sum-of-a-Matrix-With-Sorted-Rows)
    * 逐行merge, 每merge完一行取前k小的array sum

## 分清楚情况讨论

* 找零的情况: [860. Lemonade Change](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0860-Lemonade-Change)
