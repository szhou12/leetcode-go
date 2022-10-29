# [673. Number of Longest Increasing Subsequence](https://leetcode.com/problems/number-of-longest-increasing-subsequence/)

## Solution idea

### DP - LIS题的延伸

* 维护两个记事本:
    * `dp[i]`: the length of LIS ending at i
    * `count[i]`: the number of LIS ending at i

* 构建`dp[i]`与经典题LIS一样
* 如何构建 `count[i]`?
    1. if `dp[j]+1 > dp[i]`, 说明找到了一个更长的LIS, 它是目前的最长LIS, 所以`count[i]`重新继承`count[j]`
    2. if `dp[j]+1 == dp[i]`, 说明现在这个LIS属于目前最长的LIS之一, 所以`count[i]`新添加`count[j]`
* 用一个global variable `maxLeng` 在每一轮 `i` 记录看过的最长的LIS发生在以哪个i结尾的LIS
    * 用于最后遍历`dp`并把相等元素的index对应的 `count` 加入到结果中

Time complexity = $O(n^2)$

## Resource
[代码随想录-673.最长递增子序列的个数](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0673.%E6%9C%80%E9%95%BF%E9%80%92%E5%A2%9E%E5%AD%90%E5%BA%8F%E5%88%97%E7%9A%84%E4%B8%AA%E6%95%B0.md)