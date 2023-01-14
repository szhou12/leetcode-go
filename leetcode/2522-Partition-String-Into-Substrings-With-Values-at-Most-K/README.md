# [2522. Partition String Into Substrings With Values at Most K](https://leetcode.com/problems/partition-string-into-substrings-with-values-at-most-k/description/)

## Solution idea

### 1-D DP

Define `DP[i] = ` the minimum number of substrings in a good partition of `s[1...i]` (prepend `s` with a placeholder)

Base Case: `DP[0] = 0` because `s[0]` represents empty string

Recurrence (naive):

`DP[i] = 1 + min(DP[j-1])` for $1 \leq j < i $ and $s[j:i] \leq k$

(注意：这种写法需要$O(n^2)$，题目给出的量级会超时。需要通过观察`DP[i]`的性质，优化Recurrence)

**KEY: `DP[i]`的重要性质**
* 注意到，`DP[i]`一定是 **non-decreasing**. i.e. 长度为100的字符串的good partition数量一定**不可能**比长度99时的该字符串的good partition数量**少**
* 利用这个性质，每次更新`DP[i]`时，就不用看过去所有的`j`，只用看两个地方的`j`:
    1. `len(s[j:i]) == len(k)`: 首先看`j`使得`s[j:i]`位数与`k`的位数相同时的情况，如果此时框住的`s[j:i]`的值 $\leq k$, 说明`DP[i]`的最小是`DP[j-1]+1`。`j`不能再往前看了，因为框住的`s[j:i]`的位数比`k`的位数多，所以值一定比`k`大，不满足题意。
    2. 如果Case 1中框住的`s[j:i]`的值 $ > k$，说明相同位数时不能满足题意，那么，我们就需要缩短一位数`s[j+1:i]`，这样`s[j+1:i]`的值 一定$\leq k$ (e.g. 两位数一定小于三位数)，`DP[i]`的最小是`DP[j]+1`。`j`不用再往后看了，因为根据 **`DP[i]`的重要性质**，`j`再往后看，`DP[j]`也只会增加不会减少，所以肯定不存在最优

**注意**一下无解的情况。如果 `k` 长度是1，且`s[i]`里有一个字符大于`k`，那么说明即使区间长度是1 (最差情况下) 也无法满足要求。

Time complexity = $O(n)$

Space complexity = $O(n)$

## Resource

[【每日一题】LeetCode 2522. Partition String Into Substrings With Values at Most K](https://www.youtube.com/watch?v=k-3gzqWXAD4&ab_channel=HuifengGuan)