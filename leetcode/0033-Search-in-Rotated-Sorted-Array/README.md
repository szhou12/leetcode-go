# [33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)

## Solution idea

本题的题眼1：经过rotate, 升序排列的数组自然分成两段升序排列的数组 (左区间，右区间).

本题的题眼2: 左区间，右区间都和**一个元素**`nums[0]`(突破口!!!)有关系. 左区间的元素都 $\geq$ `nums[0]` (等于 因为包括`nums[0]`自己), 右区间的元素都 $<$ `nums[0]`.

本题的题眼3: 因为有左区间，右区间两个区间, `nums[mid]` 和 `target` 就有可能随机落在任意一个区间. 本题的关键就是判别`nums[mid]` 和 `target` 分别落在了哪个区间.

Time complexity = $O(\log n)$

## Resource

[33. Search-in-Rotated-Sorted-Array](https://github.com/wisdompeak/LeetCode/tree/master/Binary_Search/033.Search-in-Rotated-Sorted-Array)