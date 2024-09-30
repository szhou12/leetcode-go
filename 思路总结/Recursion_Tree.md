# Recursion & Tree

* 遇到 Tree 的题，首先思考能否用 Recursion 解决，因为, Tree 天生适合 Recursion

* Tree Recursion 三段式结构
```go
func treeRecursion(TreeNode x, ...) ... {
    /*** Base Case(s) ***/
    [...]

    /*** Recursion Call ***/
    left := treeRecursion(x.Left, ...)
    right := treeRecursion(x.Right, ...)


    /*** 当前层 Return ***/
    [...]
    return ....
    
}
```

## 由Traversal Order 构造 Binary Tree

思路都差不多:

1. 找到root (preorder的头部, postorder的尾部)

2. 在inorder中定位root的index, 并进行左右分割

* 用中序遍历+后序遍历建二叉树：[106. Construct Binary Tree from Inorder and Postorder Traversal](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0106-Construct-Binary-Tree-from-Inorder-and-Postorder-Traversal)

* 用中序遍历+前序遍历建二叉树：[105. Construct Binary Tree from Preorder and Inorder Traversal](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0105-Construct-Binary-Tree-from-Preorder-and-Inorder-Traversal)




### **构造 Binary Search Tree**

* Convert ascending array to BST: [108. Convert Sorted Array to Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0108-Convert-Sorted-Array-to-Binary-Search-Tree)

* 平衡 BST: [1382. Balance a Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1382-Balance-a-Binary-Search-Tree)




## 找到所有 Unique Binary Search Tree

* Recursion to print all BST: [95. Unique Binary Search Trees II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0095-Unique-Binary-Search-Trees-II)

* DP to find #: [96. Unique Binary Search Trees](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0096-Unique-Binary-Search-Trees)




## Complete Binary Tree 的性质

* Complete Binary Tree一定会有一个 full binary subtree (满二叉树)

* [222. Count Complete Tree Nodes](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0222-Count-Complete-Tree-Nodes)




## 找Tree的所有路径

* [257. Binary Tree Paths](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0257-Binary-Tree-Paths)
    * 前序遍历: 因为这样才方便让父节点指向孩子节点，找到对应的路径。
    * 在这道题目中将第一次涉及到回溯，因为我们要把路径记录下来，需要回溯来回退 以便于 从一个路径在进入另一个路径。
        * **注意！！！**在Go语言的实现中，回溯反而不能写出来，否则会出错。
        * 我的理解是：Go里的append操作相当于是一个copy (ie. 切片的地址变了)，所以下一层返回当前层时，当前层的path不会包含下一层的信息 (因为下一层的path是另一个地址的)
    * 据说是Google面试题




## Post-Order Traversal 需要脑筋拐个弯的题

* Left Leaf: [404. Sum of Left Leaves](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0404-Sum-of-Left-Leaves)
    * 突破点：明确Left Leaf的定义, 根据定义来写当前层的逻辑
    * 判断当前节点是不是左叶子是无法判断的，必须要通过节点的父节点来判断其左孩子是不是左叶子

* Delete Node in BST: [450. Delete Node in a BST](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0450-Delete-Node-in-a-BST)
    * 突破点 & 难点：分类讨论，明确各种情况，对症下药
    * 情况一: 当前节点的**左孩子**为空, 直接返回右孩子
    * 情况二: 当前节点的**右孩子**为空, 直接返回左孩子
    * 情况三: 当前节点的**左、右孩子**都非空, 把左子树挂到右子树最左边的叶子节点下面

* Merge & Overlap two trees: [617. Merge Two Binary Trees](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0617-Merge-Two-Binary-Trees)

* Maximum Binary Tree: [654. Maximum Binary Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0654-Maximum-Binary-Tree)

* 修剪 BST: [669. Trim a Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0669-Trim-a-Binary-Search-Tree)
    * 看到BST本能地想用In-order
    * 但本题BST性质的运用是在于如何砍掉合适的node, recursion用Post-order就可以解

