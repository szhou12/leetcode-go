# [342. Power of Four](https://leetcode.com/problems/power-of-four/)

## Solution idea
不断对input除以4，如果input不是$4^x$, 那么总会在某一步时`mod 4`会不等于$0$.

Time complexity = $O(n)$