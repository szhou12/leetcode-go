# [84. Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/)

## Solution idea

**核心思路**: 对任一 `bar i`, 以它的高为限, 能得到最大的面积. 既然高度限制在`heights[i]`, 那么，只需要中心开心，往两边找分别遇上比 `bar i` 矮的就停下，这样，我们就有了以高为`heights[i]` 能找到最大的宽。所求面积就是`bar i`的最大面积.

### DP

遵循**核心思路**

Time complexity = $O(n^2)$

### 单调栈 Monotonic Stack

同样遵循**核心思路**, 只是把 DP 里的记事本 用 Stack 来实现

## Resource

[【每日一题】LeetCode 84. Largest Rectangle in Histogram](https://www.youtube.com/watch?v=mesaogfSjD4&ab_channel=HuifengGuan)

[代码随想录-84.柱状图中最大的矩形](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0084.%E6%9F%B1%E7%8A%B6%E5%9B%BE%E4%B8%AD%E6%9C%80%E5%A4%A7%E7%9A%84%E7%9F%A9%E5%BD%A2.md)