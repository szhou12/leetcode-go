# [1539. Kth Missing Positive Number](https://leetcode.com/problems/kth-missing-positive-number/description/)


## Solution idea

### Binary Search 猜答案

* 对当前一个猜测的答案 `mid`, 
    * 小于`mid`的自然数总数是 `mid - 1` (`[1, mid-1]`), 记为 `T`
    * 计算有多少小于`mid`的数已经存在于`arr`, 使用 `lowerBound()`找到第一个 `>= mid` 的index, 那么之前的数就都是满足条件的 (利用 `arr` 单增的特性), 记为 `P`
    * `missing = T-P = ` 小于`mid`的缺失自然数的个数
        * if `missing <= k-1`, `mid` 有可能是答案, `left = mid`
        * if `missing > k-1`, 说明比 `mid`小的缺失自然数个数都已经至少k个了, `mid`绝对不会是第k个缺失的自然数, `right = mid - 1`
    * 特别注意，当 `missing == k-1` 的时候，`mid` 可能并不是最终答案，因为`mid`可能也存在于数组中，所以`mid`可以再往大猜(即`left=mid`)。因此这个分支在上面的代码里与`missing < k-1`合并

Time complexity = $O(\log D * \log N)$ where $D = $ difference between largest possible value (`arr`最后一个元素往后数k个) - smallest possible value (1)


## Resource

[http://wisdompeak/LeetCode/Binary_Search/1539.Kth-Missing-Positive-Number/](https://github.com/wisdompeak/LeetCode/tree/master/Binary_Search/1539.Kth-Missing-Positive-Number)