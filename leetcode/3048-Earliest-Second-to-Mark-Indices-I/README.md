# [3048. Earliest Second to Mark Indices I](https://leetcode.com/problems/earliest-second-to-mark-indices-i/description/)

## Solution idea
### Binary Search 猜答案 + Greedy
1. 理解题意: `changeIndices[]`代表一个时间线，每一个index表示一个时刻t。每一个时刻t可以做两件事情: 1. decrement: 挑选任意一个`nums[]`中一个非0元素减1 (注意: `nums[]`中元素数值与`changeIndices[]`的元素数值没有关系，单纯表示一个量，要减多少次为0)；2. mark: 如果时刻t对应的`nums[]`位置上的数值(`nums[changeIndices[t]]`)已经减为0，可以进行mark操作。题目问：找到最早的时刻使得所有`nums[]`的位置可以被mark。
2. 首要直觉：**单调性** - 时间越长，越容易清0，越容易mark尽可能多的位置。4秒钟能完成任务，那么8秒钟也一定能完成任务。
3. **单调性** + **最优问题** => **Binary Search**
4. Why Use Binary Search? BS把最优问题转化为可行性问题，即，判断一个解是否可行。判断可行解永远比解决最优问题容易得多。

## Resource
[【每日一题】LeetCode 3048. Earliest Second to Mark Indices I](https://www.youtube.com/watch?v=xXqQeLepYm4&ab_channel=HuifengGuan)