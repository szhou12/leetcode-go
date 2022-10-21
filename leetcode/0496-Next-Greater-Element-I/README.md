# [496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/)

## Solution idea

### Naive approach - map
这道题的暴力解法很好想到，就是对每个元素后面都进行扫描，找到第一个更大的元素就行了

Time complexity = $O(n^2)$

### Stack

* 栈顶元素的物理意义: 维护栈顶元素, 使得它总是**大于**当前元素

Time complexity = $O(n)$

分析它的时间复杂度，要从整体来看：总共有 n 个元素，每个元素都被 push 入栈了一次，而最多会被 pop 一次，没有任何冗余操作。所以总的计算规模是和元素规模 n 成正比的，也就是 $O(n)$ 的复杂度。

## Resource
[单调栈结构解决三道算法题](https://labuladong.github.io/algo/2/23/63/)