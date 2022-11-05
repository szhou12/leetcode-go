# [724. Find Pivot Index](https://leetcode.com/problems/find-pivot-index/description/)

## Solution idea

* 第i个元素作为挡板
* 在第i轮：
    * rightSum - 第i个元素
    * 判断是否 leftSum == rightSum
    * leftSum + 第i个元素

Time complexity = $O(n)$