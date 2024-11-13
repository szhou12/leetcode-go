# [1526. Minimum Number of Increments on Subarrays to Form a Target Array](https://leetcode.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/description/)

## Solution idea
### Segment Tree

Time complexity = $O(n\log n)$

### Greedy
- 对于`target[]`任意一点i，如果需要“水涨船高”从0到`target[i]`，那么有两种情况: 1. 如果 i 号元素 比 i-1 号元素大，那么“接续”涨到i-1号元素的增量，继续增加到i号元素的值。如何选择subarray? 选择最大 j，使得 `[i ... j]` 中j号元素是i之后第一个遇见的小于等于i号元素的元素。2. 如果 i 号元素 比 i-1 号元素小，那么说明i号元素的值在之前某一次的subarray选择中涵盖到了，不需要额外增量，但是此时“水平面”需要退回到当前的i号元素的值，因为之后的增量是基于此时的i号元素的值。

Time complexity = $O(n)$

## Resource
[【每日一题】1526. Minimum Number of Increments on Subarrays to Form a Target Array, 3/31/2021](https://www.youtube.com/watch?v=LA8NMbeF4Xg&ab_channel=HuifengGuan)