* Set Camera: [968. Binary Tree Cameras](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0968-Binary-Tree-Cameras)
    * 明确划分四大状态，以及伴生的16 cases
    * 用四个优先级把16 cases分门别类





## In-Order Traversal

* :lock: :red_circle: BST上找与target最近的k个nodes: [272. Closest Binary Search Tree Value II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0272-Closest-Binary-Search-Tree-Value-II)
    * In-Order Traversal + Queue

* BST上In-Order Traversal模拟一维递增数组: [501. Find Mode in Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0501-Find-Mode-in-Binary-Search-Tree)

* BST上In-Order Traversal 先走右子树，再走左子树 得到降序排列: [538. Convert BST to Greater Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0538-Convert-BST-to-Greater-Tree)

## Binary Search Tree (BST) 题
* 都是明示要利用BST性质的题
* 构造 BST:
    * Convert ascending array to BST: [108. Convert Sorted Array to Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0108-Convert-Sorted-Array-to-Binary-Search-Tree)
    * 平衡 BST: [1382. Balance a Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1382-Balance-a-Binary-Search-Tree)
* Validate BST: [98. Validate Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0098-Validate-Binary-Search-Tree)
* Search Node in BST: [700. Search in a Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0700-Search-in-a-Binary-Search-Tree)
* Delete Node in BST: [450. Delete Node in a BST](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0450-Delete-Node-in-a-BST)
* Insert Node in BST: [701. Insert into a Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0701-Insert-into-a-Binary-Search-Tree)
* Trim BST: [669. Trim a Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0669-Trim-a-Binary-Search-Tree)
* BST上In-Order Traversal模拟一维递增数组: [501. Find Mode in Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0501-Find-Mode-in-Binary-Search-Tree)
* BST上In-Order Traversal 先走右子树，再走左子树 得到降序排列: [538. Convert BST to Greater Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0538-Convert-BST-to-Greater-Tree)
* Recursion to print all BST: [95. Unique Binary Search Trees II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0095-Unique-Binary-Search-Trees-II)
* DP to find #: [96. Unique Binary Search Trees](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0096-Unique-Binary-Search-Trees)


## LCA
* LCA in BST: [235. Lowest Common Ancestor of a Binary Search Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0235-Lowest-Common-Ancestor-of-a-Binary-Search-Tree)

* LCA I: [236. Lowest Common Ancestor of a Binary Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0236-Lowest-Common-Ancestor-of-a-Binary-Tree)
    * input nodes `p`, `q` **一定**存在于tree中

* LCA II: [1644. Lowest Common Ancestor of a Binary Tree II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1644-Lowest-Common-Ancestor-of-a-Binary-Tree-II)
    * input nodes `p`, `q` **不一定**存在于tree中

* LCA III: [1650. Lowest Common Ancestor of a Binary Tree III](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1650-Lowest-Common-Ancestor-of-a-Binary-Tree-III)
    * Tree node 有 parent指针，直接逆流而上，记录"沿途风景"

* LCA IV: [1676. Lowest Common Ancestor of a Binary Tree IV](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1676-Lowest-Common-Ancestor-of-a-Binary-Tree-IV)
    * input nodes array **一定**存在于tree中
    * 额外使用一个 **HashMap** 充当 input nodes的set

* 最深叶子节点的LCA: [1123. Lowest Common Ancestor of Deepest Leaves](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1123-Lowest-Common-Ancestor-of-Deepest-Leaves)

* 二叉树中连环，计算环的长度: [2509. Cycle Length Queries in a Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2509-Cycle-Length-Queries-in-a-Tree)
    * 与一般的 LCA 的题型要简单，不一定非要用recursion，因为任意node的parent可以计算得到 `parent(X) = X/2`


## Find Paths in Tree
* :red_circle: 找到所有可以形成回文串的路径: [2791. Count Paths That Can Form a Palindrome in a Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2791-Count-Paths-That-Can-Form-a-Palindrome-in-a-Tree)

## Digit Counting & Finding
* :red_circle: 找到字符串倍增游戏中的第k个字符: [3307. Find the K-th Character in String Game II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3307-Find-the-K-th-Character-in-String-Game-II)