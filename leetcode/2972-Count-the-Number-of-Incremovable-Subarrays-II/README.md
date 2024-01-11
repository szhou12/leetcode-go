# [2972. Count the Number of Incremovable Subarrays II](https://leetcode.com/problems/count-the-number-of-incremovable-subarrays-ii/description/)

## Solution idea
### Binary Search
1. 理解题意：找到一段 subarray，删去以后，剩下的元素严格递增。
2. 把 array 分成三部分：L + M + R。L, R 两部分严格递增，M 是要删去的 subarray。
    * find the largest index L s.t. `[0 ... L]` is strictly increasing
    * find the smallest index R s.t. `[R ... len-1]` is strictly increasing
3. 机制：for any `i` in `[0 ... L]`, find the first index `j` in `[R ... len-1]` s.t. `nums[i] < nums[j]`
    * 对于一个 `i`, 可以删去的 subarray 是 `[i+1 ... j]`。这样的 `j` 有 `(len-1) - j + 1` 个。注意：`[i+1 ... j-1]` 也是一个可以删去的 subarray，操作 append `MaxInt` at the end 可以把 这个 subarray 包含进 `[i+1 ... j]`，也就是说，`(len-1) - j + 1` 可以同时算入 `[i+1 ... j-1]`
    * 使用 binary search 的 `upperBound()`
3. 考虑三种特殊情况的处理方法：
    * L部分为空：prepend `MinInt` at the beginning。`i=0`就是L部分为空的情况，可以删去的subarray个数就是`[0...j]` where `R <= j <= len-1`
    * R部分为空：append `MaxInt` at the end。作用一：标识R部分为空。作用二：详见第3项。
    * M部分为空：意味着整个 array 严格递增。使用排列组合方法计算任意subarray总数 = 任意选择一个index作为起点，任意选择另一个index作为终点，一共有 $\frac{n \times (n-2)}{2}$ 种选法 + 一个元素作为subarray的情况 (起点终点在同一个index)，一共有 $n$ 种选法。所以一共有 $\frac{n \times (n-2)}{2} + n$ 种选法。
    
Time complexity = $O(n\log n)$

## Resource
[【每日一题】LeetCode 2972. Count the Number of Incremovable Subarrays II](https://www.youtube.com/watch?v=aBikXIvNtNU&ab_channel=HuifengGuan)