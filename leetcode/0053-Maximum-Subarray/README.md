# [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)

## Solution idea
### DP (Maximum Subarray Sum) aka. Kadane's Algorithm
- DP template:
```
Definition:
    DP[i] := Maximum subarry sum ending at i.
Base Case:
    DP[0] = nums[0]
Recurrence:
    DP[i] = max(DP[i-1], 0) + nums[i] for 1 <= i < n
```

Time complexity = $$O(n)