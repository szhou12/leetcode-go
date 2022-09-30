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
    * DP: `k=1`
    * 倒着看，每次用看到的最大值 - 当前值
    
* [122. Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/)
    * DP: `k= +infinity`
    * Greedy Algorithm: 累加相邻两天的差为正数的case


