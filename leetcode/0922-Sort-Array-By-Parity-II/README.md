# [922. Sort Array By Parity II](https://leetcode.com/problems/sort-array-by-parity-ii/description/)

## Solution idea

### Naive Solution
* 开两个数组分别存下偶数元素，奇数元素; 再合并起来

Time complexity = $O(n)$, Space complexity = $O(n)$

### In-place
* 检查每个偶数位置，如果发现是奇数元素，则从奇数位置从小到大找存有偶数元素的，进行交换

Time complexity = $O(n)$, Space complexity = $O(1)$

## Resource
[代码随想录-922. 按奇偶排序数组II](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0922.%E6%8C%89%E5%A5%87%E5%81%B6%E6%8E%92%E5%BA%8F%E6%95%B0%E7%BB%84II.md)