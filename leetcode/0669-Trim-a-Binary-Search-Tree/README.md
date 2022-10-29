# [669. Trim a Binary Search Tree](https://leetcode.com/problems/trim-a-binary-search-tree/)

## Solution idea

### Post-order Traversal

* 这道题给的是BST, 所以会自然让人想用 In-Order Traversal解
* 但其实不是, 这道题BST性质的运用是在左、右子树和头节点的关系上
    * 左子树所有node的值 小于 头节点
    * 右子树所有node的值 大于 头节点
* 用Post-Order Traversal 就可以解
* 定义: 返回的是 以输入头节点为子树经过trim后返回的头节点
* recursion: 找到root左子树经过trim后返回的头节点，找到root右子树经过trim后返回的头节点
* 当前层: 
    * 如果root在合法范围内, 需要保留root, 只需要让root重新链接左、右子树的新头节点 并返回root
    * 如果root不在合法范围内:
        1. Case 1: 如果 root 小于 合法范围，利用BST性质，root左子树都小于合法范围，砍掉左子树以及root，返回右子树的新头节点
        2. Case 2: 如果 root 超过 合法范围，利用BST性质，root右子树都超过合法范围，砍掉右子树以及root，返回左子树的新头节点

Time complexity = $O(n)$

## Resource
[代码随想录-669. 修剪二叉搜索树](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0669.%E4%BF%AE%E5%89%AA%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91.md)
