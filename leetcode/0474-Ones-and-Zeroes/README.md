# [474. Ones and Zeroes](https://leetcode.com/problems/ones-and-zeroes/)

## Solution idea

### DP - 2维 0/1 Knapsack

* 典型的 0/1 背包的题型. 在两个维度上分别做 1维 0/1 Knapsack

* 在 n 个String中选出一些String，尽量完全填满 m 维和 n 维 这两个背包

#### Definition:

`dp[i][j]`: 最多有i个0和j个1的背包装下的String总数

#### Base Case:

01背包的dp数组初始化为0就可以。

因为物品价值不会是负数，初始为0，保证递推的时候`dp[i][j]`不会被初始值覆盖 (因为是用 max())。

#### Recurrence
```
dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones] + 1)
```

* 注意！！！遍历背包容量要**从后向前遍历**！这是为了防止double counting.

举一个例子：物品0的重量`weight[0] = 1`，价值`value[0] = 15`

如果正序遍历

`dp[1] = dp[1 - weight[0]] + value[0] = 15`

```
dp[2] = dp[2 - weight[0]] + value[0]
      = dp[2-1] + value[0] 
      = dp[1] + value[0] 
      = (dp[1 - weight[0]] + value[0]) + value[0]
```
`dp[1]` 就计算了添加物品0的情况.

此时`dp[2]`在考虑`dp[1]`时，意味着物品0，又被放入了一次，所以不能正序遍历。

倒序就是先算`dp[2]`, 也就是说, 等会儿算`dp[1]` 时, `dp[1]`并不会考虑`dp[2]`的情况.

所以从后往前循环，每次取得状态不会和之前取得状态重合，这样每种物品就只取一次了.

Time complexity = $O(n * M * N)$ where $n=$ # of strings, $M, N$ 是两个维度上的背包最大容量

## Resource

* [代码随想录-474.一和零](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0474.%E4%B8%80%E5%92%8C%E9%9B%B6.md)

* 详细解释倒序的原因: [动态规划：关于01背包问题，你该了解这些！（滚动数组）](https://programmercarl.com/%E8%83%8C%E5%8C%85%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%8001%E8%83%8C%E5%8C%85-2.html#%E4%B8%80%E7%BB%B4dp%E6%95%B0%E7%BB%84-%E6%BB%9A%E5%8A%A8%E6%95%B0%E7%BB%84)

* [halfrost/LeetCode-Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0474.Ones-and-Zeroes)