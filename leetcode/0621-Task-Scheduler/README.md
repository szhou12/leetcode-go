# [621. Task Scheduler](https://leetcode.com/problems/task-scheduler/description/)

## Solution idea
### Solution 1: 全模拟 Using Max Heap
1. 理解题意: 当一个任务当下被安排以后，它不能在接下来的n个时间单位内再次安排。换句话说，在一个连续的`n+1`时间间隔内，安排`n+1`个不同的任务。当一个连续的`n+1`时间窗内，没法找全`n+1`个不同的任务，则用idle补足空缺。问: 怎么安排任务可以使总耗时最少？
2. 突破点: 怎么可以总耗时最少？很显然，用idle补空缺的机会尽量少。之所以会用idle补空缺，是因为任务的多样性不足。如何在尽可能安排不同任务的同时，尽可能多地保留剩余任务的多样性？**频次高的任务优先安排**，尽量延后使用idle的时机。如果反过来，频次低的优先安排，那么任务的多样性就会很快耗完，用idle补空缺的时机就会很快到来。

Time complexity = $O(n\log n)$

### Solution 2: Greedy Algorithm

## Resource
[【每日一题】621. Task Scheduler, 8/8/2021](https://www.youtube.com/watch?v=3DZE7cfgYyg&t=161s&ab_channel=HuifengGuan)
- Solution 1: 0:00 - 20:22
- Solution 2: 20:00 - end