# [2222. Number of Ways to Select Buildings](https://leetcode.com/problems/number-of-ways-to-select-buildings/description/)

## Solution idea
### DP - 子序列个数
1. 容易想到使用DP解题，并设计出三维`dp[i][j][k]`，因为选或者不选i，我们需要知道: 1. 已经选了多少元素？2. 上一次选的元素和当前这个是否是同一个类型？难点是：1. 如何逻辑通顺地定义DP; 2. 如何设计转移方程; 3. 如何确定Base Case的值(种子)。
2. 难点解析:
    1. DP定义: 
        - `dp[i][j][k]` is defined as after looking at `s[i]`, there have been `j (=0,1,2,3)` elements selected, and the last element is `k (=0,1)`. 截止到i，我们已经选了j个元素(注意: j个中包含考虑i号元素)，且上一个元素是k类型。
    2. 状态转移方程: 思路在于，在i这个时间节点上，如何操作i号元素？那么显然，操作无非两种 - 选 或者 不选 第i号元素。
        1. NOT select i: 直接继承i-1时的状态: `dp[i-1][j][k]`
        2. select i: “有条件地”继承i-1时预留一个空位的状态: `dp[i-1][j-1][1-k]`。条件为: 1. 合法。即 `j-1>=0`; 2. i-1号元素和i号元素不同类型，即`1-k` (trick: 通过 1-x 来进行0, 1取反)。
        - 两者相加，即为截止到i号元素，所有的选择方法 = `dp[i-1][j][k] + dp[i-1][j-1][1-k]`
    3. Base Cases:
        - 注意到，状态转移方程都是通过累加前面的状态得到。也就意味着，我们要在Base Case设置“种子”(不为0的值)，才能在后续中有效地累加。
        - 根据DP定义来设置"种子":
            - `dp[0][0][0] = 1`: 表示“截止到头元素，没有任何元素入选”。“什么都不选“也是一种valid的选择，所以设置为1。
            - `dp[0][0][1] = 1`: 意义同上。因为此时什么都没选，所以上一个元素是什么并没有意义。所以k=0和k=1都设置为1。
            - `dp[0][1][s[0]] = 1`: 表示“截止到头元素，选了一个元素 (就是这个头元素)，且k是这个唯一选择元素的值”。这是一种valid的选择，所以设置为1。
    4. 总数 = `dp[n-1][3][0] + dp[n-1][3][1]`。因为我们要选3个元素，且最后一个元素是0或者1。
3. 为方便定义Base Case，也可以使用trick: 前置一个虚空占位符 (`s = '#'+s`)。这样，只需要设置两个"种子"。
    - `dp[0][0][0] = 1`: 表示"截止到什么元素都还没有，不选任何元素"。“什么都不选“也是一种valid的选择，所以设置为1。
    - `dp[0][0][1] = 1`: 意义同上。因为此时什么都没选，所以上一个元素是什么并没有意义。所以k=0和k=1都设置为1。
    - 总数 = `dp[n][3][0] + dp[n][3][1]`。


Time complexity = $O(n*4*2)$ = $O(n)$

## Resource
[【每日一题】LeetCode 2222. Number of Ways to Select Buildings](https://www.youtube.com/watch?v=PI6Dwkdw5hQ&ab_channel=HuifengGuan)