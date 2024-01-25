# [3007. Maximum Number That Sum of the Prices Is Less Than or Equal to K](https://leetcode.com/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/description/)

## Solution idea
### Binary Search + Digital Counting
1. Binary Search: 首先要有直觉，题目具有单调性：num越大，price越高。所以可以使用 Binary Search 来猜答案。
2. Digital Counting: 本题难点在于对于任意一个猜的`mid`，折半的条件是什么？或者说，对应的`price`怎么计算？这里使用 Digital Counting。有两种方法，下面会分别讨论。

#### Digital Counting - Method 1
1. 每猜一个`num`，从 1～`num` 每一个自然数都分别计算 bit 1 的个数比较笨拙。Method 1的方法是：与其每一个数计算有多少个 bit 1，逆向思维，对于`num`的每一个要求的bit位`i`，计算有多少个组合数满足：1. bit位`i=1`; 2. 这个组合的十进制表示数值`<= num`。

##### Digital Counting - Method 2

## Resource
[【每日一题】LeetCode 3007. Maximum Number That Sum of the Prices Is Less Than or Equal to K](https://www.youtube.com/watch?v=tw6jJCIq0lU&t=20s&ab_channel=HuifengGuan)

[Leetcode Top Solution](https://leetcode.com/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/solutions/4563689/c-solution-bit-manipulation-binary-search-comprehensive-visualization)