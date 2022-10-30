# [700. Search in a Binary Search Tree](https://leetcode.com/problems/search-in-a-binary-search-tree/)

## Solution idea

* 又是一道Binary Search Tree的题，明示了要利用BST性质解题

* 如何利用BST性质？
    * 当遇到空节点, 说明要找的节点tree里没有, 返回空
    * 当前节点就是要找的, 返回当前节点
    * 如果当前节点的值 > target value: 目标节点只可能出现在左子树 (根据BST性质, 右子树都 > target)
    * 如果当前节点的值 < target value: 目标节点只可能出现在右子树 (根据BST性质, 左子树都 < target)

Time complexity = $O(\log n)$ because every recursion only traveres half of tree