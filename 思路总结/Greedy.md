# Greedy Algorithm

## 目录
* [Scheduling 类型题](#scheduling-类型题)
* [找规律](#找规律)
* [分清楚情况讨论](#分清楚情况讨论)
* [K sorted lists](#k-sorted-lists)
* [Arrangement With Stride](#arrangement-with-stride)

## Scheduling 类型题

* 区间调度问题思路三步走:
    1. 从区间集合 `intervals` 中选择一个区间 `x`，这个 `x` 是在当前所有区间中**结束最早的**（end 最小）
    2. 把所有与 `x` 区间相交的区间从区间集合 `intervals` 中删除。
    3. 重复 Step 1 和 Step 2，直到 `intervals` 为空为止。之前选出的那些 `x` 就是最大不相交子集。
* 注意事项:
    1. Scheduling问题不能死板地记：sort by start/end date.
    2. sort谁由下一步的`finding compatible jobs function`的定义决定。即，组成 compatible set的元素是按什么条件筛选的。e.g. 1353中，筛选规则就是挑start date不晚于loop的当前天的。所以, sort by start date方便进行贪心(每一步不回头看)。

* :yellow_circle: 统计不开会的“好日子”: [3169. Count Days Without Meetings]()

* :secret: :red_circle: schedule jobs同时最大化利益：[1235. Maximum Profit in Job Scheduling](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1235-Maximum-Profit-in-Job-Scheduling)
    * 非常好非常综合的题目！！！
    * Greedy (Sort by end date) + DP + Binary Search (UpperBound)
    * 不使用Greedy解，但是沿用sort的思路

* :red_circle: deadline前可以完成最多任务: [1353. Maximum Number of Events That Can Be Attended](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1353-Maximum-Number-of-Events-That-Can-Be-Attended)
    * Sort + Priority Queue: Sort by start date; minHeap by end date

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

* :yellow_circle: 聚类黑白球-只允许相邻位置的交换: [2938. Separate Black and White Balls](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2938-Separate-Black-and-White-Balls)

## 分清楚情况讨论

* :red_circle: 最小相等的array sum: [2918. Minimum Equal Sum of Two Arrays After Replacing Zeros](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2918-Minimum-Equal-Sum-of-Two-Arrays-After-Replacing-Zeros)

* 找零的情况: [860. Lemonade Change](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0860-Lemonade-Change)

## K sorted lists
* :red_circle: 最小范围包含所有list中至少一个元素: [632. Smallest Range Covering Elements from K Lists](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0632-Smallest-Range-Covering-Elements-from-K-Lists)

## Arrangement With Stride
* :red_circle: 给定时间窗内安排不同任务: [621. Task Scheduler](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0621-Task-Scheduler)
    * Solution 1: `maxHeap{(freq, task), ...}`