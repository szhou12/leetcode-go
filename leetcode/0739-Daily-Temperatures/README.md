# [739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)

## Solution idea

### Stack

* 套用单调栈模版 - 从后往前看，挤掉栈顶矮个子!
* 这里，**Stack 内存入元素的index, 而不是元素本身!!!**
    * 原因: 题目要求算距离 ==> 算距离用index ==> 保存元素index

Time complexity = $O(n)$

Similar question: [503. Next Greater Element II](https://leetcode.com/problems/next-greater-element-ii/)

## Resource
[单调栈结构解决三道算法题](https://labuladong.github.io/algo/2/23/63/)