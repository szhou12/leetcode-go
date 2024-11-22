# [2398. Maximum Number of Robots Within Budget](https://leetcode.com/problems/maximum-number-of-robots-within-budget/description/)

## Solution idea
### Binary Search 猜答案 + Sliding Window (Subarray Sum + Max Deque)
1. 用二分法猜一个robots的数量 k。
2. 用sliding window检验这个k是否可以小于budget。
    1. sliding window维护 subarray sum 较简单：吃进新的，吐掉旧的
    2. sliding window max (deque) 算法维护 subarray max

### Sliding Window Max (Deque)
1. Deque维护单调递减的序列。
2. 每一次一个新元素进来: 
    1. 淘汰屁股上的矮子: 如果deque末尾元素 < 新元素, 弹掉。重复直到deque内没有小于新元素。
    2. 新元素从屁股进来。
    3. 淘汰头部的老大个: 如果deque头部元素已经老了 (小于当前sliding window的start index)，也弹掉。 
4. 这样，deque**头元素**始终是 sliding window 的最大值。

Time complexity = $O(n\log n)$

## Resource
- [【每日一题】LeetCode 2398. Maximum Number of Robots Within Budget](https://www.youtube.com/watch?v=TVh5wAIJ2y8&ab_channel=HuifengGuan)
    - [5:55 - 9:45]Sliding Window Max
