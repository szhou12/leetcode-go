# [3082. Find the Sum of the Power of All Subsequences](https://leetcode.com/problems/find-the-sum-of-the-power-of-all-subsequences/description/)

## Solution idea
### DP (3-D)
1. 理解题意: 通过例子可以发现，题目描述有中有一层嵌套关系。即，从`nums[]`中先找到一个subsequnece，然后再在这个subsequence中数有多少个sub-subseqs，其元素的总和为 k，累计所有符合这个条件的sub-subseqs.
2. 切入点: 直接面对嵌套关系很难想到解法。想想**你会什么？**给一个array，找所有和为k的subsequences你会不会？好像容易一点:
    1. 分析：对于元素`i`，假设现在已知前`i-1`个元素中有多少subseqs并且它们的和是多少，那么现阶段我们要做的就是判断元素`i`要不要参与其中。
    2. 注意到，假设的已知信息有两条: 1. 有多少subseqs? 2. 它们对应的和是多少? 显然，定义一维DP是不够的: `dp[i]= `# of subseqs。所以，需要定义第二个维度动态描述和的情况！
    3. Definition: `dp[i][s] = ` # of subsequnces in `nums[1:i]` (inclusive) with their sum = `s`
    4. Recurrence: `dp[i][s] = dp[i-1][s] + dp[i-1][s-nums[i]]` (if `s-nums[i] >= 0`)
        1. `dp[i-1][s]`: `nums[i]`不参与, 和`s`已经够了
        2. `dp[i-1][s-nums[i]]`: `nums[i]`参与
    5. `dp[n][k]`就是长度为n的`nums[]`中有多少和为k的subseqs 
3. 现在解决了里层嵌套的问题，外层嵌套怎么解决？看待问题的角度稍微变一变。对于一个和为k的subseq，我们能不能知道有多少supersequence可以包含它？可以的，**我们只需要一个额外的信息(维度)：此时的subseq的长度是多少？**
    1. 假设，对于一种subseq的长度是l，我们可以从`nums`中找到多少个superseqs包含这个subseq?
    2. 答: $2^{n-l}$ (表示`nums`剩下的元素选和不选两种情况)
    3. $2^{n-l}$ `* dp[n][k][l]` = (多少superseqs) * (多少种长度l和为k的subseqs)
    4. Definition: `dp[i][s][l] = ` # of subseqs in `nums[1:i]` (inclusive) with their sum = `s` and length = `l` (除了表示和为k长度为l的subseqs的数量以外, 也表示该类型subseqs的种类)
    5. Recurrence: `dp[i][s][l] = dp[i-1][s][l] + dp[i-1][s-nums[i]][l-1]` (if `s-nums[i] >= 0` AND `l-1 >= 0`)
        1. `dp[i-1][s][l]`: `nums[i]`不参与, 和`s`以及长度`l`都已经够了
        2. `dp[i-1][s-nums[i]][l-1]`: `nums[i]`参与
    6. Base case: `dp[0][0][0] = 1` “种子”。表示空array中找到长度为0的子序列并且它的和是0, 这样的子序列有一种, 就是空子序列。

Time complexity = $O(n\cdot k \cdot n)$
## Resource
[【每日一题】LeetCode 3082. Find the Sum of the Power of All Subsequences](https://www.youtube.com/watch?v=N1fohv72yug&ab_channel=HuifengGuan)