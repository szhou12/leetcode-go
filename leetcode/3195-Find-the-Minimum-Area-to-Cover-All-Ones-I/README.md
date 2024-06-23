# [3195. Find the Minimum Area to Cover All Ones I](https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/description/)

## Solution idea
```
 0  0 0 0  0
    _____
 0 |0 1 0| 0 
   |     |
 0 |1 0 0| 0 
   |     |
 0 |0 0 1| 0 
    -----
 0  0 0 0  0 
```
1. 找到能覆盖所有1的最小矩形，首先关注到左上角和右下角，确定它们的坐标。
2. 最小矩形的行和列的坐标可以分开独立来找，也就是说，行和列并不是互相关联，不会牵一发而动全身。那么，确定左上角的坐标的行 = 出现1的最小行，列 = 出现1的最小列；确定右下角的坐标的行 = 出现1的最大行，列 = 出现1的最大列。
3. 最后注意处理没有出现1的情况。

Time complexity = $O(M\cdot N)$