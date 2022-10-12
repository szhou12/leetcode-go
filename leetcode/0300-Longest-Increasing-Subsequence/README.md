# [300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)

## Solution idea

### DP

`DP[i] = ` longest subseq of `nums[:i]` ending at index i (注意：ending at index i 指以i为结尾的数组的最长递增子序列, 说明`DP[n]`不是最后我们要求的, 因为全局的最长递增子序列不一定会包括最后一个index)

Base Cases:

```
DP[0] = 0
DP[1...n] = 1
```

Recurrence:

```
DP[i] = max(DP[j] + 1) for 1 <= j < i and nums[j] < nums[i]
```

Time complexity = $O(n^2)$