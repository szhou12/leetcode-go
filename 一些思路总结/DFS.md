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

## 岛屿沉没类 - 找neighbor
* Roblox - Candy Crush

* 四面八方, 潮水涌来 [417. Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)

* [200. Number of Islands](https://leetcode.com/problems/number-of-islands/submissions/)

* [130. Surrounded Regions](https://leetcode.com/problems/surrounded-regions/)

* [1020. Number of Enclaves](https://leetcode.com/problems/number-of-enclaves/)

## All Subsets类

### Subsets

* 基础题 [78. Subsets](https://leetcode.com/problems/subsets/)
    * 有两种写法, **都要掌握！！！**
    * 写法一 适合[90. Subsets II](https://leetcode.com/problems/subsets-ii/)
    * 写法二 适合[491. Increasing Subsequences](https://leetcode.com/problems/increasing-subsequences/)

* 长度为k [77. Combinations](https://leetcode.com/problems/combinations/)

* 有重复元素 - Type I [90. Subsets II](https://leetcode.com/problems/subsets-ii/)
    * **Sort** 去重

* 有重复元素 - Type II [491. Increasing Subsequences](https://leetcode.com/problems/increasing-subsequences/)
    * 当前层 **Map** 去重. 注意：是每一层一个新 Map, 而不是一个 全局Map 因为层与层之间的相同元素**不能**去重

* 找IP地址 [93. Restore IP Addresses](https://leetcode.com/problems/restore-ip-addresses/)

### Combination Sum
* 0039, 0040解法属于一类；0077, 0216解法属于一类

* [39. Combination Sum](https://leetcode.com/problems/combination-sum/)

* [40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii/)

* [77. Combinations](https://leetcode.com/problems/combinations/)

* [216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii/)

## 树的深搜遍历

* 所有路径和: [129. Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers/)

## 行程规划

* [332. Reconstruct Itinerary](https://leetcode.com/problems/reconstruct-itinerary/)
    * 较难的DFS变体, 
    * 难点一：首先要想到可以用DFS暴力解
    * 难点二：如何记录映射关系 - 一个机场映射多个机场，机场之间要靠字母序排列
    * 难点三：Base Case 终止条件
    * 难点四：DFS function 返回值 bool 的意义
