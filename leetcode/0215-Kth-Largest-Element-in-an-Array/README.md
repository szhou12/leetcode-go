# [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

## Solution idea

### Solution - Min Heap
Use `min heap` to store `k` largest element: every time the top element in min heap is smaller than the current element, pop then push.

Time complexity = $O(n\log k)$

### Solution - Binary Search by Value
Binary search on `MinInt` and `MaxInt`. For each `mid`, count the number of elements in `nums` that are $\geq$ to the `mid`.
(Think `nums` as a sorted array will help visualize the binary search case)

Time complexity = $O(n\log C)$ where $C = 2^{32}$.