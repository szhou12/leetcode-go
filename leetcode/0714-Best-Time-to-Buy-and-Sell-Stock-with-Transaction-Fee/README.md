# [714. Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)

## Solution idea

### DP

[188. Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/) `k=+infinity`的情况

1. Definition: 
    - `dp[i][0] :=` max profit for NOT holding the stock on i-th day
    - `dp[i][1] :=` max profit for holding the stock on i-th day
2. Base case:
    - `dp[0][0] = 0`
    - `dp[0][1] = -prices[0]-fee` 理解成第0天“贷款”买入股票
3. Recurrence
    - `dp[i][0] = max(dp[i-1][1] + prices[i], dp[i-1][0])` 第i天需要不再持有股票：要么第i-1天还持有股票并在第i天卖出；要么第i-1天就已经不持有股票，直接继承。
    - `dp[i][1] = max(dp[i-1][0] - prices[i]-fee, dp[i-1][1])` 第i天需要持有股票：要么第i-1天不持有股票并在第i天买入，缴纳fee；要么第i-1天就已经持有股票，直接继承

Time complexity = $O(n)$