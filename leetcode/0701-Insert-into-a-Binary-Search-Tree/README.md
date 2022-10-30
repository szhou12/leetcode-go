# [701. Insert into a Binary Search Tree](https://leetcode.com/problems/insert-into-a-binary-search-tree/)

## Solution idea

* 又是一道Binary Search Tree的题，明示了要利用BST性质解题

* 如何利用BST性质？
    * 当遇到空节点, 可以在这里安全插入新节点
    * 题目说了tree里没有和target相同的值的节点, 所以:
        1. 如果当前节点的值 > target value: 新节点要插入在左子树 (根据BST性质, 右子树都 > target, 不宜插入)
        2. 如果当前节点的值 < target value: 新节点要插入在右子树 (根据BST性质, 左子树都 < target, 不宜插入)

Time complexity = $O(\log n)$ because every recursion only traveres half of tree

## Resource
[东哥带你刷二叉搜索树（基操篇）](https://labuladong.github.io/algo/2/21/43/)