# [518. Coin Change II](https://leetcode.com/problems/coin-change-ii/)

## Solution idea

### DF - 完全背包问题

* Definition

`DP[j] = ` # of coin (物品) combinations that sum up to amount $j$ (背包容量)

* Base Case

`DP[0] = 1` 目标金额为0的组合数为1, 即, 空集

* Recurrence

`DP[j] += DP[j - coins[i]]`

* 难点: 这道题需要固定遍历顺序, 即, 只能外层遍历物品 (coins种类), 内层遍历背包容量 (目标金额)
```
for (int i = 0; i < len(coins); i++) { // 遍历物品, coins[i]一旦都尝试过，就不会再被考虑，防止了重复
    for (int j = coins[i]; j <= amount; j++) { // 遍历背包容量
        dp[j] += dp[j - coins[i]]
    }
}
```
**因为本题求组合数！！！**
假设：`coins[0] = 1`，`coins[1] = 5`

那么就是先把1加入计算，然后再把5加入计算，得到的方法数量只有{1, 5}这种情况。而不会出现{5, 1}的情况。

**所以这种遍历顺序中`dp[j]`里计算的是组合数！**

如果把两个for交换顺序:
```
for (int j = 0; j <= amount; j++) { // 遍历背包容量
    for (int i = 0; i < len(coins); i++) { // 遍历物品
        if (j - coins[i] >= 0) {
            dp[j] += dp[j - coins[i]]
        }
    }
}
```
背包容量的每一个值，都是经过 1 和 5 的计算，包含了{1, 5} 和 {5, 1}两种情况。

**此时`dp[j]`里算出来的就是排列数！**

Time complexity = $O(n*amount)$, Space complexity = $O(amount)$

### DFS - 会超时
```
# of branches = # of coins[i]
# levels = len(coins)
each level = coins[i]
```

## Resource

[代码随想录-518. 零钱兑换 II](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0518.%E9%9B%B6%E9%92%B1%E5%85%91%E6%8D%A2II.md)

[经典动态规划：完全背包问题](https://labuladong.github.io/algo/3/27/83/)