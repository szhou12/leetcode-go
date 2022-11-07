# [977. Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/description/)

## Solution idea

### Two Pointers - 双指针相向而行
* 不用找正负分界点，而是直接将双指针分别初始化在 nums 的开头和结尾
* 谁的绝对值大移动谁，输出数组从后往前添加较大绝对值元素的平方

Time complexity = $O(n)$

### 利用 MergeSort 里 merge 的思路：谁小移谁
* 首先寻找正负数的分界点，然后向左边逆序再平方，然后执行合并有序数组的逻辑

Time complexity = $O(n)$

## Resource
[代码随想录-977.有序数组的平方](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0977.%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E5%B9%B3%E6%96%B9.md)