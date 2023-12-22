# [2963. Count the Number of Good Partitions](https://leetcode.com/problems/count-the-number-of-good-partitions/description/)

## Solution idea
### Sweep Line + Diff Array
1. 理解题意: no two subarrays contain the same number = 一种元素只能出现在一个subarray中
2. 破题: 对于一种元素，如何判断可以容纳它的最小可能的subarray在哪里? e.g., `[a X X g l X X g X y]` 对于 `X` 来说，最小的subarray可以通过它的**左、右边界**来确定。如果两种元素的最小subarray有重叠，那么，这两个subarrays应该合并成一个subarray来满足题意。通过 **确定左右边界 + 合并重叠subarray**，就可以得到最多可以产生的互不重叠的subarray。
3. Sweep Line: 使用**扫描线**计算互不重叠的subarray个数。每种元素的左边界定为 +1，右边界定位 -1。扫描线从左到右扫描，累积和。再次用扫描线从左向右扫描，每次遇到0说明扫过了一个无重叠的subarray，计数器+1。
4. 排列问题中的插空题型：n 个元素插入任意个挡板，可以有多少种分隔方式？答：n-1 个空位可以插挡板，每个空位选择插入或者不插，所以有 $2^{n-1}$ 种分隔方式。

Time Complexity = $O(n)$
## Resource
[【每日一题】LeetCode 2963. Count the Number of Good Partitions](https://www.youtube.com/watch?v=qd9xo40JCyk&t=619s&ab_channel=HuifengGuan)