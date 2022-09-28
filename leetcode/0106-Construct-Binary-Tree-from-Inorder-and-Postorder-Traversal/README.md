# [106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

## Solution idea

### Recursion

* Step 1: Base case - 如果postorder数组大小为零的话，说明是空节点了。

* Step 2: 如果不为空，那么取后序数组最后一个元素作为 root (**Key Observation!!!**)

* Step 3: 找到后序数组最后一个元素在**中序数组的位置**，作为**切割点**

* Step 4: 切割中序数组，切成 左中序数组 和 右中序数组 （顺序别搞反了，一定是先切中序数组）

* Step 5: 切割后序数组，切成 左后序数组 和 右后序数组

* Step 6: 递归处理左区间和右区间

* Step 7: 链接root 和 左、右子树

Time complexity = $O(n)$ where $n$ number of nodes