# [228. Summary Ranges](https://leetcode.com/problems/summary-ranges/)

## Solution idea

因为题目是 **sorted unique** array, 所以一段区间内一定单增，且相邻元素差值 $\geq 1$. 

所以，我的想法是**快慢指针** 并使用一个 `count`.

慢指针指向一段连续区间的头，快指针不断往前。

如果当前 `nums[fast]-nums[slow]==count`，说明区间`[slow, fast]`是连续的.

如果当前 `nums[fast]-nums[slow]>count` (只可能`>`, 因为数组单增), 说明`fast` 和 `fast-1`出现断崖，此时`[slow, fast-1]` 为一段连续区间，并纳入result中。

**注意**：for-loop终止后，要post-processing 把最后一段区间纳入result中.

Time complexity = $O(n)$