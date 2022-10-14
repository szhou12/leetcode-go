# [416. Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/)

## Solution idea

### DP - 0/1背包问题

* 非常好的 0/1背包类问题

* 一件商品 = 一个元素, 商品重量 = 元素的值, 商品价值 = 元素的值

```
Definition:
DP[i][j] = whether or not we can find a subset from the first i elements whose sum can sum up to target weight j

Base Cases:
DP[0][0] = true
DP[i][0] = true // bc we can always find empty subset whose sum = 0
DP[0][j] = false for j > 1

Recurrence:
DP[i][j] = DP[i-1][j] || DP[i-1][j-nums[i]] if j >= nums[i] (也就是说背包容量可以容纳nums[i])
         = DP[i-1][j]                       o/w
```

Time complexity = $O(n^2)$

### DFS - 会超时！！！
* 求出所有可能的subsets，看哪个subset sum符合要求

Time complexity = $O(2^n)$

## Resource

0/1背包的详细解释：[代码随想录-416. 分割等和子集](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0416.%E5%88%86%E5%89%B2%E7%AD%89%E5%92%8C%E5%AD%90%E9%9B%86.md)
