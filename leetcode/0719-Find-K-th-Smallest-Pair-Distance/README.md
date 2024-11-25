# [719. Find K-th Smallest Pair Distance](https://leetcode.com/problems/find-k-th-smallest-pair-distance/description/)

## Binary Search 猜答案
1. 题目要求第k小的配对产生的绝对差值。那么，我们就用二分法高效地猜一个差值，数小于等于该差值的pair数是否够k个。如果此时pair数小于k个，说明猜小了；否侧 (大于等于k个)，说明猜的值有可能是解。
2. 那么，问题来到了，如何高效地数pair数？
    - 数所有pair数需要$O(n^2)$
    - 但是，如果给一个target差值`diff`和减数`nums[i]`，在一个排好序的数组中，可以使用`upperBound()`二分法找到第一个大于`target+nums[i]`的index `j`。那么，对于`i`而言，所有小于等于`diff`的pair数 = `(j-1) - i`。
    - 再Loop `i`就可以找到所有符合条件的pair数。

Time complexity = $O(\log n * n\log n)$


## Resource
[【每日一题】719. Find K-th Smallest Pair Distance, 10/26/2019](https://www.youtube.com/watch?v=rKVivKCchFc&ab_channel=HuifengGuan)
