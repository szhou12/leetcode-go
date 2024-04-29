# [3132. Find the Integer Added to Array II](https://leetcode.com/problems/find-the-integer-added-to-array-ii/description/)

## Solution idea
### Degree of Freedom + Sorting
1. **sorting**使`nums1[]`与`nums2[]`的数组元素一一对应。相当于`nums1[]`删掉两个元素后，整体平移X，得到`nums2[]`。
2. **只删掉两个元素**，说明**自由度=2**, 也就是说sort以后，`nums1[]`头三个(n+1=2+1)元素一定至少有一个元素会被保留下来，参与平移X。你想想是不是？`nums1[]`只比`nums2[]`长度多2，任意拿掉两个元素，`nums1[]`的前三个无论如何都至少会有一个幸存下来。有可能幸存两个，或者三个，但至少一个。那么，这三个打头阵的元素`nums1[0], nums1[1], nums[2]`分别与`nums2[0]`之间的差值，**最小的**平移量X就是从这三个里面选出。
    - **最小**平移量X至多三个候选人: `nums2[0]-nums1[0]`, `nums2[0]-nums1[1]`, `nums2[0]-nums1[2]`。
3. 解法：分别检验这三个候选平移量X，一一比对是否`nums1[i]+X == nums2[j]`。允许容错率为2，也就是允许出现至多两次`nums1[i]+X != nums2[j]`，因为可以拿掉两个元素。如果一一检验完，容错率依旧保持在2或2以下，那么此时的候选平移量就可以参加竞选最小平移量。

Time complexity = $O(n\log n)$

## Resource
[LeetCode Weekly Contest 395 - All Solutions - Python](https://www.youtube.com/watch?v=L5O37E8GnTQ&ab_channel=Alpha-Code)
[1:35 - 7:35]