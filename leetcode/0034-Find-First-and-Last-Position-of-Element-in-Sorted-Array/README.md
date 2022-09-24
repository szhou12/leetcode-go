# [34. Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

## Solution idea

非常经典的 **Binary Search** 类型题.

### 要点！！！

**一定！一定！一定要明确**target可能在数组中**没有出现**的三种情况，才能正确地写出post-processing中边界的查看条件!!!

* 情况一: target 不在数组范围中，target 过于小，在**左边界**以外. e.g. 数组{3, 4, 5}，target为2
* 情况二: target 不在数组范围中，target 过于大，在**右边界**以外，e.g. 数组{3, 4, 5}，target为6
* 情况三: target 在数组范围中，只是数组中不存在. e.g. 数组{3, 6, 7}, target为5

Time complexity = $O(\log n) + O(\log n) = O(\log n)$