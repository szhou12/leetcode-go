# [108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)

## Solution idea

**Recursion**

Input array 升序排列，而BST in-order traversal产生升序排列的数组.

故，**Key Oberservation**: input array的mid就是BST的根结点

Time complexity = $O(\log n)$