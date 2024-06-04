# [3169. Count Days Without Meetings](https://leetcode.com/problems/count-days-without-meetings/description/)

## Solution idea
### Method 1: Greedy Algorithm - Scheduling
1. 典型的 Meeting Scheduling 问题，可以用贪心算法解决。
2. 主要问题是：如果sort所有的meetings？
    1. 先按照每个meeting的**起始点从小到大**排序。
    2. 如果起始点相同，按照**结束点从大到小**排序。
        - 为什么结束点从大到小？举个例子: `[1,48], [1,6]`. 这样我们先跳到结束点靠后的时间点，之后每次与前一个比较时，就不会错误计算被覆盖的meeting。也就是说，`[1,6]`被`[1,48]`覆盖，我们应该直接跳过`[1,6]`，而不应该再拿它与前一个比较是否compatible。
3. 注意！在给出的实现中，按照结束点排序这一步其实可有可无。去掉并不影响正确性。这是因为，代码中`if curDate <= end`会确保在相同起始点时，最后落到最晚结束点并且此过程中的更新正确。
Time complexity = $O(n\log n)$

### Method 2: Sweep Line
1. 典型的 Sweep Line 问题，也可以用常规扫描线的方法解决。
2. 扫描线的使用:
    1. 每一个meeting，起始点标记+1，结束点之后一天标记-1。用一个Map储存。撸过所有meetings之后，每一个时间点就记录了这一天meetings的“增量” (可正可负)。
    2. 按照时间线，扫描一遍，计算在每一个时间点处的积分(meetings的累积量)。任意时刻的积分一定是非负的。
    3. 某一时刻的积分如果从0上升为正数，说明“好日子”到头了，计算截止到这一天（不包括这一天）的天数，记为没有meetings的时期。否则，这一时刻如果积分从正数下降为0，说明“好日子”又开始了，储存这一天为没有meetings的最早时刻。
    4. 注意！在给定的时间线`days`上，meeting不一定会出现在最后一天。也就是说，有可能从某一时刻直到最后一天是没有meetings的“好日子”。为了让扫描线捕捉到这种情况，需要在`days+1`这一天设置一个dummy增量 (赋一个正数)。
3. 注意！储存增量的Map需要有序。C++中的`std::map`是天然有序的。如果使用Go来解题，Go中map是无序的，所以需要额外使用一个slice来储存所有map中出现的key，然后进行从小到大排序。然后，用扫描线扫描这个slice所有的日子并计算积分。

Time complexity = $O(n\log n)$

## Resource
[【每日一题】LeetCode 3169. Count Days Without Meetings](https://www.youtube.com/watch?v=AD31HbsbQF8&ab_channel=HuifengGuan)