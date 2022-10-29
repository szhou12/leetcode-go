# [654. Maximum Binary Tree](https://leetcode.com/problems/maximum-binary-tree/)

## Solution idea

### Binary Tree - Post-order Traversal

* 当前层先找到input nums里的最大的值和它的index
* 左子树构建需要 `nums[:index]`
* 右子树构建需要 `nums[index+1:]`
* 用input nums里的最大的值构建当前层的node 并 链接上左、右子树

Time complexity = $O(n^2)$ because traversing every node needs $O(n)$ and constructing each node will loop `nums` by $O(n)$

## Resource
[东哥带你刷二叉树（构造篇）](https://labuladong.github.io/algo/2/21/38/)

[代码随想录-654.最大二叉树](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0654.%E6%9C%80%E5%A4%A7%E4%BA%8C%E5%8F%89%E6%A0%91.md)