# [26. Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)

## Solution idea

### Two Pointers
双指针的物理意义：

慢指针 (`l`): 指向目前为止最后一个 dedup 过的元素

快指针 (`r`): 不停的往前走

机制：

每当`r`指向的元素值不同于`l`指向的元素的值，把`r`指向的元素值赋值给`l+1`

Time complexity = $O(n)$