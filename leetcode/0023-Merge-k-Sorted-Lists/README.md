# [23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/description/)

## Solution idea
### Min Heap
1. Alwasy pop the node head with the smallest value, and push its next node (if not nil) back into the heap.

Time complexity = $O(n\logk)$