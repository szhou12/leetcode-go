# [188. Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/)

## Solution idea

### DP

* Definition:
    * `DP[i][k][0]` = Max profit on $i$-th day, with at most $k$ transactions, and don't own a stock (i.e. sell a stock)
    * `DP[i][k][0]` = Max profit on $i$-th day, with at most $k$ transactions, and own a stock (i.e. buy a stock)

* Base Cases:
    * `k==0`: 
        1. `DP[i][0][0]` = 0: if allow 0 transactions, profit we can have is 0.
        2. `DP[i][0][1]` = `-inf`: if allow 0 transactions, illegal to buy/own any stock, profit we can have is set to `-infinity`
    * `i==0`:
        1. `DP[0][1...k][0]` = 0: On first day we don't own a stock, profit we can have is 0.
        2. `DP[0][1...k][0]` = `-prices[0]`: On first day we buy a stock, profict we have is `-prices[0]`.

* Recurrence: $0 < i < n, 1 \leq k \leq k_{max}$
    * `DP[i][k][0] = max(DP[i-1][k][0], DP[i-1][k][1] + prices[i])` (either we rest or sell a stock on previous day)
    * `DP[i][k][1] = max(DP[i-1][k][1], DP[i-1][k-1][0] - prices[i])` (either we rest or buy a stock on current day)

Time complexity = $O(n*k)$

## Resource

详细解释参考：

[一个方法团灭 LEETCODE 股票买卖问题](https://labuladong.github.io/algo/3/28/96/)

一些重要的takeout:
* DP算法本质上就是穷举「状态」，然后在「选择」中选择最优解
* 这个问题，每天都有三种「选择」：买入、卖出、无操作，我们用 `buy`, `sell`, `rest` 表示这三种选择
```
for 状态1 in 状态1的所有取值：
    for 状态2 in 状态2的所有取值：
        for ...
            dp[状态1][状态2][...] = 择优(选择1，选择2...)
```
* 这个问题的「状态」有三个，第一个是天数，第二个是允许交易的最大次数，第三个是当前的持有状态（即之前说的 `rest` 的状态，我们不妨用 1 表示持有，0 表示没有持有）
```
dp[i][k][0 or 1]
0 <= i <= n - 1, 1 <= k <= K
n 为天数，大 K 为交易数的上限，0 和 1 代表是否持有股票。
此问题共 n × K × 2 种状态，全部穷举就能搞定。

for 0 <= i < n:
    for 1 <= k <= K:
        for s in {0, 1}:
            dp[i][k][s] = max(buy, sell, rest)
```

* Recurrence:
```
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
              max( 今天选择 rest,        今天选择 sell       )
```
解释：今天我没有持有股票，有两种可能，我从这两种可能中求最大利润：
1. 我昨天就没有持有，且截至昨天最大交易次数限制为 k；然后我今天选择 `rest`，所以我今天还是没有持有，最大交易次数限制依然为 k。
2. 我昨天持有股票，且截至昨天最大交易次数限制为 k；但是今天我 `sell` 了，所以我今天没有持有股票了，最大交易次数限制依然为 k。
    * 注意：条件2 为什么是 k: 只有买入之后才能卖出，也就是说，交易是从 `buy` 开始的. 而当前的行为是 `sell`, 所以反推至前一天要`buy`的话，最大交易次数限制为就要和今天的一样为 k

```
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
              max( 今天选择 rest,         今天选择 buy         )
```
解释：今天我持有着股票，最大交易次数限制为 k，那么对于昨天来说，有两种可能，我从这两种可能中求最大利润：
1. 我昨天就持有着股票，且截至昨天最大交易次数限制为 k；然后今天选择 `rest`，所以我今天还持有着股票，最大交易次数限制依然为 k。
2. 我昨天本没有持有，且截至昨天最大交易次数限制为 k - 1；但今天我选择 `buy`，所以今天我就持有股票了，最大交易次数限制为 k。

**这里着重提醒一下**，时刻牢记「状态」的定义，状态 k 的定义并不是「已进行的交易次数」，而是「最大交易次数的上限限制」。如果确定今天进行一次交易，且要保证截至今天最大交易次数上限为 k，那么昨天的最大交易次数上限必须是 k - 1。

* Base Cases
```
dp[0][1...k][0] = 0
解释：因为 i 是从 0 开始的，所以 i = 0 意味着还没有开始，这时候的利润当然是 0。

dp[0][1...k][1] = -prices[0]
解释：第一天只买入股票，所以利润是负的 -prices[0]

dp[i][0][0] = 0
解释：因为 k 是从 1 开始的，所以 k = 0 意味着根本不允许交易，这时候利润当然是 0。

dp[i][0][1] = -infinity
解释：不允许交易的情况下，是不可能持有股票的。
因为我们的算法要求一个最大值，所以初始值设为一个最小值，方便取最大值。
```
* 为什么使用 `k = max_k, k--` 的方式呢？因为这样符合语义：
    * `k--`/`k++`都是正确的
    * 你买股票，初始的「状态」是什么？应该是从第 0 天开始，而且还没有进行过买卖，所以最大交易次数限制 `k` 应该是 `max_k`；而随着「状态」的推移，你会进行交易，那么交易次数上限 `k` 应该不断减少，这样一想，`k = max_k, k--` 的方式是比较合乎实际场景的。
