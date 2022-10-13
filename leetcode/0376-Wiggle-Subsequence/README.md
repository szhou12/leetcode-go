# [376. Wiggle Subsequence](https://leetcode.com/problems/wiggle-subsequence/)

## Solution idea

### DP
```
dp[i][0] = longest wiggle subseq ending at i and diff = nums[i] - prev < 0
dp[i][1] = longest wiggle subseq ending at i and diff = nums[i] - prev > 0
Base Cases:
dp[0][0/1] = 0
dp[1...n][0/1] = 1
Recurrence:
dp[i][0] = max(dp[j][1]+1) where j < i and nums[i] - nums[j] < 0
dp[i][1] = max(dp[j][0]+1) where j < i and nums[i] - nums[j] > 0
Output:
return global max
```

Time complexity = $O(n^2)$, Space complexity = $O(n)$

### Greedy Algorithm

* 统计整个序列有最多的局部峰值，从而达到最长摆动序列

* 因为题目要求的是最长摆动子序列的长度，所以只需要统计数组的峰值数量就可以了（相当于是删除单一坡度上的节点，然后统计长度）

* 贪心所贪的地方: 让峰值尽可能的保持峰值，然后"删除"(忽略)单调区间上的节点

Time complexity = $O(n)$, Space complexity = $O(1)$

## Resource
[代码随想录-376. 摆动序列](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0376.%E6%91%86%E5%8A%A8%E5%BA%8F%E5%88%97.md)