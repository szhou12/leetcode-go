# [3165. Maximum Sum of Subsequence With Non-adjacent Elements](https://leetcode.com/problems/maximum-sum-of-subsequence-with-non-adjacent-elements/description/)

## Solution idea
### House Robber (Divide and Conquer) + Segement Tree
1. 突破点：每个query实际上是在求House Robber，但如果使用传统的DP-线性回头看(`dp[i] = max(nums[i] + dp[i-2], dp[i-1])`)的写法，会因为constraints而TLE。另一种House Robber的解法是Divide and Conquer。
2. 分治法: 把`nums[]`砍两半，合并两边的最优解。
    - 但本题我们不能直接简单的合并`dp[i][j] = dp[i][k] + dp[k+1][j]`，因为题目约束挨着的两个元素不能都取。也就是说，k元素和k+1元素不能同时取。
    - 把题目约束条件Generalize一下：我们需要在递归时，时刻监视每个子数组的头元素和尾元素。拼接时，左子数组的weave和右子数组的weave不能同时取。
```
 i       k  k+1         j
[a X X X v][v+c X X X X b]

Definition:
dp00[i][j] := 数组[i...j]头元素不取，尾元素不取
dp01[i][j] := 数组[i...j]头元素不取，尾元素取
dp10[i][j] := 数组[i...j]头元素取，尾元素不取
dp11[i][j] := 数组[i...j]头元素取，尾元素取

拼接原则：首尾取不取子数组与母数组保持一致，拼接处不能都取，i.e. 00, 01, 10

dp00[i][j]合法的拼接方法有三种:
1. dp00[i][j] = dp00[i][k] + dp00[k+1][j]
2. dp00[i][j] = dp01[i][k] + dp00[k+1][j]
3. dp00[i][j] = dp00[i][k] + dp10[k+1][j]
dp00[i][j]的最优解取这三种情况的最大值。

同理，
dp01[i][j]合法的拼接方法有三种，最优解取最大:
1. dp01[i][j] = dp00[i][k] + dp01[k+1][j]
2. dp01[i][j] = dp01[i][k] + dp01[k+1][j]
3. dp01[i][j] = dp00[i][k] + dp11[k+1][j]

dp10[i][j]合法的拼接方法有三种，最优解取最大:
1. dp10[i][j] = dp10[i][k] + dp00[k+1][j]
2. dp10[i][j] = dp11[i][k] + dp00[k+1][j]
3. dp10[i][j] = dp10[i][k] + dp10[k+1][j]

dp11[i][j]合法的拼接方法有三种，最优解取最大:
1. dp11[i][j] = dp10[i][k] + dp01[k+1][j]
2. dp11[i][j] = dp11[i][k] + dp01[k+1][j]
3. dp11[i][j] = dp10[i][k] + dp11[k+1][j]

Base cases: 分割到底，区间内只有一个元素
dp00[i][i] = 0
dp11[i][i] = nums[i]
dp01[i][i] = -inf // 取负无穷表示我们不能认可这种情况存在
dp10[i][i] = -inf // 同上

最后，一个query的最优解 = max(dp00[0][n-1], dp01[0][n-1], dp10[0][n-1], dp11[0][n-1])
```
3. 既然分治法是每次把数组砍一半，每一半记录整个区间的某一信息。自然地，这是一个线段数的结构。并且，只需要实现"单点更新"，不需要实现“区间更新”(每个query只改变一个值)和“区间查询” (因为每个query都只看0...n-1)
    - 线段数的两大优势: 1. 中间节点记录任意区间的最优解。2. "单点更新"更方便。每一个值/叶子节点更新只需要log(n)的时间复杂度。

Time complexity = $O(Q*\log n)$ where $Q = $ number of queries, $n = $ length of `nums[]`

## Resource
[【每日一题】LeetCode 3165. Maximum Sum of Subsequence With Non-adjacent Elements](https://www.youtube.com/watch?v=AzzaHQKkwnM&t=1502s&ab_channel=HuifengGuan)