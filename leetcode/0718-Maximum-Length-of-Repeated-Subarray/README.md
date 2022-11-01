# [718. Maximum Length of Repeated Subarray](https://leetcode.com/problems/maximum-length-of-repeated-subarray/)

## Solution idea

### DP - Longest common subarray

* Definition:
```
DP[i][j] = max length of longest common subarray for nums1[1...i] ending at i and nums2[1...j] ending at j
```
    * Note: `nums1` and `nums2` prepend with a placehold so meaningful elements start from index = 1
* Base Cases:
```
dp[0][0] = 0
dp[i][0] = 0 for all i
dp[0][j] = 0 for all j
```
* Recurrence
```
dp[i][j] = dp[i-1][j-1] + 1 if nums1[i] == nums2[j]
         = 0                otherwise
```

* **重点**: 注意这里状态转移方程(Recurrence)与 LCS 类型题的区别
    * 根据`DP[i][j]`定义, **必须结尾在**`nums1[i]` 和 `nums2[j]`
    * 如果`nums1[i]` != `nums2[j]`, 说明此时不可能存在 一个以`i`结尾的`nums1`的subarray 与 一个以`j`结尾的`nums2`的subarray 相同 (显然了, `i, j`指向的元素都不同), 所以直接是 0
    * 注意: 在 longest common subsequence 题型中, 这里是 `max(dp[i][j-1], dp[i-1][j])`, 这是因为`i, j`指向的元素不一定非要贡献进subsequence里

Time complexity = $O(m*n)$

## Resource
[【每日一题】718. Maximum Length of Repeated Subarray, 7/2/2020](https://www.youtube.com/watch?v=1NvUmfI5u8Y)