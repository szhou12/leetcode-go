# [15. 3Sum](https://leetcode.com/problems/3sum/)

## Solution idea

**Key 1**: Sort first in ascending order

**Key 2**: 对每个元素，用**相向而行双指针**找对应的two sum (a pair)

**Key 3**: 第i个元素可以重复出现在不同triplet中，但是，不能有重复的triplet (里面的三个值不能完全一样), 所以，要跳过之后所有与第i个元素相同的元素. 也就是说，对每个元素i我们在 `while (left < right)`中, 已经找到了所有可能的triplets.

Time complexity = $O(n\log n + n ^2)$ = $O(n^2)$