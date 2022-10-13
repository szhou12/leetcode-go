# [404. Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/)

## Solution idea

### Post-Order Traversal

* 本题难点：在于首先要明确Left Leaf的定义，然后根据定义来确定当前层要做些什么.

* Left Leaf的定义: 节点A的左孩子不为空，且左孩子的左右孩子都为空（说明是叶子节点），那么A节点的左孩子为左叶子节点

* 所以，判断当前节点是不是左叶子是无法判断的，必须要通过节点的父节点来判断其左孩子是不是左叶子。

Time complexity = $O(n)$

## Resource

[代码随想录-404.左叶子之和](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0404.%E5%B7%A6%E5%8F%B6%E5%AD%90%E4%B9%8B%E5%92%8C.md)