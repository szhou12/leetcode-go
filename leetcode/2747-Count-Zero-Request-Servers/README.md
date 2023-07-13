# [2747. Count Zero Request Servers](https://leetcode.com/problems/count-zero-request-servers/description/)

## Solution idea
### Sliding Window (Fix)
#### 思路总结
1. 想到使用Sliding Window解题不难，容易想到先要 sort `logs`和`queries`。难的是怎么实现。
2. 实现：类Fix模版
    * 每个Sliding Window 就是 每个query区间的长度由`queries[i]`决定：`[queries[i] - x, queries[i]]`
    * Sliding Window的滑动由 `queries`的顺序决定
    * Sliding Window 的**物理意义**：当前query对应的时段内响应的servers数量
    * 右边界和左边界分别向Sliding Window框住的区间移动，右边界移动时“吃”，左边界移动时“吐”
    * 用 HashMap 记录 左右边界框住的区间内的每个server的数量
    * 每当左右边界进入Sliding Window框住的区间，计算一次 待机的servers = n - 左右边界框住的servers


Time complexity = $O(n\log n)$


## Resource
[【每日一题】LeetCode 2747. Count Zero Request Servers](https://www.youtube.com/watch?v=ShYCEh77RSc&ab_channel=HuifengGuan)