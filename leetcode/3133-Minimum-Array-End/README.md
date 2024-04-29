# [3133. Minimum Array End](https://leetcode.com/problems/minimum-array-end/description/)

## Solution idea
### Bit Manipulation + Greedy
- <逻辑思维题>
- **连续 bitwise AND 所得值一定越变越小，不会大于其中的最小元素。**
- 题目要求bitwise AND 以后的值=`x`, 并且`nums[]`是递增的，所以头号元素就是x
- `x`中为1的bit位，之后的元素对应的bit位都只能是1，但凡有0，就不可能最后AND完以后是x
- 有自由度(可以自由安排bit)的只有`x`中为0的bit位
- 题目要求`nums[]`最后一个元素最小，相当于要求整个`nums[]`最小(因为递增)
- 如何最小？
- `x`抽掉为1的bit位，剩余所有为0的bit位组合在一起=0
- 那么以0为开头的最小递增数组显然是: 0, 1, 2, 3, ..., n-1
- 也就是说，第二个元素在有自由度的bit位组合起来得1，第三个元素组合起来得2, ..., 最后一个元素组合起来得n-1

### 一图胜千言


Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 3133. Minimum Array End](https://www.youtube.com/watch?v=qjo7da3eNDQ&t=65s&ab_channel=HuifengGuan)