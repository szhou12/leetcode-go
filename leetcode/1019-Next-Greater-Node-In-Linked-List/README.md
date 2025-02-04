# [1019. Next Greater Node In Linked List](https://leetcode.com/problems/next-greater-node-in-linked-list/description/)

## Solution idea
### Stack - nextGreater
还是比较简单的，本质上还是用单调递减栈找Next Greater。只是套了一层Linked List的外壳。用$O(n)$时间遍历一遍Linked List，存入一个数组就行。然后就是套模版做就行。注意，最后输出的是具体的值，而不是index。

Time complexity = $O(n)$