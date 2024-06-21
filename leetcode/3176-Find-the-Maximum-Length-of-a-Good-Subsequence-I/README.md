# [3176. Find the Maximum Length of a Good Subsequence I](https://leetcode.com/problems/find-the-maximum-length-of-a-good-subsequence-i/description/)

## Solution idea
### 2-D DP
1. Definition: `DP[i][t] :=` maximum possible length of a good subsequence of `nums[0:i]` **including i-th element** && there costs `t` times of `seq[i] != seq[i + 1]` ($0 \leq t \leq k$)
2. Base Case: `DP[0][0...k] = 1`. (无论“花费”多少，一个元素的情况只能自立门户)
3. Recurrence: `DP[i][t] = max(1, DP[j][t]+1)` for all `0 <= j < i` and `0 <= t <= k` if `nums[i] == nums[j]` or `DP[i][t] =  max(1, DP[j][t-1]+1)` for all `0 <= j < i` and `1 <= t <= k` if `nums[i] != nums[j]`
```
[X X X X j] X X i
if nums[i] == nums[j]:
    DP[i][t] = max(1, DP[j][t] + 1)
else if nums[i] != nums[j] && t-1 >= 0:
    DP[i][t] = max(1, DP[j][t-1] + 1)
```
4. 翻译：站在第 i 号元素上，回头看所有过去的元素，对于任意一个过去的 j 号元素，如果`nums[i] == nums[j]`，那么我们不用“花费”任何费用，直接把 i 接在 j 后面。如果`nums[i] != nums[j]`，我们就要“花费”一次允许不相等的费用，再把 i 接在 j 后面。
5. 注意！注意！注意！
    1. 对于第二种情况 (`nums[i] != nums[j]`)，需要检查边界条件 `t-1 >= 0`，因为显然当前次是不相等的情况，我们至少有一次“花费”，transition最极端的情况是从之前0次“花费”到现在1次“花费”，而不能从-1次“花费”到0次“花费”。
    2. 看回到`DP[i][t]`的定义，站在第 i 号元素上，`DP[i][t]`必须包含第 i 号元素，因此，我们可以start over，在第 i 号元素自立门户，也就是`DP[i][t] = 1`。
    3.  看回到`DP[i][t]`的定义，站在第 i 号元素上，`DP[i][t]`必须包含第 i 号元素，因此，全局最优解不一定出现在最后`n-1`处，而是应该在loop i and loop k 的双重循环中，随时更新全局最优解。

Time complexity = $O(n\cdot k \cdot n)$ (这个解法只适用于3176的数据量级)

## Resource
[【每日一题】LeetCode 3177. Find the Maximum Length of a Good Subsequence II](https://www.youtube.com/watch?v=lEOzA3KXV5s&t=83s&ab_channel=HuifengGuan) (0:00 ~ 12:23)