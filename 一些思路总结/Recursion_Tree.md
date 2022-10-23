# Recursion & Tree

* 遇到 Tree 的题，首先思考能否用 Recursion 解决，因为, Tree 天生适合 Recursion

## 由Traversal Order 构造 Binary Tree

思路都差不多:

1. 找到root (preorder的头部, postorder的尾部)

2. 在inorder中定位root的index, 并进行左右分割

* [106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)

* [105. Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

## 找到所有 Unique Binary Search Tree

* Recursion to print all BST: [95. Unique Binary Search Trees II](https://leetcode.com/problems/unique-binary-search-trees-ii/)

* DP to find #: [96. Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/)

## Complete Binary Tree 的性质

* Complete Binary Tree一定会有一个 full binary subtree (满二叉树)

* [222. Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes/)

## 找Tree的所有路径

* [257. Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/)
    * 前序遍历: 因为这样才方便让父节点指向孩子节点，找到对应的路径。
    * 在这道题目中将第一次涉及到回溯，因为我们要把路径记录下来，需要回溯来回退 以便于 从一个路径在进入另一个路径。
        * **注意！！！**在Go语言的实现中，回溯反而不能写出来，否则会出错。
        * 我的理解是：Go里的append操作相当于是一个copy (ie. 切片的地址变了)，所以下一层返回当前层时，当前层的path不会包含下一层的信息 (因为下一层的path是另一个地址的)
    * 据说是Google面试题

## Post-Order Traversal 需要脑筋拐个弯的题

* Left Leaf: [404. Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/)
    * 突破点：明确Left Leaf的定义, 根据定义来写当前层的逻辑
    * 判断当前节点是不是左叶子是无法判断的，必须要通过节点的父节点来判断其左孩子是不是左叶子

* Delete Node in BST: [450. Delete Node in a BST](https://leetcode.com/problems/delete-node-in-a-bst/)
    * 突破点 & 难点：分类讨论，明确各种情况，对症下药
    * 情况一: 当前节点的**左孩子**为空, 直接返回右孩子
    * 情况二: 当前节点的**右孩子**为空, 直接返回左孩子
    * 情况三: 当前节点的**左、右孩子**都非空, 把左子树挂到右子树最左边的叶子节点下面

## In-Order Traversal

* BST上In-Order Traversal模拟一维递增数组: [501. Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/)


