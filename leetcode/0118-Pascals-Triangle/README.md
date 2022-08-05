#[118. Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/)

## Solution idea
按照帕斯卡三角的机制写就可以：j-th position on i-th row = (j-1)-th position + j-th position on (i-1)-th row

Time complexity = $O(n^2)$ ($1+2+3+\cdots + n = \frac{n(n-1)}{2}$)