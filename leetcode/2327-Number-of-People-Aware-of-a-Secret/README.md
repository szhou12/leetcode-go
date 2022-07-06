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

**$DP[i] = \sum_j DP[j]$ such that $j\in (i-forget, i-delay]$**

翻译：第i天新增的人数由哪些天新增的人数贡献来的呢？显然，第`i-forget`天以及之前都不行，因为他们刚好是最后一批，到第i天会刚好忘记secret的人群。

那么，最遥远的一批可以贡献新增人数的人群来自于在第`i-forget+1`天新增的人群。

那么，哪一天产生最后一批可以贡献新增人数的人群呢？

显然，第`i-delay`天产生的新增人群是最后一批可以对第i天贡献新增人数的人群。

那么，这两天中间的每一天新增的人群都可以做贡献了。

$DP[i]$ 现在定义好了，下一个问题是，如何和题目所求建立联系呢？

我们知道，题目问的是：在第n天新增的人数，以及截止第n天还没忘记的人数。这两类人群。

其实，这第二类人群可以转化一下描述：

第一类人：在第n天新增的人数，

第二类人：**第n天前**新增的、至今还没忘记的人数。

这样，我们的 $DP$ 不就涵盖了两类人群了吗。

所以，题目要求的就是：$\sum_i DP[i]$ such that $i+forget>n$.

另一个值得注意的点是：这道题的 $DP$ 可以回头看，也可以向前看。

如果回头看，就是看 哪些天贡献 $DP[i]$. 即，$j\in (i-forget, i-delay]$.

如果向前看，就是看 $DP[i]$ 可以贡献给未来哪些天。 即，$j\in [i+delay, i+forget)$. (此处，$DP[i]$ 作为已知量)