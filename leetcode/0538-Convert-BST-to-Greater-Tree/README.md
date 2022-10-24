# [538. Convert BST to Greater Tree](https://leetcode.com/problems/convert-bst-to-greater-tree/)

## Solution idea

### BST - In-Order Traversal

* 中序遍历: 先走右子树，再当前层，最后走左子树 --> 可以得到降序排列

* 本题要求: 每一个node value 累加它右子树的和, 所以**先走右子树**是符合题意的做法
    * 相当于, 从左至右loop一个降序排列的数组, 每个元素累加之前所有元素的sum

Time complexity = $O(n)$

## Resource

[东哥带你刷二叉搜索树（特性篇）](https://labuladong.github.io/algo/2/21/42/)