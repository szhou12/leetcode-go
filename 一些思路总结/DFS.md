# Depth First Search

## N-皇后

* 求所有解法 [51. N-Queens](https://leetcode.com/problems/n-queens/)

* 求解法数量 [52. N-Queens II](https://leetcode.com/problems/n-queens-ii/)

* 究极进化 - 解数独 [37. Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)

**思路**

```
# levels = # rows
# branches = # cols
```

一个可行解的表示方法：`array = {index: row index; value: column index to place Queen}`

不合法放置Queen的三种情况：

1. 同列: column index = previous column index

2. slope = 1

3. slope = -1

## 岛屿沉没类

* 四面八方, 潮水涌来 [417. Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)

* [200. Number of Islands](https://leetcode.com/problems/number-of-islands/submissions/)

* [130. Surrounded Regions](https://leetcode.com/problems/surrounded-regions/)

* [1020. Number of Enclaves](https://leetcode.com/problems/number-of-enclaves/)

## All Subsets

* 基础题 [78. Subsets](https://leetcode.com/problems/subsets/)

* 长度为k [77. Combinations](https://leetcode.com/problems/combinations/)

* 找IP地址 [93. Restore IP Addresses](https://leetcode.com/problems/restore-ip-addresses/)

## 树的深搜遍历

* 所有路径和: [129. Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers/)