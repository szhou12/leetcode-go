# [2343. Query Kth Smallest Trimmed Number](https://leetcode.com/problems/query-kth-smallest-trimmed-number/)

## Solution idea

### Priority Queue
Use maxheap to get k-th smallest for each query

Time complexity = $O(mn\log k)$ where $m$ number of queries, $n$ number of numbers.

### Optimal Solution - DP + 基准排序
基准排序： 类似于 `Bucket Sort`。先排序一个digit， 再排序两个digits. 十位数上的数字只有 `0~9`.

Time complexity = $O(mn + \| queries \|)$ where $m$ number of numbers, $n$ length of one number (all numbers equal length).

**Resources**

[【每日一题】LeetCode 2343. Query Kth Smallest Trimmed Number](https://www.youtube.com/watch?v=SlWh1ih7p74)

