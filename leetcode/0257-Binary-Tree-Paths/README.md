# [257. Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/)


## Solution iea

### Tree Pre-Order Traversal

* 前序遍历：当前层先记录当前节点的值，再递归左子树、右子树

* 终止条件：与一般终止在空节点不同，此题终止节点为 leaf node

* 用一个数组 path 来记录走过的路径

## Resource
[代码随想录-257. 二叉树的所有路径](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0257.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%89%80%E6%9C%89%E8%B7%AF%E5%BE%84.md)