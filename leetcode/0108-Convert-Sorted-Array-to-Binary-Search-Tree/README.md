# [108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)

## Solution idea

**Recursion**

Input array 升序排列，而BST in-order traversal产生升序排列的数组.

故，**Key Oberservation**: input array的mid就是BST的根结点

Time complexity = $O(\log n)$

## Resource

[代码随想录-108.将有序数组转换为二叉搜索树](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0108.%E5%B0%86%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E8%BD%AC%E6%8D%A2%E4%B8%BA%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91.md)