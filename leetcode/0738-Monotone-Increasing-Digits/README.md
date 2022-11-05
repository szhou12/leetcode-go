# [738. Monotone Increasing Digits](https://leetcode.com/problems/monotone-increasing-digits/description/)

## Solution idea
### Greedy
* 如果input number不符题意, 从后往前, 前一个digit - 1, 当前digit变成9

* 局部最优：遇到`strNum[i - 1] > strNum[i]`的情况，让`strNum[i - 1]--`，然后`strNum[i]`给为9，可以保证这两位变成最大单调递增整数。
* 全局最优：得到小于等于N的最大单调递增的整数。

## Resource
[代码随想录-738.单调递增的数字](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0738.%E5%8D%95%E8%B0%83%E9%80%92%E5%A2%9E%E7%9A%84%E6%95%B0%E5%AD%97.md)