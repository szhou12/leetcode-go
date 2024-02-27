# [1235. Maximum Profit in Job Scheduling](https://leetcode.com/problems/maximum-profit-in-job-scheduling/description/)

## Solution idea
### Greedy (sorting by end date) + DP (2-D) + Binary Search (UpperBound) 
1. **Greedy (sorting by end date)**
* 此题涉及 "non-overlapping intervals" 的问题。大概率需要要尝试 "sort by end date"。
* :arrow_right: 用Greedy解决 "non-overlapping intervals" 的问题的一个核心前提是一定会选择第一个job。
* :arrow_right: BUT! 对于本题，总是选择第一个job并不是铁定最优解。
* :arrow_right: 所以，不能用Greedy无脑解决。
* :arrow_right: Greedy无脑解行不通，那么，就尝试DP的暴力解。同时，sort by end date这个思路依然可以沿用：一是保证了之后DP定义(在endTime维度上)的单增性；二是保证了可以使用Binary Search来进行快速的"回头看"
2. **DP (2-D)**
3. **Binary Search (UpperBound)**

## Resource
[【每日一题】1235. Maximum Profit in Job Scheduling, 1/9/2021](https://www.youtube.com/watch?v=0C7re8lam7M&ab_channel=HuifengGuan)