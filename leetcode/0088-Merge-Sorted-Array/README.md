# [88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/)

## Solution ideas

### 1. Merge part in Merge Sort

### 2. Optimal Solution = loop backwards + select larger element each time
为了不大量移动元素，就要从2个数组长度之和的最后一个位置开始，依次选取两个数组中大的数，从第一个数组的尾巴开始往头放，只要循环一次以后，就生成了合并以后的数组了。

Time complexity = $O(m+n)$

[Reference](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0088.Merge-Sorted-Array)