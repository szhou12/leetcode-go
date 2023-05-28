# [2516. Take K of Each Character From Left and Right](https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/description/)

## Solution idea
### Two Pointers - Sliding Window
#### 思路总结
1. 需要转个弯的Sliding Window的题型。使用 **模版Flex**。
2. 两个难点:
    1. 转化思想: 题目求两侧尽可能短的满足要求元素个数的substring，取反，也就是，求中间滑窗 a. 尽可能长 b. 包含的元素个数不超过一个阈值 (某一个元素的总数-滑窗内该元素个数 $\geq $ k个)。这样，就把题目转换成了 Sliding Window 的题目，并可套用模版。
    2. update result时一定！一定！一定！要分情况讨论，这个分情况讨论是指，right停下来是因为什么？是因为right越界了？还是因为拓展右边界的条件不再满足了？还是因为两者都不满足了？因为右边界延伸的代码逻辑是：当前right指向的元素记录入滑窗中，然后right前进一格。
        1. 如果情况 只是right越界了: 说明此时滑窗`[left, right-1]`内都是满足要求的
        2. 如果情况 只是拓展右边界的条件不再满足了: 说明是添加当前right指向元素导致的，然后right还又前进了一格，所以此时只有`[left, right-2]`内都是满足要求的
        3. 如果情况 两者都有: 也说明添加当前right指向元素也不行，所以此时只有`[left, right-2]`内都是满足要求的。

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 2516. Take K of Each Character From Left and Right](https://www.youtube.com/watch?v=KvQK3RXuTFc&ab_channel=HuifengGuan)