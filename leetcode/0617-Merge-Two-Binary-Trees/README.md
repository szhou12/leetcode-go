# [617. Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/)

## Solution idea

### Post-order Traversal

* Base Cases: (和`LCA`类型题有点像)
    1. 两棵树的当前node都为空
    2. tree1的当前node 不为空，tree2的当前node 为空: return tree1的当前node
    3. tree1的当前node 为空，tree2的当前node 不为空: return tree2的当前node

* 后序遍历
    * 同时递归两棵树左孩子
    * 同时递归两棵树右孩子
    * 当前层:
        * 两棵树node的值求和
        * 重新链接返回的左、右孩子

Time complexity = $O(n)$

## Resource
[代码随想录-617.合并二叉树](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0617.%E5%90%88%E5%B9%B6%E4%BA%8C%E5%8F%89%E6%A0%91.md)