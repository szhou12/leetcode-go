# [1235. Maximum Profit in Job Scheduling](https://leetcode.com/problems/maximum-profit-in-job-scheduling/description/)

## Solution idea
### Greedy (sorting by end date) + DP (2-D) + Binary Search (UpperBound) 
1. **Greedy (sorting by end date)**
* 此题涉及 "non-overlapping intervals" 的问题。大概率需要要尝试 "sort by end date"。马上应该联想到的例题是: [435. Non-overlapping Intervals](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0435-Non-overlapping-Intervals)
    * :arrow_right: 用Greedy解决 "non-overlapping intervals" 问题的一个核心前提是一定会选择第一个job。
    * :arrow_right: BUT! 对于本题，因为要求第二个维度，也就是利益的最大化，总是选择第一个job(时间线上最早的)并不是铁定最优解。
    * :arrow_right: 所以，不能用Greedy无脑解决。
    * :arrow_right: Greedy无脑解行不通，那么，就尝试DP的暴力解。同时，sort by end date这个思路依然可以沿用：一是保证了之后DP定义(在endTime维度上)的单增性；二是保证了可以使用Binary Search来进行快速的"回头看"
2. **DP (2-D)**
* `dp[i]`可以有两种定义。用第二种定义可以优化时间复杂度。
    1. `dp[i] :=` max profit if we MUST take job `i`
        * `dp[i] = max(dp[j] + profit[i])` where `j` are all compatible jobs with `i` (i.e. `endTime[j] <= startTime[i]`).
        * `dp[i]` not necessarily monotonic increasing. Thus, need to all possible `j` for each `i`.
        * Time complexity: `O(n^2)`
    2. `dp[i] :=` max profit at job `i`'s endTime (may or may not include job `i`'s profit) 
        * `dp[i] = max(dp[j] + profit[i], dp[i-1])` where `j` is the latest (WHY?) job that is compatible with `i`.
        * WHY can `j` be just the latest? Because `dp[i]` in this definition is monotonic increasing. The max profit is only gonna increase as time goes by. Thus, the latest compatible `j` is giving the higher profit than any earlier compatible jobs.
        * Time complexity: `O(nlogn)` by using Binary Search (UpperBound).
3. **Binary Search (UpperBound)**
* Jobs are sorted by end time + Using 2nd definition of `dp[i]` => We can use Binary Search to efficiently find the latest compatible job `j` for `i`
* `upperBound(nums []int, target int)` gives the first index `i` s.t. `nums[i] > target`. The latest compatible job `j` is defined as `endTime[j] <= startTime[i]`. Thus, `i-1` is what we want.
* 魔改`upperBound(nums []int, target int)` as `upperBound(dp [][]int, target int)`
    * `dp = [[endTime at job 0, maxProfit], [endTime at job 1, maxProfit], ...]`
    * `target = startTime[i]`

Time complexity = $O(n\log n)$

## Resource
[【每日一题】1235. Maximum Profit in Job Scheduling, 1/9/2021](https://www.youtube.com/watch?v=0C7re8lam7M&ab_channel=HuifengGuan)