# [1052. Grumpy Bookstore Owner](https://leetcode.com/problems/grumpy-bookstore-owner/description/)

## Solution idea

### Two Pointers: Sliding Window
* **吃了吐**的情况分析:
    * 吃进去1: 增加吃进去的值
    * 吃进去0: 不变
    * 吐出来1: 减去吐出去的值
    * 吐出来0: 不变

面试时曾经有一道极为类似的题: 股票buy & sell, 选择 k consecutive days to sell, 求最大收益
```
prices = [3, 1, 2, 4]
algo = [0, 1, 1, 0] // 0-sell, 1-buy
k = 2
```

Time complexity = $O(n)$

## Resource
[【每日一题】1052. Grumpy Bookstore Owner, 9/16/2020](https://www.youtube.com/watch?v=qxwjueQh82M&ab_channel=HuifengGuan)