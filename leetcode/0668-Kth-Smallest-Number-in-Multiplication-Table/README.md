# [668. Kth Smallest Number in Multiplication Table](https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/description/)

## Solution idea

### Binary Search - Guess K-th Element

* 题目的 Multiplication Table 都是正数相乘得到，所以 Sorted Table. 思路就与 [378. Kth Smallest Element in a Sorted Matrix](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/) 一样

* 小细节: 本题table是 1-indexed, 所以涉及matrix index的部分要做微小调整

Time complexity = $O((m+n)\log D)$ where $D = m*n - 1$
