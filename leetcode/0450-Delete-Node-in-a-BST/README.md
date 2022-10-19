# [450. Delete Node in a BST](https://leetcode.com/problems/delete-node-in-a-bst/)

## Solution idea

### Post-Order Traversal
* **难点**: 分类讨论 所有可能遇到的删除节点的情况， 并保证所有情况下的操作都能保持BST特性
    * 情况一: 当前节点的**左孩子**为空, 直接返回右孩子
    * 情况二: 当前节点的**右孩子**为空, 直接返回左孩子
    * 情况三: 当前节点的**左、右孩子**都非空, 把左子树挂到右子树最左边的叶子节点下面

Time complexity = $O(n)$

## Resource

有很好的图示：[代码随想录-450.删除二叉搜索树中的节点](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0450.%E5%88%A0%E9%99%A4%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E4%B8%AD%E7%9A%84%E8%8A%82%E7%82%B9.md)
