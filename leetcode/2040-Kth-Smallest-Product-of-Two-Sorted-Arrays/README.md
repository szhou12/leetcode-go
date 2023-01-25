# [2040. Kth Smallest Product of Two Sorted Arrays](https://leetcode.com/problems/kth-smallest-product-of-two-sorted-arrays/description/)

## Solution idea

### Binary Search 猜答案

**Key**: input array中有负数，所以matrix中的product可正可负，所以此时需要**分类讨论**:

Given `nums1[i]`, find `j` s.t. `nums1[i] * nums2[j] <= mid` in order to find # of products less than or equal to `mid`

1. **Case 1**: if `nums1[i] > 0`, then `nums2[j] <= mid / nums1[i]`. Since `nums2` sorted in decreasing-order, `j` should point to the last element `<= mid / nums1[i]`. Thus, `0...j` all satisfy the requirement. # of products less than or equal to `mid` = `j-0+1`.
    * Note 1: How to find `j`? Binary search on `nums2` with target value `mid / nums1[i]`
        * Use `upperBound()` method to find the first index `k` s.t. `nums2[k] > target`. Thus, `j=k-1`
    * Note 2: It may NOT the case that `nums1[i] | mid`. Thus, take `Floor(mid / nums1[i])`
2. **Case 2**: if `nums1[i] = 0`, then `0 * nums2[j] <= mid`
    * **Case 2.**1: if `mid < 0`, then no `j` can work. # of products less than or equal to `mid` = 0
    * **Case 2.2**: if `mid >= 0`, then all `j` will work.  # of products less than or equal to `mid` = `len(nums2)`
3. **Case 3**: if `nums1[i] < 0`, then `nums2[j] >= mid / nums1[i]`. Since `nums2` sorted in decreasing-order, `j` should point to the first element `>= mid / nums1[i]`. Thus, `j...n-1` all satisfy the requirement. # of products less than or equal to `mid` = `n-1-j+1`.
    * Note 1: How to find `j`? Binary search on `nums2` with target value `mid / nums1[i]`
        * Use `lowerBound()` method to find the first index `k` s.t. `nums2[k] >= target`. Thus, `j=k`
        * Note 2: It may NOT the case that `nums1[i] | mid`. Thus, take `Ceil(mid / nums1[i])`

Time complexity = $O(\log n * (n \log n)) = O(n\log n)$

## Resource
[【每日一题】LeetCode 2040. Kth Smallest Product of Two Sorted Arrays, 10/28/2021](https://www.youtube.com/watch?v=Ct-seYTr1dM&ab_channel=HuifengGuan)