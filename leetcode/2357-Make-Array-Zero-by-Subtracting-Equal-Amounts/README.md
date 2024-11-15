# [2357. Make Array Zero by Subtracting Equal Amounts](https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/description/)

## Solution idea
### 逻辑思维 + HashMap
思路类似[1526. Minimum Number of Increments on Subarrays to Form a Target Array](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1526-Minimum-Number-of-Increments-on-Subarrays-to-Form-a-Target-Array)的简化版：总是挑全局最小的正数降0，比它大的数会跟着减少对应的次数，并且需要额外一次操作把剩余的量降为0。所以，最小操作次数就是所有正数的种类数。

Time complexity = $O(n)$