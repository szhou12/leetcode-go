# [494. Target Sum](https://leetcode.com/problems/target-sum/)

## Solution idea

### DFS - All Subsets
* 使用 All Subsets第一种写法

Time complexity = $O(2^n)$

### DP - 0/1 Knapsack
* 需要数学推导，不容易想到

* 数学推导:

题目要求在数组元素前加上 + 或者 - 号，其实相当于把数组分成了 2 组，一组全部都加 + 号，一组都加 - 号。记 + 号的一组 P ，记 - 号的一组 N，那么可以推出以下的关系:

```
sum(P) - sum(N) = target
sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
                       2 * sum(P) = target + sum(nums)
```

等号两边都加上 `sum(N) + sum(P)`，于是可以得到结果 `2 * sum(P) = target + sum(nums)`，那么这道题就转换成了，能否在数组中找到这样一个集合，和等于 `(target + sum(nums)) / 2`。那么这题就转化为了第 416 题了.

* Definition

$DP[i] = $ # of ways to form positive sequence that sums up to == `i`

注意: 不合法的情况: 如果和不是偶数，即不能被 2 整除，或者和是负数，那说明找不到满足题目要求的解了，直接输出 0

* Base Case
$DP[0] = 1$ 装满容量为0的背包，有1种方法，就是装0件物品

* Recurrence
    * 不考虑`nums[i]`的情况下，填满容量为j的背包，有`dp[j]`种方法。
    * 那么考虑`nums[i]`的话（只要搞到`nums[i]`），凑成`dp[j]`就有`dp[j - nums[i]]` 种方法。
```
dp[j] += dp[j - nums[i]]
```

## Resource
[代码随想录-494. 目标和](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0494.%E7%9B%AE%E6%A0%87%E5%92%8C.md)

[halfrost/LeetCode-Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0494.Target-Sum)