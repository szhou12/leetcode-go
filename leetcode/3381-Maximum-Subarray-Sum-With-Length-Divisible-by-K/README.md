# [3381. Maximum Subarray Sum With Length Divisible by K](https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/description/)

## Solution idea
### DP (Maximum Subarray Sum) + Prefix Sum
- 本题考查对 Classic Kadane's Algorithm ([53. Maximum Subarray](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0053-Maximum-Subarray)) 本质有多了解。
    1. Let's review **Classic Kadane's Algorithm**: 本质上是从任意长度的subarray (size >= 1)中找到和最大的
- DP template:
```
Definition:
    DP[i] := Maximum k-sized subarry sum ending at i.
Base Case:
    DP[k-1] = sum(nums[0:k-1]) (inclusive) = prefixSum[k-1]
Recurrence:
    DP[i] = max(DP[i-k], 0) + (prefixSum[i] - prefixSum[i-k]) for k <= i < n
```

Time complexity = $O(n)$

## Resource
- [Prefix Sum Kadane](https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/solutions/6124522/prefix-sum-kadane/)
    - Look for *David Espinosa* comment
- [DP in 5 Lines of Python3](https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/solutions/6124606/dp-in-5-lines-of-python3/)
- [7ms - O(N) - Kadane with benefits](https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/solutions/6124512/7ms-o-n-kadane-with-benefits/)