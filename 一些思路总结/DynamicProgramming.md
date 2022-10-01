# Dynamic Programming

## String Matching

一般是**二维DP**

* [72. Edit Distance](https://leetcode.com/problems/edit-distance/)

* [115. Distinct Subsequences](https://leetcode.com/problems/distinct-subsequences/)

## Longest Ascending Subarray 类型

* [485. Max Consecutive Ones](https://leetcode.com/problems/max-consecutive-ones/)

## 买卖股票

* [188. Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/)
    * 这种类型的通法

* [121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)
    * DP: `k=1`; 因为`k=0`已经在 base cases里，`k=1` 就这一种情况， 对状态转移已经没有影响了，不再需要描述状态 `k`了
    * 倒着看，每次用看到的最大值 - 当前值

* [123. Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/)
    * DP: `k= 2`
    
* [122. Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)
    * DP: `k= +infinity`; 因为 `k` 无限大，所以, `k`和`k-1`等同的，就不再需要描述状态 `k`了
    * Greedy Algorithm: 累加相邻两天的差为正数的case

* [309. Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)
    * DP: `k= +infinity`
    * 将`cooldown`融入状态转移方程。即，当天如果要买入，则需要往前看两天卖出的情况 (而不是前一天卖出的情况)
    * 同时，增加 `i==1`的 base cases

* [714. Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)
    * DP: `k= +infinity`
    * 将`fee`融入状态转移方程。即，当天如果要买入，则需要额外减掉`fee`

