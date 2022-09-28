# [96. Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/)

## Solution idea

## DP

$DP[i] = $ number of BSTs when there are $i$ number of nodes

Base Cases:

$DP[0] = 1$

$DP[1] = 1$

Recurrence:

$DP[i] = \sum_{j=1}^i DP[j-1] * DP[i-j]$

如何理解 $DP[j-1] * DP[i-j]$:

* $DP[j-1]$: 根结点`value = j`时，左子树的BST数量

* $DP[i-j]$: 根结点`value = j`时，右子树的BST数量

* 注意，node value 和 node数量 是重合的，所以用同样的 `i, j` 表示

```
node values:
1, ..., j-1, j,  j+1, ..., i
|--------|  root |---------|
 j-1 nodes         i-j nodes
```

Time complexity = $O(n^2)$

## Resource

[代码随想录-096](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0096.%E4%B8%8D%E5%90%8C%E7%9A%84%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91.md)