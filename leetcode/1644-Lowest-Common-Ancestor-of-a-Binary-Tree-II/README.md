# [1644. Lowest Common Ancestor of a Binary Tree II](https://leetcode.ca/all/1644.html)

## Solution idea

### LCA
* LCA 经典题型: 两个input nodes `p, q` **不一定**存在于树中
    * 与 `0236` 对比着学习

* 用一个 helper function, return (LCA, count)
    * return 的物理意义: subtree中是否找到LCA, subtree中看到了几个input nodes

* Base Cases:
    * current node 已经为空，说明input node一个没找到，老老实实返回 `(nil, 0)`
    * **注意: 如果遇到root等于`p`或`q`的时候，是不能和`0236`一样去立刻返回root的，因为不能判断另一个节点是否在树中。这里采取的方法是即使遇到root等于`p`或`q`的时候，也继续往下遍历左右子树.**

* Recursion call:
    * 问问看: 是否能在左子树中找到LCA, 左子树中看到了几个input nodes (记为 `lcount`)
    * 问问看: 是否能在右子树中找到LCA, 右子树中看到了几个input nodes (记为 `rcount`)

* Return 的情况
    1. 如果root等于`p`或`q`: 返回 (root, 1+lcount+rcount)
    2. 如果左子树和右子树返回的都不是空节点: `p`和`q`必定都在树中, 且都被找到了, 返回 (root, 2)
    3. 如果一边的子树找到了, 而另一边没找到 (i.e. 返回的是空节点), 说明: current node 不是 LCA, 返回找到(非空)节点的 (子树的结果, 子树的count)

Time complexity = $O(h)$

## Resource
code参考: [LeetCode 1644. Lowest Common Ancestor of a Binary Tree II](https://zhenchaogan.gitbook.io/leetcode-solution/leetcode-1644-lowest-common-ancestor-of-a-binary-tree-ii)
解释参考: [1644 - Lowest Common Ancestor of a Binary Tree II](https://leetcode.ca/2020-05-31-1644-Lowest-Common-Ancestor-of-a-Binary-Tree-II/)