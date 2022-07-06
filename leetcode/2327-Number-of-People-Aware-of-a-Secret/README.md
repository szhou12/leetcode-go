# [2327. Number of People Aware of a Secret](https://leetcode.com/problems/number-of-people-aware-of-a-secret/)

## Solution idea

### DP

这道题不寻常的点在于，如果按题目所问直接定义 $DP$ 是不够的。

i.e. $DP[i] = $ the number of people who know the secret at the end of day $i$

**注意到**，此时的 $DP[i]$ 既包含了第i天新增的人数，又包含了**第i天前**新增的、至今还没忘记的人数。即，$DP[i]$ 实际包含了两类人。

这样并不好使 $DP[i]$ 与 $DP[i-X]$ 建立联系。

如果，我们只定义一类人群，即，每天新增的人数呢？

Define $DP[i] = $ the number of NEWLY added people who know the secret on $i$-th day.

这样，似乎比较容易建立 Recurrence:

$DP[i] = \sum_j DP[j]$ such that $j\in (i-forget, i-delay]$