# [1382. Balance a Binary Search Tree](https://leetcode.com/problems/balance-a-binary-search-tree/description/)

## Solution idea

### 构造BST
* Step 1: 中序遍历获取 BST 的有序结果
* Step 2: 然后用[108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)的解法代码，将这个有序数组转化成平衡 BST

Time complexity = $O(n) + O(\log n) = O(n)$

## Resrouce
[代码随想录-1382.将二叉搜索树变平衡](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/1382.%E5%B0%86%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E5%8F%98%E5%B9%B3%E8%A1%A1.md)