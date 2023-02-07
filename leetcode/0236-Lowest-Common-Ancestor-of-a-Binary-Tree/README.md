# [236. Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)

## Solution idea

### LCA

* LCA 经典题型: 两个input nodes `p, q` **一定**存在于树中

* Base Cases:
    1. current node 已经为空，说明input node一个没找到，老老实实返回 `nil`
    2. current node 是input node任意一个，说明找到了一个input node, 返回current node

* Recursion
    1. 问问current node的左子树能不能找到任意一个input node啊？
    2. 问问current node的右子树能不能找到任意一个input node啊？

* **Recursion Return的物理意义**: current node的左、右子树 **是否找到任意一个 input node**
    1. 如果左子树找到了一个 input node, 右子树找到了另一个 input node, (i.e. 两个子树返回的都不是空节点) 说明: 好，current node 就是 LCA, 返回current node
        * 显然，如果左、右子树各找到了input node，它们找到的肯定不是同一个input node，因为一个node怎么可能同时存在于两个地方呢
    2. 如果一边的子树找到了, 而另一边没找到 (i.e. 返回的是空节点), 说明: current node 不是 LCA, 返回找到节点(非空)的子树的结果

Time complexity = $O(h)$ where $h = $ tree height.

## Note:
The testing file is not finished because inputs p, q are not actual nodes from input tree, which they should have been.
This will cause the algorithm always return nil.
BUT the algorithm is correct.