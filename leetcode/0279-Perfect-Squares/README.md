# [279. Perfect Squares](https://leetcode.com/problems/perfect-squares/)

## Solution idea

### DP

`DP[i] := ` minimum number of perfect square numbers that sum to `i`

* Base Cases:
```
DP[0] = 0
DP[1] = 1 
```

我最开始的Recurrence构造方法是:
```
dp[i] = 1 if i is a perfect square
      = min(dp[j] + 1) for 1 <= j <= i-1 && i-j is a perfect square o/w
```
这样写会超时，因为对每个`i`, 相当于把它前面的所有数都看了一遍.

如下定义更效率，物理意义也更清晰:
```
dp[i] = 1 if i is a perfect square
      = min(dp[i-j*j] + 1) for 1 <= j*j <= i o/w
```
这里，每个`j`的物理意义是小与`i`的所有可能的平方数


Time complexity = $O(n^2)$

## Resource
[279. Perfect Squares](https://leetcode.com/problems/perfect-squares/)

