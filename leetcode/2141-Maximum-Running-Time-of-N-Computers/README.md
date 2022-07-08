# [2141. Maximum Running Time of N Computers](https://leetcode.com/problems/maximum-running-time-of-n-computers/)

## Solution idea

### 要点 1
**遇事不决想二分**

**求最大/最小若没有思路就想二分法**

对时间 `time` 进行二分：如果当前时间可行，当前时间至少是一个解，再看比它大的数。如果当前时间过大，当前时间一定不是解，看比它小的数。

### 要点 2
题目中，batteries 可随意插拔，随意分配。这就意味着，我们不用管具体的如何分配，只需检查所有 batteries 加起来是否满足要求的 `time*N` 时间.

同时，**注意** 每个battery最多贡献time时间，多余的时间没用。所以，要 `min(time, battery)`

Time complexity = $O(n\log m)$ where $n$ number of batteries, $\log m$ is for binary search.

## Resources
[【每日一题】LeetCode 2141. Maximum Running Time of N Computers](https://www.youtube.com/watch?v=rUq28MHmuoo)