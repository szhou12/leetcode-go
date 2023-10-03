# Dynamic Programming

## 目录
* [基础题](#基础题)
* [买卖股票](#买卖股票)
* [抢劫房子](#抢劫房子)
* [背包问题](#背包问题)
* [回文串](#回文串)
* [Subarray 子数组类型题](#subarray-子数组类型题)
    * [类型一: Longest Ascending Subarray - 最长递增子数组](#类型一-longest-ascending-subarray---最长递增子数组)
    * [类型二: Longest Common Subarray - 最长公共子数组](#类型二-longest-common-subarray---最长公共子数组)
    * [类型三: 最大子数组](#类型三-最大子数组)
* [Subsequnce 子序列类型题](#subsequnce-子序列类型题)
    * [类型一: Longest Increasing Subsequence - 最长递增子序列](#类型一-longest-increasing-subsequence---最长递增子序列)
    * [类型二: Longest Common Subsequence - 最长公共子序列](#类型二-longest-common-subsequence---最长公共子序列)
    * [类型三: String Matching + 编辑距离](#类型三-string-matching--编辑距离)
    * [类型四: 子序列个数](#类型四-子序列个数)
* [二维/高维 DP](#二维高维-dp)
    * [一维种类型](#一维种类型)
    * [区间型 (interval)](#区间型-interval)


## 基础题

* Fibonacci Number: [509. Fibonacci Number](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0509-Fibonacci-Number)

* 爬楼梯: [70. Climbing Stairs](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0070-Climbing-Stairs)

* 最小花费爬楼梯: [746. Min Cost Climbing Stairs](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0746-Min-Cost-Climbing-Stairs)

* 有限制条件的切字符串: [2522. Partition String Into Substrings With Values at Most K](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2522-Partition-String-Into-Substrings-With-Values-at-Most-K)
    * 切东西类型的DP题
    * 需要通过总结 `DP[i]` 的规律性质来优化算法的 time complexity


## 买卖股票

* 买卖股票4: [188. Best Time to Buy and Sell Stock IV](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0188-Best-Time-to-Buy-and-Sell-Stock-IV)
    * 这种类型的通法

* 买卖股票1: [121. Best Time to Buy and Sell Stock](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0121-Best-Time-to-Buy-and-Sell-Stock)
    * DP: `k=1`; 因为`k=0`已经在 base cases里，`k=1` 就这一种情况， 对状态转移已经没有影响了，不再需要描述状态 `k`了
    * 倒着看，每次用看到的最大值 - 当前值

* 买卖股票3: [123. Best Time to Buy and Sell Stock III](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0123-Best-Time-to-Buy-and-Sell-Stock-III)
    * DP: `k= 2`
    
* 买卖股票2: [122. Best Time to Buy and Sell Stock II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0122-Best-Time-to-Buy-and-Sell-Stock-II)
    * DP: `k= +infinity`; 因为 `k` 无限大，所以, `k`和`k-1`等同的，就不再需要描述状态 `k`了
    * Greedy Algorithm: 累加相邻两天的差为正数的case

* 买卖股票+冷静期: [309. Best Time to Buy and Sell Stock with Cooldown](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0309-Best-Time-to-Buy-and-Sell-Stock-with-Cooldown)
    * DP: `k= +infinity`
    * 将`cooldown`融入状态转移方程。即，当天如果要买入，则需要往前看两天卖出的情况 (而不是前一天卖出的情况)
    * 同时，增加 `i==1`的 base cases

* 买卖股票+交易费: [714. Best Time to Buy and Sell Stock with Transaction Fee](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0714-Best-Time-to-Buy-and-Sell-Stock-with-Transaction-Fee)
    * DP: `k= +infinity`
    * 将`fee`融入状态转移方程。即，当天如果要买入，则需要额外减掉`fee`




## 抢劫房子
* 抢房子1: [198. House Robber](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0198-House-Robber)

* 抢房子2: [213. House Robber II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0213-House-Robber-II)

* 抢房子3: [337. House Robber III](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0337-House-Robber-III)
    * 树形DP: post-order traversal + DP

* 抢房子4: [2560. House Robber IV](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2560-House-Robber-IV)
    * Binary Search 猜答案 + DP



## 背包问题
* [279. Perfect Squares](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0279-Perfect-Squares)

* [322. Coin Change](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0322-Coin-Change)

* [518. Coin Change II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0518-Coin-Change-II)
    * 题目要求 **组合数**
    * 只能外层遍历物品 (coins种类), 内层遍历背包容量 (目标金额)
    * 以防止 `{1, 5}`, `{5, 1}` 的情况, 他们是等价的combination

* [416. Partition Equal Subset Sum](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0416-Partition-Equal-Subset-Sum)
    * 0/1背包问题 (一个元素只能放入一次)

* [494. Target Sum](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0494-Target-Sum)
    * 0/1背包问题
    * 需要数学推导，不好想

* [474. Ones and Zeroes](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0474-Ones-and-Zeroes)
    * 0/1背包问题
    * 二维DP
    * 倒序遍历背包容量!!!




## 回文串

* [5. Longest Palindromic Substring](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0005-Longest-Palindromic-Substring)

* [647. Palindromic Substrings](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0647-Palindromic-Substrings)

* [131. Palindrome Partitioning](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0131-Palindrome-Partitioning)
    * DFS 解法容易想到

* [132. Palindrome Partitioning II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0132-Palindrome-Partitioning-II)
    * 两步 DP




## Subarray 子数组类型题

### 类型一: Longest Ascending Subarray - 最长递增子数组

* 最长递增子数组/最长连续递增子序列: [674. Longest Continuous Increasing Subsequence](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0674-Longest-Continuous-Increasing-Subsequence)
    * 经典题，需要明确区分与**最长递增子序列**在解法上的的异同

* 最长连续1: [485. Max Consecutive Ones](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0485-Max-Consecutive-Ones)

### 类型二: Longest Common Subarray - 最长公共子数组
* 与 **Longest Common Subsequence** 不同的处理在于: 当以`i`结尾的元素 与 以`j`结尾的元素不相同时，直接subarray长度为0, 不考虑其他.
* 最长公共子数组: [718. Maximum Length of Repeated Subarray](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0718-Maximum-Length-of-Repeated-Subarray)

### 类型三: 最大子数组
* 最大子数组: [53. Maximum Subarray](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0053-Maximum-Subarray)



## Subsequnce 子序列类型题

### 类型一: Longest Increasing Subsequence - 最长递增子序列
* LIS: [300. Longest Increasing Subsequence](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0300-Longest-Increasing-Subsequence)
    * 1维 LIS
    * 注意：求的是全局最长，不一定发生在最后

* 俄罗斯套娃信封: [354. Russian Doll Envelopes](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0354-Russian-Doll-Envelopes)
    * 2维 LIS
    * Sort rule: 因为题目只允许 每个维度 `<` (严格递增), Sort by `w` in increasing order; then if `w` identical, sort by `h` by **decreasing order**
    * 最优解：Greedy + Binary Search

* 堆积如山: [1691. Maximum Height by Stacking Cuboids](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1691-Maximum-Height-by-Stacking-Cuboids)
    * 3维 LIS
    * 题目允许rotate: 每个cuboid就会有最多6种变体需要考虑
    * Sort rule: 因为题目允许 每个维度 `<=`, sort by `width` in increasing order; then if `width` identical, sort by `length` by increasing order; ...
    * 加入index 信息，防止同一个cuboid因为idential变体被放入LIS超过1次
    * 最优解: 每个cuboid只能需要考虑1种变体 `(min(l, w, h), mid(l, w, h), max(l, w, h))`

* Longest Wiggle Subsequence: [376. Wiggle Subsequence](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0376-Wiggle-Subsequence)
    * 类二维DP: `DP[i][0, 1]`

* LIS个数: [673. Number of Longest Increasing Subsequence](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0673-Number-of-Longest-Increasing-Subsequence)
    * 1维 LIS
    * 维护两个记事本:
        * `dp[i]`: length of LIS ending at `i`
        * `count[i]`: number of LIS ending at `i`

### 类型二: Longest Common Subsequence - 最长公共子序列

**公共子序列**的一个重要特征: 相对顺序不改变. 意即, 数字4在字符串A中数字1的后面，那么数字4也应该在字符串B数字1的后面

* LCS: [1143. Longest Common Subsequence](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1143-Longest-Common-Subsequence)
* 只允许Delete转换两个string: [583. Delete Operation for Two Strings](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0583-Delete-Operation-for-Two-Strings)
    * 求出LCS, 然后减去LCS长度即为需要删除的次数
* 最小ASCII删除和: [712. Minimum ASCII Delete Sum for Two Strings](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0712-Minimum-ASCII-Delete-Sum-for-Two-Strings)
* 不想交的线的最大数量: [1035. Uncrossed Lines](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1035-Uncrossed-Lines)
    * "连线间互不相交" == **相对顺序不改变**
    * 根据这个重要信息，把题目转化为 LCS 题求解

### 类型三: String Matching + 编辑距离

一般是**二维DP**

* Edit Distance: [72. Edit Distance](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0072-Edit-Distance)
* 只允许Delete转换两个string: [583. Delete Operation for Two Strings](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0583-Delete-Operation-for-Two-Strings)
* 不同的子序列数量: [115. Distinct Subsequences](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0115-Distinct-Subsequences)

### 类型四: 子序列个数

* :red_circle: :secret: 不同非空子序列个数II: [940. Distinct Subsequences II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0940-Distinct-Subsequences-II)
    * *940 与 1987 思路非常相近*

* :red_circle: :secret: 二进制字符串中的"好"子序列个数: [1987. Number of Unique Good Subsequences](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1987-Number-of-Unique-Good-Subsequences)
    * *940 与 1987 思路非常相近*



## 二维/高维 DP

### 一维种类型
#### 第一维 = 位置 (index); 第二维 = 值的种类 (value)

* :red_circle: 奇偶变换数组中挑选总和最大化: [2786. Visit Array Positions to Maximize Score](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2786-Visit-Array-Positions-to-Maximize-Score)
    * `DP[index][value]`

* :red_circle: 最少次变换元素分三组: [2826. Sorting Three Groups](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2826-Sorting-Three-Groups)
    * `DP[index][value]`

### 区间型 (interval)
#### 第一维 = 区间起始点 (start); 第二维 = 区间终止点 (end) or offset

* 猜测数高了还是低了II: [375. Guess Number Higher or Lower II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0375-Guess-Number-Higher-or-Lower-II)
    * 方法1: **二维DP**. 外层循环length, 内层循环starting position
    * 方法2: DP Memoization (Recursion + DP), 用 DP memo 来剪枝
 
