# [373. Find K Pairs with Smallest Sums](https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/)

## Solution idea

### Binary Search 猜答案

* **KEY**: 我们有two sorted input arrays，如果构造一个matrix，其中一个array作行，另一个array作列，每一个element是sum，我们可以得到一个行与列都是sorted matrix。

* 一旦察觉到上述要素，剩下的解法基本上就与[378. Kth Smallest Element in a Sorted Matrix](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/)一致: 从左下/右上角出发，计数小于等于当前元素的个数是否为k

* 此题，我们不用显性地生成这个matrix，只需要操作`nums1` 和 `nums2` 的index 来模拟 matrix traversal

Time complexity = $O((m+n)\log D + m*k) = O((m+n)\log D)$ where $m = $ length of `nums1`, $n = $ length of `nums2`, $D = $ difference between the max value in the sum matrix and the min value in the sum matrix.

## Resource

[【每日一题】LeetCode 373. Find K Pairs with Smallest Sums](https://www.youtube.com/watch?v=TsOzIxkzh1E&ab_channel=HuifengGuan)