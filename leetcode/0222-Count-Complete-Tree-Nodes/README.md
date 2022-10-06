# [222. Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes/)

## Solution idea

### 普通二叉树的解法

Post-Order Traversal: 左子树汇报，右子树汇报，当前层干点什么

Time complexity = $O(n)$

### 利用性质

Complete Binary Tree一定会有一个 full binary subtree (满二叉树)

Time complexity $< O(n)$

## Resource

[代码随想录-222.完全二叉树的节点个数](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0222.%E5%AE%8C%E5%85%A8%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E8%8A%82%E7%82%B9%E4%B8%AA%E6%95%B0.md)