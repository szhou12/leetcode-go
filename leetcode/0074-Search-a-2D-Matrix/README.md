# [74. Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/)

## Solution idea
### Binary Search
1. 把2-D matrix 看成一个1-D array (indexed from left to right, then top to bottom 0 -> m*n-1)
2. Then for any 1-D index `mid`, its corresponding 2-D index is `(mid/n, mid%n)`
3. 然后就可以用二分搜索了

Time complexity = $O(\log(m*n))$

Space complexity = $O(1)$