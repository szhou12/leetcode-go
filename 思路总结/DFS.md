# Depth First Search

## 目录
* [N-Queens](#n-Queens)
* [岛屿沉没 - 找neighbor](#岛屿沉没---找neighbor)
* [类DP](#类dp)
* [Permutation](#permutation)
* [All Subsets](#all-subsets)
    * [Subsets](#subsets)
    * [Combination Sum](#combination-sum)
* [图的遍历](#图的遍历)
    * [树的深搜遍历](#树的深搜遍历)
* [行程规划](#行程规划)
* [Min-Max Strategy](#min-max-strategy)



## N-Queens

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



## 岛屿沉没 - 找neighbor

* 岛屿数量: [200. Number of Islands](https://leetcode.com/problems/number-of-islands/submissions/)

* 飞地的数量: [1020. Number of Enclaves](https://leetcode.com/problems/number-of-enclaves/)

* 封闭岛屿数量: [1254. Number of Closed Islands](https://leetcode.com/problems/number-of-closed-islands/)
    * **突破点**: 提前沉没掉**贴着上、下、左、右边界**的岛屿, 剩下的就是符合题意的岛屿
    * 提前沉没岛屿的姐妹题: [1905. Count Sub Islands](https://leetcode.com/problems/count-sub-islands/)

* 子岛屿数量: [1905. Count Sub Islands](https://leetcode.com/problems/count-sub-islands/)
    * **突破点**: 提前沉没掉**grid2中还是陆地但 grid1中已经是海水**的岛屿, 剩下的就是符合题意的岛屿
    * 提前沉没岛屿的姐妹题: [1254. Number of Closed Islands](https://leetcode.com/problems/number-of-closed-islands/)

* 最大岛屿面积: [695. Max Area of Island](https://leetcode.com/problems/max-area-of-island/)

* 最大人工岛屿面积: [827. Making A Large Island](https://leetcode.com/problems/making-a-large-island/)
    * 分两步: 1. DFS记录每个岛屿的面积; 2. 遍历所有为0的cell，尝试连接邻居岛屿，并统计连接所得最大面积

* 四面八方, 潮水涌来 [417. Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)
    * 上、下、左、右边界出发，分别向内做DFS

* [130. Surrounded Regions](https://leetcode.com/problems/surrounded-regions/)

* Roblox - Candy Crush

## 类DP

* :red_circle: 特殊全排列: [2741. Special Permutations](https://leetcode.com/problems/special-permutations/)

* :red_circle: 结合字符串: [2746. Decremental String Concatenation](https://leetcode.com/problems/decremental-string-concatenation/description/)


## Permutation

* :red_circle: 全排列: [46. Permutations](https://leetcode.com/problems/permutations/)
    * Permutation 类型题 **母题**

* :red_circle: 有重复元素的全排列: [47. Permutations II](https://leetcode.com/problems/permutations-ii/)




## All Subsets

### Subsets

* 基础题 [78. Subsets](https://leetcode.com/problems/subsets/)
    * 有两种写法, **都要掌握！！！**
    * 写法一 适合[90. Subsets II](https://leetcode.com/problems/subsets-ii/)
    * 写法二 适合[491. Increasing Subsequences](https://leetcode.com/problems/increasing-subsequences/)

* 长度为k [77. Combinations](https://leetcode.com/problems/combinations/)

* 有重复元素 - Type I [90. Subsets II](https://leetcode.com/problems/subsets-ii/)
    * **Sort** 去重
    * All Subsets 第一种写法

* 有重复元素 - Type II [491. Increasing Subsequences](https://leetcode.com/problems/increasing-subsequences/)
    * 当前层 **Map** 去重. 注意：是每一层一个新 Map, 而不是一个 全局Map 因为层与层之间的相同元素**不能**去重
    * All Subsets 第二种写法

* 找IP地址 [93. Restore IP Addresses](https://leetcode.com/problems/restore-ip-addresses/)

* 当前元素正负值 [494. Target Sum](https://leetcode.com/problems/target-sum/)

### Combination Sum
* 0039, 0040解法属于一类；0077, 0216解法属于一类

* [39. Combination Sum](https://leetcode.com/problems/combination-sum/)

* [40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii/)

* [77. Combinations](https://leetcode.com/problems/combinations/)

* [216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii/)



## 图的遍历
* DAG图中找从`node 0`到`node n-1`的所有合法路径: [797. All Paths From Source to Target](https://leetcode.com/problems/all-paths-from-source-to-target/description/)

* 判断在有向有环图中能否从`node 0`到达所有其他 nodes: [841. Keys and Rooms](https://leetcode.com/problems/keys-and-rooms/description/)
    * 用 hashmap `visited`记录遍历过的节点

### 树的深搜遍历

* 所有路径和: [129. Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers/)


## 行程规划

* [332. Reconstruct Itinerary](https://leetcode.com/problems/reconstruct-itinerary/)
    * 较难的DFS变体, 
    * 难点一：首先要想到可以用DFS暴力解
    * 难点二：如何记录映射关系 - 一个机场映射多个机场，机场之间要靠字母序排列
    * 难点三：Base Case 终止条件
    * 难点四：DFS function 返回值 bool 的意义


## Min-Max Strategy
* :red_circle: 先手能否赢: [464. Can I Win](https://leetcode.com/problems/can-i-win/description/)

* :red_circle: 石头游戏I: [877. Stone Game](https://leetcode.com/problems/stone-game/description/)

* :yellow_circle: 判断赢家: [486. Predict the Winner](https://leetcode.com/problems/predict-the-winner/description/)
    * 与 0877 完全一样的解法

* :yellow_circle: 石头游戏II: [1140. Stone Game II](https://leetcode.com/problems/stone-game-ii/description/)