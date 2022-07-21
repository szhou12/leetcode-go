# [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)

## Solution idea
**Binary Search**

思路：找到array中最后一个小于 target的数 (the largest element < target)

注意：最后要分情况处理

Case 1: 如果找不到最后一个小于target的数，即，收敛完后的元素 $\geq$ target, 那么就挤占这个元素的位置。

Case 2: 如果找到了最后一个小于target的数，那么应该插在这个元素的后一位。

Time complexity = $O(\log n)$