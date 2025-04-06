# [3489. Zero Array Transformation IV](https://leetcode.com/problems/zero-array-transformation-iv/description/)

## Solution idea
### DP - Knapsack
1. 0/1背包问题Review
    ```
    背包i的容量        物品集合: {物品1容量, 物品2容量, ...}
    nums[i]          itemSet = {val1, val2, val3, ... }

    0/1: 指物品i放入or不放入背包

    dp[v] = True 物理意义: 背包容量为v时，通过挑选itemSet中的物品，组成subset，可以刚好占满整个背包

    0/1背包基础版：只有一个背包时，模版如下:
    // 外层循环: 放入or不放入当前物品 (占用容量d)
    for (d in itemSet) {
        // 内层循环: 当背包容量为v时, 放入d可以刚好占满背包容量(减到0) 当且仅当 上一轮的背包容量为v-d时是可以被占满的
        // 即, d可以入选subset占满v 当前仅当 之前可以找到一个subset' (d未入选时)占满v-d
        for v := 0; v <= maxCapacity; v++ {
            dp[v] = dp[v-d]
        }
    }

    Trick: 内层循环从大到小遍历
    上面模版中的内层循环有一个问题: 从小到大遍历时, dp[v-d]有可能使用的是本轮更新的值, 违反了定义。
    e.g. v=100, d=50, dp[0] = True. 
    那么, 当前轮中, dp[50] = dp[50-50] = dp[0] = True, 
    同时, dp[100] = dp[100-50] = dp[50] = True.
    dp[100]更新用到的dp[50]实际上是在本轮中才更新的. 物品i (重量=d) 放入了"两次"。但是, v-d容量应该对应的是尚未看到物品d时的subset, 也就是上一轮，从物品{val0, ..., vali-1}中组成的subset'。

    所以，正确的写法是:

    for (d in itemSet) {
        for v := maxCapacity; v >= 0 ; v-- {
            dp[v] = dp[v-d]
        }
    }
    ```

## Resource
[【每日一题】LeetCode 3489. Zero Array Transformation IV](https://www.youtube.com/watch?v=ilvbOwh1o_U&ab_channel=HuifengGuan)