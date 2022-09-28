# [95. Unique Binary Search Trees II](https://leetcode.com/problems/unique-binary-search-trees-ii/)

## Solution idea

### Recursion

1. 穷举 `root` 节点的所有可能

2. 递归构造出左、右子树的所有合法 BST

3. 给 `root` 节点穷举所有左、右子树的组合

Time complexity = $O(\frac{4^n}{n^{1/2}})$

## Resource

[Unique Binary Search Trees II](https://aaronice.gitbook.io/lintcode/dynamic_programming/unique-binary-search-trees-ii)