# [2560. House Robber IV](https://leetcode.com/problems/house-robber-iv/description/)

## Solution idea

### Binary Search 猜答案 + 2-D DP

* **题目难点**: 需要理解题目中 capability, k 的含义以及它们的关系
    * capability: 指打劫n个房子中，打劫到最高金额的房子 $\Rightarrow $ 翻译: 打劫一个房子的金额**上限**
    * k: 打劫n个房子, $n \geq k$

* **难点突破**: capability, # of houses robbed, k 三者之间的关系
    * capability 越高 $\Rightarrow $ 打劫一个房子的金额上限越高 $\Rightarrow $ 越多的house可以rob $\Rightarrow $ 越容易满足 k 的要求
    * capability 越小 $\Rightarrow $ 打劫一个房子的金额上限越低 $\Rightarrow $ 越少的house可以rob $\Rightarrow $ 越难以满足 k 的要求
    * 根据以上的两条规律制定 Binary Search:
        * 猜 capability
        * 用一个helper function判断 猜测的capability 是否可以rob 至少k个房子
            1. 可以: 说明 猜测的capability 是一个可行解, 试试看再降一点可不可以
            2. 不可以: 说明 猜测的capability 绝对不是一个可行解, 试试看增大一点

* 实现 helper function: **2-D DP**
    * Definition:
        * `DP[i][0] =` # of houses we can rob in `nums[0...i]` if we DO NOT rob i-th house `nums[i]`
        * `DP[i][1] =` # of houses we can rob in `nums[0...i]` if we DO rob i-th house `nums[i]`
    * Base cases:
        1. `DP[0][0] = 0` if we DO NOT rob first house
        2. `DP[0][1] = 1` if we DO rob first house and `nums[0] <= cap`; otherwise, `DP[0][1] = MinInt`  if we DO rob first house but `nums[0] > cap` (用无限小代表无意义, 因为实际上超过了上限rob不到, 因为`DP`更新是用`max`, 所以用无限小)
    * Recurrrence:
        1. if `nums[i] > cap`:
            * `DP[i][0] = max(DP[i-1][0], DP[i-1][1])`
            * `DP[i][1] = MinInt`
        2. if `nums[i] <= cap`:
            * `DP[i][0] = max(DP[i-1][0], DP[i-1][1])`
            * `DP[i][1] = DP[i-1][0] + 1` (因为题目要求不能连续rob房子, 决定rob当前房子就不能rob前一个房子)

Time complexity = $O(n*\log D)$ where $D = $ difference between smallest possible capability (0) and largest possible capability (max(nums))


## Resource
[【每日一题】LeetCode 2560. House Robber IV](https://www.youtube.com/watch?v=_-nBUCSeU98&ab_channel=HuifengGuan)