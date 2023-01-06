# [532. K-diff Pairs in an Array](https://leetcode.com/problems/k-diff-pairs-in-an-array/description/)

## Solution idea

### Two Pointers - 一个数组 + 双指针同向而行

**Key 1**: Sort - 把input数组首先增序排列

**Key 2**: 为了避免重复，保证`left`指针跳过重复的元素。

**Key 3**: 另外注意的一个细节就是，每确定一个左指针`left`，都要重新定位右指针`right=left+1`

Time complexity = $O(n\log n) + O(n) = O(n\log n)$

## Resource

[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Two_Pointers/532.K-diff-Pairs-in-an-Array)