# [3152. Special Array II](https://leetcode.com/problems/special-array-ii/description/)

## Solution idea
### DP
1. 思路类似于prefix sum。从左到右撸一遍`nums[]`，累计沿途遇到的parity的个数.
2. Dynamic Programming:
    - `dp[i] :=` number of dis-parity pairs ending at `nums[i]` (dis-parity means a pair of adjacent elements `dp[i]` and `dp[i-1]` are both odd or both even).
    - Base case: `dp[0] = 0`
    - Recurrence: `dp[i] = dp[i-1] + (nums[i] % 2 == nums[i-1] % 2)`
3. 判断一个subarray `[l, r]`是否"special", 只需要判断`dp[r] - dp[l] == 0`即可。

Time complexity = $O(n)$
