# Recursion & Tree

* 遇到 Tree 的题，首先思考能否用 Recursion 解决，因为, Tree 天生适合 Recursion

## 由Traversal Order 构造 Binary Tree

思路都差不多:

1. 找到root (preorder的头部, postorder的尾部)

2. 在inorder中定位root的index, 并进行左右分割

* [106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

* [105. Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)


