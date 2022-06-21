# [198. House Robber](https://leetcode.com/problems/house-robber/)

## Solution idea

### My Solution
Define $DP[i]$: max sum ending at i (e.g. max sum of `nums[0...i]`)

Recurrence:

Base case 1: $DP[0] = nums[0]$

Base case 2: $DP[1] = \max\{nums[0], nums[1]\}$

Case 1: $DP[i] = \max\{DP[j] + nums[i]\}$ where $0 \leq j < i-1$

Case 2: $DP[i] = DP[i-1]$

$DP[i] = \max\{ Case 1, Case 2 \}$

result $= DP[n-1]$

Time = $O(n^2)$

But my solution should work for the situation where there are negative entries.

### Optimizing Solution
Assume all entries $\geq 0$.

Recurrence:

$DP[i] = \max\{DP[i-1], DP[i-2] + nums[i]\}$ since we know $DP[i]$ must be non-decreasing as i increases.

Time = $O(n)$, Space = $O(n)$

### Optimizing Solution Again for Space

Since $DP[i]$ only depends on $DP[i-1], DP[i-2]$, we can use two variables to record them instead of creating a whole slice
